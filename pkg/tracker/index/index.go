// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package index

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/timestamp"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/http"
	"github.com/tellor-io/telliot/pkg/util"
	"github.com/yalp/jsonpath"
)

const ComponentName = "indexTracker"

type Config struct {
	Interval       util.Duration
	FetchTimeout   util.Duration
	ApiFile        string
	ManualDataFile string
}

type IndexTracker struct {
	logger       log.Logger
	ctx context.Context
	stop context.CancelFunc
	tsDB        *tsdb.DB
	cfg         *Config
	dataSources map[string][]DataSource
	prices      *prometheus.GaugeVec
	volumes     *prometheus.GaugeVec
	getErrors   *prometheus.CounterVec
}

func New(
	logger log.Logger, 
	ctx context.Context, 
	cfg *Config, 
	tsDB *tsdb.DB, 
	client contracts.ETHClient,
	) (*IndexTracker, error) {
	dataSources, err := createDataSources(logger, ctx, cfg, client)
	if err != nil {
		return nil, errors.Wrap(err, "create data sources")
	}

	ctx,stop:=context.WithCancel(ctx)

	return &IndexTracker{
		logger:       log.With(logger, "component", ComponentName),
		ctx:ctx,
		stop:stop,
		dataSources: dataSources,
		tsDB:        tsDB,
		cfg:         cfg,
		getErrors: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "errors_total",
			Help:      "The total number of get errors. Usually caused by API throtling.",
		}, []string{"source"}),
		prices: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "price",
			Help:      "The currency price",
		},
			[]string{"source"},
		),
		volumes: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "volume",
			Help:      "The currency trade ammount",
		},
			[]string{"source"},
		),
	}, nil
}

func createDataSources(logger log.Logger, ctx context.Context, cfg *Config, client contracts.ETHClient) (map[string][]DataSource, error) {
	// Load index file.
	byteValue, err := ioutil.ReadFile(cfg.ApiFile)
	if err != nil {
		return nil, errors.Wrapf(err, "read index file path:%s", cfg.ApiFile)
	}
	// Parse to json.
	indexes := make(map[string][]Api)
	err = json.Unmarshal(byteValue, &indexes)
	if err != nil {
		return nil, errors.Wrap(err, "parse index file")
	}

	dataSources := make(map[string][]DataSource)

	for symbol, apis := range indexes {
		for _, api := range apis {
			api.URL = os.Expand(api.URL, func(key string) string {
				return os.Getenv(key)
			})

			var source DataSource

			// Default value for the api type.
			if api.Type == "" {
				api.Type = httpIndexType
			}

			if int64(api.Interval.Duration) != 0 {
				api.Interval = cfg
			}

			// Default value for the parser.
			if api.Parser == "" {
				api.Parser = jsonPathParser
			}
			switch api.Type {
			case httpIndexType:
				{
					source = NewJSONapi(logger, api.Interval, cfg.FetchTimeout.Duration, api.URL, NewParser(api.Parser))
				}
			case fileIndexType:
				{
					source = NewJSONfile(api.URL, NewParser(api.Parser))
				}
			case ethereumIndexType:
				{
					// Getting current network id from geth node.
					networkID, err := client.NetworkID(ctx)
					if err != nil {
						return nil, err
					}
					// Validate and pick an ethereum address for current network id.
					address, err := util.GetAddressForNetwork(api.URL, networkID.Int64())
					if err != nil {
						return nil, errors.Wrap(err, "getting address for network id")
					}
					if api.Parser == uniswapParser {
						source = NewUniswap(symbol, address,api.Interval, client)

					} else if api.Parser == balancerParser {
						source = NewBalancer(symbol, address,api.Interval, client)
					} else {
						return nil, errors.Wrapf(err, "unknown source for on-chain index tracker")
					}
				}
			default:
				return nil, errors.New("unknown index type for index object")
			}

			dataSources[symbol] = append(dataSources[symbol], source)
		}

	}
	return dataSources, nil

}

func (self *IndexTracker) Run() error {
	for symbol,dataSource:= self.dataSources{
		// Use the default interval when not set.
		interval := dataSource.Interval.Duration
		if int64(interval) == 0{
			interval = self.cfg.Interval
		}

		go func(symbol string,interval time.Duration,dataSource DataSource){
			ticker := time.NewTicker(interval)

			logger = log.With(logger, "source", dataSource.Source())
			for {
				select{
				case <-self.ctx.Done():
					level.Debug(self.logger).Log("msg","data source loop exited","source", dataSource.Source())
					return
				default:
				}
				appender := self.tsDB.Appender(ctx)

				vals,err:= dataSource.Get(self.ctx)
				if err!=nil{
					level.Error(logger).Log("msg","getting values from data source","err",err)
					self.getErrors.With(prometheus.Labels{"source":dataSource.Source()}).(prometheus.Counter).Inc()
				}
				appender.Append(0, labels.Labels{labels.Label{Name: "__name__", Value: symbol + "_value"}}, timestamp.FromTime(time.Now().Round(0)), vals[0])
				
				// TODO for manual entries this is broken.
				// It should return 0 volume.
				// At the moment it returns the timestamp.
				appender.Append(0, labels.Labels{labels.Label{Name: "__name__", Value: symbol + "_volume"}}, timestamp.FromTime(time.Now().Round(0)), vals[1])

				if err := appender.Commit(); err != nil {
					level.Error(logger).Log("msg","adding values to the DB","err",err)
					continue
				}

				self.price.With(prometheus.Labels{"source": url}).(prometheus.Gauge).Set(vals[0])
				self.volume.With(prometheus.Labels{"source": url}).(prometheus.Gauge).Set(vals[1])

				<-ticker.C
			}

		}(symbol,interval,dataSource)
	}
}

func (self *IndexTracker) Stop() {
	self.stop()
}

// IndexType -> index type for Api.
type IndexType string

const (
	httpIndexType     IndexType = "http"
	ethereumIndexType IndexType = "ethereum"
	fileIndexType     IndexType = "file"
)

// ParserType -> index parser for Api.
type ParserType string

const (
	fileParser ParserType = "jsonPath"
	jsonPathParser ParserType = "jsonPath"
	uniswapParser  ParserType = "Uniswap"
	balancerParser ParserType = "Balancer"
)

// Api will be used in parsing index file.
type Api struct {
	URL      string          `json:"URL"`
	Type     IndexType       `json:"type"`
	Parser   ParserType      `json:"parser"`
	Param    string          `json:"param"`
	Interval config.Duration `json:"interval"`
}

func NewJSONapi(logger log.Logger, interval time.Duration, retryDelay time.Duration, url string,parser Parser) *JSONapi {
	return &JSONapi{
		logger:     logger,
		url:        url,
		retryDelay: retryDelay,
		interval:   interval,
		parser: parser,
	}
}

type JSONapi struct {
	url        string
	logger     log.Logger
	retryDelay time.Duration
	interval   time.Duration
	parser Parser
}

func (self *JSONapi) Get(ctx context.Context) (float64, error) {
	
	vals,err:= http.Fetch(ctx, self.logger, self.url, self.retryDelay)
	if err!=nil{
		return 0,errors.Wrap(err,"fetching data from API")
	}
	 return self.Parser(vals)
}

func (self *JSONapi) Interval() time.Duration {
	return self.interval
}

func (self *JSONapi) Source() time.Duration {
	return self.url
}

func NewJSONfile(filepath string) *JSONfile {
	return &JSONfile{filepath: filepath}
}

type JSONfile struct {
	filepath string
}

func (self *JSONfile) Get(_ context.Context) (float64, error) {
	return ioutil.ReadFile(self.filepath)
}

func (self *JSONfile) Interval() time.Duration {
	return 0
}

func (self *JSONapi) Source() time.Duration {
	return self.filepath
}

type DataSource interface {
	// Source returns the data source.
	Source() string
	// Get returns current index price and volume.
	Get(context.Context) ([]float64, error)
	// The recommended interval for calling the Get method.
	// Some APIs will return an error if called more often
	// Due to API rate limiting of the provider.
	Interval() time.Duration 
}

type Parser interface {
	Parse([]byte)(vals []float64, err error)
}

type FileParser struct{}

func (*FileParser) Parse(payload []byte)(vals []float64, err error){
	var decodedPayload, result interface{}
	err = json.Unmarshal(payload, &decodedPayload)
	if err != nil {
		return 0,
	}
}

type JsonPathParser struct{}

func (*JsonPathParser) Parse(payload []byte)(vals []float64, err error){
	result = decodedPayload
	if len(strings.TrimSpace(self.Param)) > 0 {
		if err != nil {
			return
		}
		result, err = jsonpath.Read(decodedPayload, self.Param)
		if err != nil {
			return
		}
	}

	// Parse each item of slice to a float.
	vals = make([]float64, 0)
	for _, a := range resultList {
		strValue := fmt.Sprintf("%v", a)
		// Normalize based on american locale.
		strValue = strings.Replace(strValue, ",", "", -1)
		val, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			return nil, errors.Wrap(err, "JSON value needs to be a valid float")
		}
		vals = append(vals, val)
	}
}


func NewParser(t ParserType) Parser {
	switch t {
	case jsonPathParser:
		return &JsonPathParser{}
	default: 
		return &FileParser{}
	}
}

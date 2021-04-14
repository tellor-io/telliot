// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package index

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tellor-io/telliot/pkg/apiOracle"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/http"
	"github.com/tellor-io/telliot/pkg/util"
	"github.com/yalp/jsonpath"
)

const ComponentName = "index"

var clck clock.Clock

func init() {
	clck = clock.New()
}

var indexes map[string][]*IndexTracker

// GetIndexes returns indexes for outside package usage.
func GetIndexes() map[string][]*IndexTracker {
	return indexes
}

// BuildIndexTrackers creates and initializes a new tracker instance.
func BuildIndexTrackers(logger log.Logger, cfg *config.Config, db db.DataServerProxy, client contracts.ETHClient) ([]*IndexTracker, error) {
	err := apiOracle.EnsureValueOracle(logger, cfg)
	if err != nil {
		return nil, err
	}

	// Load trackers from the index file,
	// and build a tracker for each unique URL, symbol
	indexers, symbolsForAPI, err := ParseApiFile(logger, cfg, db, client)
	if err != nil {
		return nil, err
	}

	var sortedIndexers []string
	// Set the reverse map.
	for api, symbols := range symbolsForAPI {
		indexers[api].Symbols = symbols
		sortedIndexers = append(sortedIndexers, api)
	}

	// Sort the Indexer array so we return the same order every time.
	sort.Strings(sortedIndexers)

	// Make an array of trackers to be sent to Runner.
	trackers := make([]*IndexTracker, len(indexers))
	for idx, api := range sortedIndexers {
		trackers[idx] = indexers[api]
	}

	// Start the PSR system that will feed from these indexes.
	err = InitPSRs()
	if err != nil {
		return nil, errors.Wrap(err, "initialize PSRs")
	}

	return trackers, nil
}

// ParseApiFile parses api.json file and returns a *IndexTracker,
// for every URL in index file, also a map[string][]string that describes which APIs
// influence which symbols.
func ParseApiFile(logger log.Logger, cfg *config.Config, DB db.DataServerProxy, client contracts.ETHClient) (trackersPerURL map[string]*IndexTracker, symbolsForAPI map[string][]string, err error) {

	// Load index file.
	byteValue, err := ioutil.ReadFile(cfg.ApiFile)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "read index file @ %s", cfg.ApiFile)
	}
	// Parse to json.
	baseIndexes := make(map[string][]IndexObject)
	err = json.Unmarshal(byteValue, &baseIndexes)
	if err != nil {
		return nil, nil, errors.Wrap(err, "parse index file")
	}
	// Keep track of tracker per symbol.
	indexes = make(map[string][]*IndexTracker)
	// Build a tracker for each unique URL.
	trackersPerURL = make(map[string]*IndexTracker)
	// Keep track of which APIs influence which symbols so we know what to update later.
	symbolsForAPI = make(map[string][]string)

	psr := NewPsr(
		logger,
		promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "values",
			Help:      "The tracker values",
		},
			[]string{"dataID"},
		),
	)

	for symbol, apis := range baseIndexes {
		for _, api := range apis {
			// Tracker for this API already added?
			_, ok := trackersPerURL[api.URL]
			if !ok {
				// Expand any env variables with their values from the .env file.
				_, err := godotenv.Read(cfg.EnvFile)
				// Ignore file doesn't exist errors.
				if err != nil && !os.IsNotExist(err) {
					return nil, nil, errors.Wrap(err, "reading .env file")
				}
				api.URL = os.Expand(api.URL, func(key string) string {
					return os.Getenv(key)
				})

				var name string
				var source DataSource

				// Create an index tracker based on the api type.
				// Default value for the api type.
				if api.Type == "" {
					api.Type = httpIndexType
				}
				switch api.Type {
				case httpIndexType:
					{
						source = NewJSONapi(logger, api.URL, cfg.Trackers.FetchTimeout.Duration)
						u, err := url.Parse(api.URL)
						if err != nil {
							return nil, nil, errors.Wrapf(err, "invalid API URL: %s", api.URL)
						}
						name = u.Host
					}
				case fileIndexType:
					{
						source = NewJSONfile(api.URL)
						name = filepath.Base(api.URL)
					}
				case ethereumIndexType:
					{
						// Getting current network id from geth node.
						networkID, err := client.NetworkID(context.Background())
						if err != nil {
							return nil, nil, err
						}
						// Validate and pick an ethereum address for current network id.
						address, err := util.GetAddressForNetwork(api.URL, networkID.Int64())
						if err != nil {
							return nil, nil, errors.Wrap(err, "getting address for network id")
						}
						if api.Parser == uniswapIndexParser {
							source = NewUniswap(symbol, address, client)

						} else if api.Parser == balancerIndexParser {
							source = NewBalancer(symbol, address, client)
						} else {
							return nil, nil, errors.Wrapf(err, "unknown source for on-chain index tracker")
						}
						name = fmt.Sprintf("%s(%s)", api.Type, api.URL)
					}
				default:
					return nil, nil, errors.New("unknown index type for index object")
				}

				if api.Interval.Duration > 0 && (api.Interval.Duration < cfg.Trackers.SleepCycle.Duration) {
					return nil, nil, errors.New("api interval can't be smaller than the global tracker cycle")
				}

				// Default value for the parser.
				if api.Parser == "" {
					api.Parser = jsonPathIndexParser
				}
				current := &IndexTracker{
					Name:       name,
					Identifier: api.URL,
					Source:     source,
					DB:         DB,
					Interval:   api.Interval.Duration,
					Param:      api.Param,
					Type:       api.Type,
					cfg:        cfg,
					psr:        psr,
				}

				trackersPerURL[api.URL] = current
			}
			// Now we definitely have one.
			thisOne := trackersPerURL[api.URL]

			// Insert add it and it's more specific variant to the symbol -> api map.
			indexes[symbol] = append(indexes[symbol], thisOne)
			specificName := fmt.Sprintf("%s~%s", symbol, thisOne.Name)
			indexes[specificName] = append(indexes[specificName], thisOne)

			// Save this for later so we can build the api->symbol map.
			symbolsForAPI[api.URL] = append(symbolsForAPI[api.URL], symbol, specificName)
		}
	}
	return

}

// IndexType -> index type for IndexObject.
type IndexType string

const (
	httpIndexType     IndexType = "http"
	ethereumIndexType IndexType = "ethereum"
	fileIndexType     IndexType = "file"
)

// IndexParser -> index parser for IndexObject.
type IndexParser string

const (
	jsonPathIndexParser IndexParser = "jsonPath"
	uniswapIndexParser  IndexParser = "Uniswap"
	balancerIndexParser IndexParser = "Balancer"
)

// IndexObject will be used in parsing index file.
type IndexObject struct {
	URL      string          `json:"URL"`
	Type     IndexType       `json:"type"`
	Parser   IndexParser     `json:"parser"`
	Param    string          `json:"param"`
	Interval config.Duration `json:"interval"`
}

type IndexTracker struct {
	DB               db.DataServerProxy
	Name             string
	Identifier       string
	Symbols          []string
	Source           DataSource
	Interval         time.Duration
	Param            string
	Type             IndexType
	lastRunTimestamp time.Time
	cfg              *config.Config
	psr              *PSR
}

type DataSource interface {
	Get(context.Context) ([]byte, error)
}

func NewJSONapi(logger log.Logger, url string, retryDelay time.Duration) *JSONapi {
	return &JSONapi{
		logger:     logger,
		url:        url,
		retryDelay: retryDelay,
	}
}

type JSONapi struct {
	url        string
	logger     log.Logger
	retryDelay time.Duration
}

func (self *JSONapi) Get(ctx context.Context) ([]byte, error) {
	return http.Fetch(ctx, self.logger, self.url, self.retryDelay)
}

func NewJSONfile(filepath string) *JSONfile {
	return &JSONfile{filepath: filepath}
}

type JSONfile struct {
	filepath string
}

func (j *JSONfile) Get(_ context.Context) ([]byte, error) {
	return ioutil.ReadFile(j.filepath)
}

func (i *IndexTracker) Exec(ctx context.Context) error {
	now := time.Now()
	if now.Sub(i.lastRunTimestamp) < i.Interval {
		return nil
	}
	i.lastRunTimestamp = now

	payload, err := i.Source.Get(ctx)
	if err != nil {
		return err
	}

	vals, err := i.ParsePayload(payload)
	if err != nil {
		return err
	}

	volume := 0.0
	if len(vals) >= 2 {
		volume = vals[1]
	}

	//save the value into our local data window (set 0 volume for now)
	apiOracle.SetRequestValue(i.Identifier, clck.Now(), apiOracle.PriceInfo{Price: vals[0], Volume: volume})
	//update all the values that depend on these symbols
	return i.psr.UpdatePSRs(ctx, i.cfg, i.DB, i.Symbols)
}

func (i *IndexTracker) String() string {
	return fmt.Sprintf("%s on %s", strings.Join(i.Symbols, ","), i.Name)
}

// ParsePayload parses the input JSON payload to a slice of float64
// The input JSON will get queried using JSONPath query language if
// the JSONPath expression is not empty.
func (i *IndexTracker) ParsePayload(payload []byte) (vals []float64, err error) {

	var decodedPayload, result interface{}
	err = json.Unmarshal(payload, &decodedPayload)
	if err != nil {
		return
	}

	// Query the json payload using JSONPath expression if needed.
	result = decodedPayload
	if len(strings.TrimSpace(i.Param)) > 0 {
		if err != nil {
			return
		}
		result, err = jsonpath.Read(decodedPayload, i.Param)
		if err != nil {
			return
		}
	}

	// Expect result to be a slice of float or a single float value.
	var resultList []interface{}
	switch result := result.(type) {
	case []interface{}:
		resultList = result
	default:
		resultList = []interface{}{result}
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
	return
}

// IndexProcessor consolidates the recorded API values to a single value.
type IndexProcessor func([]*IndexTracker, time.Time, float64) (apiOracle.PriceInfo, float64, error)

type ValueGenerator interface {
	// Require reports what a PSR requires to produce a value.
	Require() map[string]IndexProcessor

	// ValueAt returns the best estimate of a value at a given time, and the confidence
	// if confidence == 0, the value has no meaning
	ValueAt(map[string]apiOracle.PriceInfo) float64

	// Granularity returns the currency granularity.
	Granularity() int64

	// Symbol returns the tracker Symbol.
	Symbol() string
}

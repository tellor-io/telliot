package aggregator

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/tracker/index"
	"github.com/tellor-io/telliot/pkg/util"
)

type AggregatorRule struct {
	Symbol      string
	Transform   string
	Granularity int
}

const ComponentName = "aggregator"

var aggr = map[int]AggregatorRule{
	1: {
		Symbol:      "ETH/USD",
		Transform:   "MedianAt",
		Granularity: 1000000,
	},
}

type Config struct {
	LogLevel      string
	MinConfidence float64
	Interval      util.Duration
}

type PriceInfo struct {
	Price, Volume float64
}

type Aggregator struct {
	logger       log.Logger
	ctx          context.Context
	stop         context.CancelFunc
	tsDB         *tsdb.DB
	promqlEngine *promql.Engine
	cfg          Config
	prices       *prometheus.GaugeVec
}

func New(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	tsDB *tsdb.DB,
	client contracts.ETHClient,
) (*Aggregator, error) {

	ctx, stop := context.WithCancel(ctx)

	opts := promql.EngineOpts{
		Logger:     nil,
		Reg:        nil,
		MaxSamples: 10,
		Timeout:    10 * time.Second,
	}
	engine := promql.NewEngine(opts)

	return &Aggregator{
		logger:       log.With(logger, "component", ComponentName),
		ctx:          ctx,
		stop:         stop,
		tsDB:         tsDB,
		promqlEngine: engine,
		cfg:          cfg,
		prices: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "price",
			Help:      "The currency price",
		},
			[]string{"source"},
		),
	}, nil
}

func (self *Aggregator) getLatest(symbol string) ([]PriceInfo, float64, error) {

	// 100% confidence is when all apis have returned a value within the last 5 minutes.
	// For every minute above the 5min mark the calculation subtracts some confidence level.

	// Get the total number of APIs for this Symbol.
	query, err := self.promqlEngine.NewInstantQuery(self.tsDB, util.SanitizeMetricName(symbol+index.ApiCountSuffix), time.Now())
	if err != nil {
		return nil, 0, err
	}
	defer query.Close()
	apiCount := query.Exec(self.ctx)
	if apiCount.Err != nil {
		return nil, 0, errors.Wrapf(apiCount.Err, "error evaluating query:%v", query)
	}

	if apiCount.Value.String() == "" {
		return nil, 0, errors.New("query returns no results for API count")
	}

	// Get the timestamp of the last value.
	query, err = self.promqlEngine.NewInstantQuery(self.tsDB, util.SanitizeMetricName(symbol+index.TsLastValueSuffix), time.Now())
	if err != nil {
		return nil, 0, err
	}
	defer query.Close()
	_tsLastValue := query.Exec(self.ctx)
	if _tsLastValue.Err != nil {
		return nil, 0, errors.Wrapf(_tsLastValue.Err, "error evaluating query:%v", query)
	}

	if _tsLastValue.Value.String() == "" {
		return nil, 0, errors.New("query returns no results for last value TS")
	}

	// Get the last price for all APIs.
	tsLastValueSources, err := _tsLastValue.Vector()
	if err != nil {
		return nil, 0, errors.New("expected tsLastValue to be a vector")
	}

	for _, v := range tsLastValueSources {
		fmt.Println("point", v.Point)
		fmt.Println("metric", v.Metric)
	}

	// for tsLastValueSources

	// ts, err := strconv.ParseInt(_tsLastValue.Value.String(), 10, 64)
	// if err != nil {
	// 	return nil, 0, errors.New("parsing  _tsLastValue")
	// }
	// tsLastValue := time.Unix(ts, 0)
	// query, err = self.promqlEngine.NewInstantQuery(self.tsDB, util.SanitizeMetricName(symbol+index.PriceSuffix), tsLastValue)
	// if err != nil {
	// 	return nil, 0, err
	// }
	// defer query.Close()
	// price := query.Exec(self.ctx)
	// if price.Err != nil {
	// 	return nil, 0, errors.Wrapf(price.Err, "error evaluating query:%v", query)
	// }

	// if price.Value.String() == "" {
	// 	return nil, 0, errors.New("query returns no results for symbol price")
	// }

	return nil, 0, nil

	// var values []PriceInfo
	// totalConf := 0.0
	// for _, api := range apis {
	// 	b, _ := apiOracle.GetNearestTwoRequestValue(api.Identifier, at)
	// 	if b != nil {
	// 		// Penalize values more than 5 minutes old.
	// 		totalConf += math.Min(5/at.Sub(b.Created).Minutes(), 1.0)
	// 		values = append(values, b.PriceInfo)
	// 	}
	// }
	// return values, totalConf / float64(len(apis))
}

func (self *Aggregator) Run() error {
	ticker := time.NewTicker(self.cfg.Interval.Duration)
	for {
		if _, _, err := self.getLatest("ALGO/USD"); err != nil {
			level.Error(self.logger).Log("msg", "get latest", "err", err)
		}
		<-ticker.C

	}
}

func (self *Aggregator) Stop() {
	self.stop()
}

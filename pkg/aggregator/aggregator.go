package aggregator

import (
	"context"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/prometheus/pkg/timestamp"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/logging"
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
	LogLevel            string
	MinConfidence       float64
	ConfidIntvThreshold util.Duration
}

type PriceInfo struct {
	Price, Volume float64
}

type Aggregator struct {
	logger       log.Logger
	ctx          context.Context
	stop         context.CancelFunc
	tsDB         storage.SampleAndChunkQueryable
	promqlEngine *promql.Engine
	cfg          Config
	prices       *prometheus.GaugeVec
}

func New(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	tsDB storage.SampleAndChunkQueryable,
	client contracts.ETHClient,
) (*Aggregator, error) {

	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	ctx, stop := context.WithCancel(ctx)

	opts := promql.EngineOpts{
		Logger:        logger,
		Reg:           nil,
		MaxSamples:    10000,
		Timeout:       10 * time.Second,
		LookbackDelta: cfg.ConfidIntvThreshold.Duration, // Any value below this Duration is considered invalid.
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

func (self *Aggregator) Run() error {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-self.ctx.Done():
			return nil
		default:
		}
		price, volume, confidence, err := self.MedianAt("ETH/USD", time.Now())
		if err != nil {
			level.Error(self.logger).Log("msg", "get latest", "err", err)
			<-ticker.C
			continue
		}
		fmt.Println("price, volume, confidence", price, volume, confidence)
		select {
		case <-self.ctx.Done():
			return nil
		case <-ticker.C:
			continue
		}

	}
}

func (self *Aggregator) MedianAt(symbol string, at time.Time) (float64, float64, float64, error) {
	values, confidence, err := self.valuesAt(symbol, at)
	if err != nil {
		return 0, 0, 0, err
	}
	if len(values) == 0 {
		return 0, 0, 0, errors.New("no values")
	}
	price, volume := self.median(values)
	return price, volume, confidence, nil
}

func (self *Aggregator) MedianAtEOD(symbol string, at time.Time) (float64, float64, float64, error) {
	d := 24 * time.Hour
	eod := time.Now().Truncate(d)
	return self.MedianAt(symbol, eod)
}

func (self *Aggregator) MeanAt(symbol string, at time.Time) (float64, float64, float64, error) {
	values, confidence, err := self.valuesAt(symbol, at)
	if err != nil {
		return 0, 0, 0, err
	}
	price, volume := self.mean(values)
	return price, volume, confidence, nil
}

func (self *Aggregator) mean(vals []PriceInfo) (float64, float64) {
	priceSum := 0.0
	volSum := 0.0
	for _, val := range vals {
		priceSum += val.Price
		volSum += val.Volume
	}
	return priceSum / float64(len(vals)), volSum
}

func (self *Aggregator) TimeWeightedAvg(
	symbol string,
	from time.Time,
	lookBack time.Duration,
	resolution float64,
	weightFn func(float64) (float64, float64),
) (float64, float64, float64, error) {
	sum := 0.0
	weightSum := 0.0
	series, err := self.valuesForTime(symbol, from, lookBack)
	if err != nil {
		return 0, 0, 0, err
	}
	for _, ser := range series {
		for _, v := range ser.Points {
			normDelta := from.Sub(timestamp.Time(v.T)).Seconds() / lookBack.Seconds()
			weight, _ := weightFn(normDelta)
			sum += v.V * weight
			weightSum += weight
		}
	}

	// Number of APIs * rate * interval.
	apiCount, err := self.apiCount(symbol)
	if err != nil {
		return 0, 0, 0, err
	}
	maxWeight := apiCount * (1 / resolution) * lookBack.Seconds()
	// Average weight is the integral of the weight fn over [0,1].
	_, avgWeight := weightFn(0)
	targetWeight := maxWeight * avgWeight

	price := sum / weightSum

	// Use the highest volume seen over all values.
	// Works well when the time averaging window is equal to the interval of volume reporting
	// ie, 24 hour average on an api that returns 24hr volume
	volume, err := self.maxVolumeFor(symbol, from, lookBack)
	if err != nil {
		return 0, 0, 0, err
	}
	return price, volume, math.Min(weightSum/targetWeight, 1.0), nil
}

// ExpDecay maps values of x between 0 (brand new) and 1 (old) to weights between 0 and 1
// also returns the integral of the weight over the range [0,1]
// weights the oldest data (1) as being 1/3 as important (1/e).
func ExpDecay(x float64) (float64, float64) {
	return 1 / math.Exp(x), 0.63212
}

// NoDecay weights all data in the time interval evenly.
func NoDecay(x float64) (float64, float64) {
	return 1, 1
}

func (self *Aggregator) median(values []PriceInfo) (float64, float64) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].Price < values[j].Price
	})
	price := values[len(values)/2].Price
	volume := 0.0
	for _, price := range values {
		volume += price.Volume
	}
	return price, volume
}

func (self *Aggregator) Stop() {
	self.stop()
}

func (self *Aggregator) valuesForTime(symbol string, from time.Time, lookBack time.Duration) (promql.Matrix, error) {
	query, err := self.promqlEngine.NewRangeQuery(
		self.tsDB,
		util.SanitizeMetricName(symbol+index.PriceSuffix),
		from,
		from.Truncate(lookBack),
		time.Second,
	)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	result := query.Exec(self.ctx)
	if result.Err != nil {
		return nil, errors.Wrapf(result.Err, "error evaluating query:%v", query)
	}

	return result.Value.(promql.Matrix), nil

}

func (self *Aggregator) maxVolumeFor(symbol string, from time.Time, lookBack time.Duration) (float64, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		"max(max_over_time("+util.SanitizeMetricName(symbol+index.VolumeSuffix)+"["+lookBack.String()+"]))",
		from,
	)
	if err != nil {
		return 0, err
	}
	defer query.Close()
	result := query.Exec(self.ctx)
	if result.Err != nil {
		return 0, errors.Wrapf(result.Err, "error evaluating query:%v", query)
	}

	return result.Value.(promql.Vector)[0].V, nil

}

// valuesAt returns the value for a given symbol with the confidence level.
// 100% confidence is when all apis have returned a value within the last 5 minutes.
// For every minute above the 5min mark the calculation subtracts some confidence level.
func (self *Aggregator) valuesAt(symbol string, at time.Time) ([]PriceInfo, float64, error) {
	// Get the value of the last price for all APIs.
	var priceInfos []PriceInfo
	totalConf := 0.0

	prices, err := self.pricesAt(symbol, at)
	if err != nil {
		return nil, 0, err
	}
	volumes, err := self.volumesAt(symbol, at)
	if err != nil {
		return nil, 0, err
	}
	for _, price := range prices {
		volume := 0.0
		for _, _volume := range volumes {
			if price.Metric[1].Value == _volume.Metric[1].Value {
				volume = _volume.V
				break
			}
		}

		// Penalize values more than 5 minutes old.
		tsT := timestamp.Time(price.T)
		totalConf += math.Min(5/time.Since(tsT).Minutes(), 1.0)
		priceInfos = append(priceInfos, PriceInfo{
			Price:  price.V,
			Volume: volume,
		})
	}
	apiCount, err := self.apiCount(symbol)
	if err != nil {
		return nil, 0, err
	}

	return priceInfos, totalConf / apiCount, nil
}
func (self *Aggregator) volumesAt(symbol string, at time.Time) (promql.Vector, error) {
	query, err := self.promqlEngine.NewInstantQuery(self.tsDB, util.SanitizeMetricName(symbol+index.VolumeSuffix), at)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	result := query.Exec(self.ctx)
	if result.Err != nil {
		return nil, errors.Wrapf(result.Err, "error evaluating query:%v", query)
	}
	return result.Value.(promql.Vector), nil
}

func (self *Aggregator) pricesAt(symbol string, at time.Time) (promql.Vector, error) {
	query, err := self.promqlEngine.NewInstantQuery(self.tsDB, util.SanitizeMetricName(symbol+index.PriceSuffix), at)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	result := query.Exec(self.ctx)
	if result.Err != nil {
		return nil, errors.Wrapf(result.Err, "error evaluating query:%v", query)
	}

	return result.Value.(promql.Vector), nil
}

func (self *Aggregator) apiCount(symbol string) (float64, error) {
	query, err := self.promqlEngine.NewInstantQuery(self.tsDB, util.SanitizeMetricName(symbol+index.ApiCountSuffix), time.Now())
	if err != nil {
		return 0, err
	}
	defer query.Close()
	result := query.Exec(self.ctx)
	if result.Err != nil {
		return 0, errors.Wrapf(result.Err, "error evaluating query:%v", query)
	}

	if result.Value.String() == "" {
		return 0, errors.New("no results for API count")
	}

	return result.Value.(promql.Vector)[0].V, nil
}

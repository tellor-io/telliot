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
	LogLevel      string
	MinConfidence float64
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
		Logger:     logger,
		Reg:        nil,
		MaxSamples: 30000,
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

func (self *Aggregator) Run() error {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-self.ctx.Done():
			return nil
		default:
		}
		price, volume, confidence, err := self.TimeWeightedAvg("TRB/ETH", time.Now(), time.Hour)
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

// TimeWeightedAvg returns price,volume and confidence level for a given symbol.
// Confidence is calculated based on the total amount of data points and the total count of data sources.
// For example with 1h look back, resolution 1m and 10 data sources 100% confidence is
// when the total count of data points is 600 (60m/1m * 10).
// In other words total duration divided by resolution multiplied by sources count.
func (self *Aggregator) TimeWeightedAvg(
	symbol string,
	from time.Time,
	lookBack time.Duration,
) (float64, float64, float64, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg_over_time(`+index.PriceMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
		from,
	)
	if err != nil {
		return 0, 0, 0, err
	}
	defer query.Close()
	result := query.Exec(self.ctx)
	if result.Err != nil {
		return 0, 0, 0, errors.Wrapf(result.Err, "error evaluating query:%v", query)
	}
	fmt.Println("avg", result)

	query, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`count_over_time(`+index.PriceMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
		from,
	)
	if err != nil {
		return 0, 0, 0, err
	}
	defer query.Close()
	result = query.Exec(self.ctx)
	if result.Err != nil {
		return 0, 0, 0, errors.Wrapf(result.Err, "error evaluating query:%v", query)
	}

	fmt.Println("count", result)

	return 0, 0, 0, err

	// // Number of APIs * rate * interval.
	// apiCount, err := self.apiCount(symbol)
	// if err != nil {
	// 	return 0, 0, 0, err
	// }

	// // Use the highest volume seen over all values.
	// // Works well when the time averaging window is equal to the interval of volume reporting
	// // ie, 24 hour average on an api that returns 24hr volume
	// volume, err := self.maxVolumeFor(symbol, from, lookBack)
	// if err != nil {
	// 	return 0, 0, 0, err
	// }
	// return price, volume, math.Min(weightSum/targetWeight, 1.0), nil
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

func (self *Aggregator) valuesForTime(symbol string, from time.Time, lookBack time.Duration, resolution time.Duration) (promql.Matrix, error) {
	query, err := self.promqlEngine.NewRangeQuery(
		self.tsDB,
		util.SanitizeMetricName(symbol+index.PriceMetricName),
		from,
		from.Truncate(lookBack),
		resolution*time.Millisecond, // The parameter expects milliseconds.
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
		"max(max_over_time("+util.SanitizeMetricName(symbol+index.VolumeMetricName)+"["+lookBack.String()+"]))",
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
	query, err := self.promqlEngine.NewInstantQuery(self.tsDB, util.SanitizeMetricName(symbol+index.VolumeMetricName), at)
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
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		index.PriceMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}`,
		at,
	)
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
	query, err := self.promqlEngine.NewInstantQuery(self.tsDB, util.SanitizeMetricName(symbol+index.ApiCountMetricName), time.Now())
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

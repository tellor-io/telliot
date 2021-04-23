package aggregator

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
// Confidence is calculated based on maximum possible samples over the actual samples for a given period.
// avg(maxPossibleSamplesCount/actualSamplesCount)
// For example with 1h look back and source interval of 60sec maxPossibleSamplesCount = 36
// with actualSamplesCount = 18 this is 50% confidence.
// The same calculation is done for all source and the final value is the averege of all.
func (self *Aggregator) TimeWeightedAvg(
	symbol string,
	from time.Time,
	lookBack time.Duration,
) (float64, float64, error) {

	// Avg value.
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg_over_time(`+index.PriceMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
		from,
	)
	if err != nil {
		return 0, 0, err
	}
	defer query.Close()
	avg := query.Exec(self.ctx)
	if avg.Err != nil {
		return 0, 0, errors.Wrapf(avg.Err, "error evaluating query:%v", query)
	}

	// Confidence level.
	// an example for 1h period.
	// avg(count_over_time(indexTracker_prices{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval))
	query, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.PriceMetricName+`{
				symbol="`+util.SanitizeMetricName(symbol)+`"
			}[`+lookBack.String()+`]) /
			(`+strconv.Itoa(int(lookBack.Nanoseconds()))+` / indexTracker_interval))`,
		from,
	)
	if err != nil {
		return 0, 0, err
	}
	defer query.Close()
	confidence := query.Exec(self.ctx)
	if confidence.Err != nil {
		return 0, 0, errors.Wrapf(confidence.Err, "error evaluating query:%v", query)
	}

	return avg.Value.(promql.Vector)[0].V, confidence.Value.(promql.Vector)[0].V, err
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

// func (self *Aggregator) valuesForTime(symbol string, from time.Time, lookBack time.Duration, resolution time.Duration) (promql.Matrix, error) {
// 	query, err := self.promqlEngine.NewInstantQuery(
// 		self.tsDB,
// 		util.SanitizeMetricName(symbol+index.PriceMetricName),
// 		from,
// 		from.Add(-lookBack),
// 		resolution*time.Millisecond, // The parameter expects milliseconds.
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer query.Close()
// 	result := query.Exec(self.ctx)
// 	if result.Err != nil {
// 		return nil, errors.Wrapf(result.Err, "error evaluating query:%v", query)
// 	}

// 	return result.Value.(promql.Matrix), nil

// }

func (self *Aggregator) maxVolumeFor(symbol string, from time.Time, lookBack time.Duration) (float64, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`max_over_time(`+index.VolumeMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
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

// valuesAt returns the value from all sources for a given symbol with the confidence level.
// 100% confidence is when all apis have returned a value within the last 10 minutes.
// For every missing value the calculation subtracts some confidence level.
// Confidence is calculated totalValues/totalApiCount
func (self *Aggregator) valuesAt(symbol string, at time.Time) ([]PriceInfo, float64, error) {
	var priceInfos []PriceInfo
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
		for _, _volume := range volumes { // Find the volume for this symbol.
			if price.Metric[1].Value == _volume.Metric[1].Value {
				volume = _volume.V
				break
			}
		}
		priceInfos = append(priceInfos, PriceInfo{
			Price:  price.V,
			Volume: volume,
		})
	}
	apiCount, err := self.apiCount(symbol)
	if err != nil {
		return nil, 0, err
	}

	return priceInfos, float64(len(prices)) / apiCount, nil
}
func (self *Aggregator) volumesAt(symbol string, at time.Time) (promql.Vector, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`last_over_time{`+index.VolumeMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"})[10m]`,
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

func (self *Aggregator) pricesAt(symbol string, at time.Time) (promql.Vector, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`last_over_time(`+index.PriceMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}[10m])`,
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

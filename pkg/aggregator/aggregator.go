package aggregator

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
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

type Aggregator struct {
	logger       log.Logger
	ctx          context.Context
	tsDB         storage.SampleAndChunkQueryable
	promqlEngine *promql.Engine
	maxLookback  time.Duration
	cfg          Config
	value        *prometheus.GaugeVec
	confidence   *prometheus.GaugeVec
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

	maxLookback := 5 * time.Minute
	opts := promql.EngineOpts{
		Logger:               logger,
		Reg:                  nil,
		MaxSamples:           30000,
		Timeout:              10 * time.Second,
		LookbackDelta:        maxLookback,
		EnableAtModifier:     true,
		EnableNegativeOffset: true,
	}
	engine := promql.NewEngine(opts)

	return &Aggregator{
		logger:       log.With(logger, "component", ComponentName),
		ctx:          ctx,
		tsDB:         tsDB,
		promqlEngine: engine,
		cfg:          cfg,
		maxLookback:  maxLookback,
		value: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "value",
			Help:      "The aggregated value",
		},
			[]string{"symbol", "type", "interval"},
		),
		confidence: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "confidence",
			Help:      "The aggregated value confidence",
		},
			[]string{"symbol", "type", "interval"},
		),
	}, nil
}

func (self *Aggregator) MedianAt(symbol string, at time.Time) (float64, float64, error) {
	values, confidence, err := self.valuesAtWithConfidence(symbol, at)
	if err != nil {
		return 0, 0, err
	}
	if len(values) == 0 {
		return 0, 0, errors.New("no values")
	}
	price := self.median(values)
	return price, confidence, nil
}

func (self *Aggregator) MedianAtEOD(symbol string, at time.Time) (float64, float64, error) {
	d := 24 * time.Hour
	eod := time.Now().Truncate(d)
	return self.MedianAt(symbol, eod)
}

func (self *Aggregator) MeanAt(symbol string, at time.Time) (float64, float64, error) {
	values, confidence, err := self.valuesAtWithConfidence(symbol, at)
	if err != nil {
		return 0, 0, err
	}
	price := self.mean(values)
	return price, confidence, nil
}

func (self *Aggregator) mean(vals []float64) float64 {
	priceSum := 0.0
	for _, val := range vals {
		priceSum += val
	}
	return priceSum / float64(len(vals))
}

// TimeWeightedAvg returns price,volume and confidence level for a given symbol.
// Confidence is calculated based on maximum possible samples over the actual samples for a given period.
// avg(maxPossibleSamplesCount/actualSamplesCount)
// For example with 1h look back and source interval of 60sec maxPossibleSamplesCount = 36
// with actualSamplesCount = 18 this is 50% confidence.
// The same calculation is done for all source and the final value is the average of all.
// Example for 1h.
// maxDataPointCount is calculated by deviding the seconds in 1h by how often the tracker queries the APIs.
// avg(count_over_time(indexTracker_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval)).
func (self *Aggregator) TimeWeightedAvg(
	symbol string,
	from time.Time,
	lookBack time.Duration,
) (float64, float64, error) {
	// Avg value over the look back period.
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg_over_time(`+index.ValueMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
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
	query, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
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

	if len(avg.Value.(promql.Vector)) == 0 {
		return 0, 0, errors.New("no result")
	}

	return avg.Value.(promql.Vector)[0].V, confidence.Value.(promql.Vector)[0].V, err
}

// VolumWeightedAvg returns price,volume and confidence level for a given symbol.
// Confidence is calculated based on maximum possible samples over the actual samples for a given period.
// avg(maxPossibleSamplesCount/actualSamplesCount)
// For example with 1h look back and source interval of 60sec maxPossibleSamplesCount = 36
// with actualSamplesCount = 18 this is 50% confidence.
// The same calculation is done for all source and the final value is the average of all.
// Example for 1h.
// maxDataPointCount is calculated by deviding the seconds in 1h by how often the tracker queries the APIs.
// avg(count_over_time(indexTracker_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval)).
func (self *Aggregator) VolumWeightedAvg(
	symbol string,
	start time.Time,
	end time.Time,
	aggrWindow time.Duration,
) (float64, float64, error) {
	_timeWindow := end.Sub(start).Round(time.Minute).Seconds()
	timeWindow := strconv.Itoa(int(_timeWindow)) + "s"

	q, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			sum_over_time(
				(
					sum_over_time(indexTracker_value{symbol="`+util.SanitizeMetricName(symbol)+`_VOLUME"}[`+aggrWindow.String()+`])
					*  on(domain)
					avg_over_time(indexTracker_value{symbol="`+util.SanitizeMetricName(symbol)+`"}[`+aggrWindow.String()+`])
				)[`+timeWindow+`:`+aggrWindow.String()+`])
			/ on(domain)
			sum_over_time(
				(
					sum_over_time(indexTracker_value{symbol="`+util.SanitizeMetricName(symbol)+`_VOLUME"}[`+aggrWindow.String()+`]
				)
			)[`+timeWindow+`:`+aggrWindow.String()+`])
		)`,
		end,
	)
	if err != nil {
		return 0, 0, err
	}
	defer q.Close()

	_result := q.Exec(self.ctx)
	if _result.Err != nil {
		return 0, 0, errors.Wrapf(_result.Err, "error evaluating query:%v", q)
	}
	result := _result.Value.(promql.Vector)
	if len(result) == 0 {
		return 0, 0, errors.New("no result for values")
	}

	// Confidence level for prices and volumes.
	// an example for 1h period.
	// confidence = actualDataPointCountFor1h/maxDataPointCountFor1h
	// avg(count_over_time(indexTracker_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval))

	// For prices.
	q, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+util.SanitizeMetricName(symbol)+`"
			}[`+timeWindow+`]) /
			(`+strconv.Itoa(int(end.Sub(start).Nanoseconds()))+` / indexTracker_interval))`,
		end,
	)
	if err != nil {
		return 0, 0, err
	}
	defer q.Close()

	confidenceP := q.Exec(self.ctx)
	if confidenceP.Err != nil {
		return 0, 0, errors.Wrapf(confidenceP.Err, "error evaluating query:%v", q)
	}

	// For volumes.
	q, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+util.SanitizeMetricName(symbol)+`_VOLUME"
			}[`+timeWindow+`]) /
			(`+strconv.Itoa(int(end.Sub(start).Nanoseconds()))+` / indexTracker_interval))`,
		end,
	)
	if err != nil {
		return 0, 0, err
	}
	defer q.Close()
	confidenceV := q.Exec(self.ctx)
	if confidenceV.Err != nil {
		return 0, 0, errors.Wrapf(confidenceV.Err, "error evaluating query:%v", q)
	}

	if len(confidenceP.Value.(promql.Vector)) == 0 || len(confidenceV.Value.(promql.Vector)) == 0 {
		return 0, 0, errors.New("no result for confidence")
	}

	// Use the smaller confidence.
	confidence := confidenceP.Value.(promql.Vector)[0].V
	if confidence > confidenceV.Value.(promql.Vector)[0].V {
		confidence = confidenceV.Value.(promql.Vector)[0].V
	}

	self.value.With(
		prometheus.Labels{
			"type":     "vwap",
			"interval": start.Sub(end).String(),
			"symbol":   util.SanitizeMetricName(symbol),
		},
	).(prometheus.Gauge).Set(result[0].V)

	self.confidence.With(
		prometheus.Labels{
			"type":     "vwap",
			"interval": start.Sub(end).String(),
			"symbol":   util.SanitizeMetricName(symbol),
		},
	).(prometheus.Gauge).Set(confidence)

	return result[0].V, confidence, nil
}

func (self *Aggregator) median(values []float64) float64 {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	price := values[len(values)/2]
	return price
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

// func (self *Aggregator) maxVolumeFor(symbol string, from time.Time, lookBack time.Duration) (float64, error) {
// 	query, err := self.promqlEngine.NewInstantQuery(
// 		self.tsDB,
// 		`max_over_time(`+index.VolumeMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
// 		from,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer query.Close()
// 	result := query.Exec(self.ctx)
// 	if result.Err != nil {
// 		return 0, errors.Wrapf(result.Err, "error evaluating query:%v", query)
// 	}

// 	return result.Value.(promql.Vector)[0].V, nil

// }

// valuesAt returns the value from all sources for a given symbol with the confidence level.
// 100% confidence is when all apis have returned a value within the last 10 minutes.
// For every missing value the calculation subtracts some confidence level.
// Confidence is calculated actualDataPointCount/maxDataPointCount.
func (self *Aggregator) valuesAtWithConfidence(symbol string, at time.Time) ([]float64, float64, error) {
	var prices []float64
	pricesVector, err := self.valuesAt(symbol, at)
	if err != nil {
		return nil, 0, err
	}

	for _, price := range pricesVector {
		prices = append(prices, price.V)

	}

	// Confidence level.
	// an example for 1h period.
	// confidence = actualDataPointCountFor1h/maxDataPointCountFor1h
	// avg(count_over_time(indexTracker_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval))
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+util.SanitizeMetricName(symbol)+`"
			}[`+self.maxLookback.String()+`]) /
			(`+strconv.Itoa(int(self.maxLookback.Nanoseconds()))+` / indexTracker_interval))`,
		time.Now(),
	)
	if err != nil {
		return nil, 0, err
	}
	defer query.Close()
	confidence := query.Exec(self.ctx)
	if confidence.Err != nil {
		return nil, 0, errors.Wrapf(confidence.Err, "error evaluating query:%v", query)
	}

	return prices, confidence.Value.(promql.Vector)[0].V, nil
}

// func (self *Aggregator) volumesAt(symbol string, at time.Time) (promql.Vector, error) {
// 	query, err := self.promqlEngine.NewInstantQuery(
// 		self.tsDB,
// 		`last_over_time{`+index.VolumeMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"})[10m]`,
// 		at,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer query.Close()
// 	result := query.Exec(self.ctx)
// 	if result.Err != nil {
// 		return nil, errors.Wrapf(result.Err, "error evaluating query:%v", query)
// 	}
// 	return result.Value.(promql.Vector), nil
// }

func (self *Aggregator) valuesAt(symbol string, at time.Time) (promql.Vector, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`last_over_time(`+index.ValueMetricName+`{symbol="`+util.SanitizeMetricName(symbol)+`"}[`+self.maxLookback.String()+`])`,
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

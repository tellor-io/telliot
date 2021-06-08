// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package aggregator

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"github.com/tellor-io/telliot/pkg/format"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/tracker/index"
)

const ComponentName = "aggregator"

type IAggregator interface {
	TimeWeightedAvg(symbol string, start time.Time, lookBack time.Duration) (float64, float64, error)
}

type Config struct {
	LogLevel       string
	ManualDataFile string
}

type Aggregator struct {
	logger       log.Logger
	ctx          context.Context
	tsDB         storage.SampleAndChunkQueryable
	promqlEngine *promql.Engine
	cfg          Config
}

func New(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	tsDB storage.SampleAndChunkQueryable,
) (*Aggregator, error) {

	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	opts := promql.EngineOpts{
		Logger:               logger,
		Reg:                  nil,
		MaxSamples:           30000,
		Timeout:              10 * time.Second,
		LookbackDelta:        5 * time.Minute,
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
	}, nil
}

func (self *Aggregator) ManualValue(oracleName string, reqID int64, ts time.Time) (float64, error) {
	jsonFile, err := os.Open(self.cfg.ManualDataFile)
	if err != nil {
		return 0, errors.Wrapf(err, "manual data file read Error")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]map[string]map[string]float64
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return 0, errors.Wrap(err, "unmarshal manual data file")
	}

	oracleManualVals, ok := result[oracleName]
	if !ok {
		return 0, errors.Wrapf(err, "malformatted json file for oracle:%v", oracleName)
	}
	val := oracleManualVals[strconv.FormatInt(reqID, 10)]["VALUE"]
	if val != 0 {
		_timestamp := int64(oracleManualVals[strconv.FormatInt(reqID, 10)]["DATE"])
		timestamp := time.Unix(_timestamp, 0)
		if ts.After(timestamp) {
			return 0, errors.Errorf("manual entry value has expired:%v", ts)
		}
	}
	return val, nil
}

func (self *Aggregator) MedianAt(symbol string, at time.Time) (float64, float64, error) {
	values, confidence, err := self.valuesAtWithConfidence(symbol, at)
	if err != nil {
		return 0, 0, err
	}
	if len(values) == 0 {
		return 0, 0, errors.Errorf("no values at:%v", at)
	}
	median := self.median(values)
	return median, confidence, nil
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

// TimeWeightedAvg returns price and confidence level for a given symbol.
// Confidence is calculated based on maximum possible samples over the actual samples for a given period.
// avg(maxPossibleSamplesCount/actualSamplesCount)
// For example with 1h look back and source interval of 60sec maxPossibleSamplesCount = 36
// with actualSamplesCount = 18 this is 50% confidence.
// The same calculation is done for all source and the final value is the average of all.
//
// Example for 1h.
// maxDataPointCount is calculated by deviding the seconds in 1h by how often the tracker queries the APIs.
// avg(count_over_time(indexTracker_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval)).
func (self *Aggregator) TimeWeightedAvg(
	symbol string,
	start time.Time,
	lookBack time.Duration,
) (float64, float64, error) {
	// Avg value over the look back period.
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg_over_time(`+index.ValueMetricName+`{symbol="`+format.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
		start,
	)
	if err != nil {
		return 0, 0, err
	}
	defer query.Close()
	_result := query.Exec(self.ctx)
	if _result.Err != nil {
		return 0, 0, errors.Wrapf(_result.Err, "error evaluating query:%v", query.Statement())
	}
	if len(_result.Value.(promql.Vector)) == 0 {
		return 0, 0, errors.Errorf("no result for values query:%v", query.Statement())
	}

	result := _result.Value.(promql.Vector)[0].V

	// Confidence level.
	query, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+format.SanitizeMetricName(symbol)+`"
			}[`+lookBack.String()+`]) /
			(`+strconv.Itoa(int(lookBack.Nanoseconds()))+` / `+index.IntervalMetricName+`))`,
		start,
	)
	if err != nil {
		return 0, 0, err
	}
	defer query.Close()
	confidence := query.Exec(self.ctx)
	if confidence.Err != nil {
		return 0, 0, errors.Wrapf(confidence.Err, "error evaluating query:%v", query.Statement())
	}

	if len(confidence.Value.(promql.Vector)) == 0 {
		return 0, 0, errors.Errorf("no result for confidence query:%v", query.Statement())
	}

	return result, confidence.Value.(promql.Vector)[0].V, err
}

// VolumWeightedAvg returns price and confidence level for a given symbol.
// Confidence is calculated based on maximum possible samples over the actual samples for a given period.
// avg(maxPossibleSamplesCount/actualSamplesCount)
// For example with 1h look back and source interval of 60sec maxPossibleSamplesCount = 36
// with actualSamplesCount = 18 this is 50% confidence.
// The same calculation is done for all source and the final value is the average of all.
// maxDataPointCount is calculated by deviding the seconds in 1h by how often the tracker queries the APIs.
//
// Example for 1h.
// avg(count_over_time(indexTracker_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval)).
func (self *Aggregator) VolumWeightedAvg(
	symbol string,
	start time.Time,
	end time.Time,
	aggrWindow time.Duration,
) (float64, float64, error) {
	_timeWindow := end.Sub(start).Round(time.Minute).Seconds()
	timeWindow := strconv.Itoa(int(_timeWindow)) + "s"

	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			sum_over_time(
				(
					sum_over_time(`+index.ValueMetricName+`{symbol="`+format.SanitizeMetricName(symbol)+`_VOLUME"}[`+aggrWindow.String()+`])
					*  on(domain)
					avg_over_time(`+index.ValueMetricName+`{symbol="`+format.SanitizeMetricName(symbol)+`"}[`+aggrWindow.String()+`])
				)[`+timeWindow+`:`+aggrWindow.String()+`])
			/ on(domain)
			sum_over_time(
				(
					sum_over_time(`+index.ValueMetricName+`{symbol="`+format.SanitizeMetricName(symbol)+`_VOLUME"}[`+aggrWindow.String()+`]
				)
			)[`+timeWindow+`:`+aggrWindow.String()+`])
		)`,
		end,
	)
	if err != nil {
		return 0, 0, err
	}
	defer query.Close()

	_result := query.Exec(self.ctx)
	if _result.Err != nil {
		return 0, 0, errors.Wrapf(_result.Err, "error evaluating query:%v", query.Statement())
	}
	result := _result.Value.(promql.Vector)
	if len(result) == 0 {
		return 0, 0, errors.Errorf("no result for values query:%v", query.Statement())
	}

	// Confidence level for prices.
	query, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+format.SanitizeMetricName(symbol)+`"
			}[`+timeWindow+`]) /
			(`+strconv.Itoa(int(end.Sub(start).Nanoseconds()))+` / `+index.IntervalMetricName+`))`,
		end,
	)
	if err != nil {
		return 0, 0, err
	}
	defer query.Close()

	confidenceP := query.Exec(self.ctx)
	if confidenceP.Err != nil {
		return 0, 0, errors.Wrapf(confidenceP.Err, "error evaluating query:%v", query.Statement())
	}

	// Confidence level for volumes.
	query, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+format.SanitizeMetricName(symbol)+`_VOLUME"
			}[`+timeWindow+`]) /
			(`+strconv.Itoa(int(end.Sub(start).Nanoseconds()))+` / `+index.IntervalMetricName+`))`,
		end,
	)
	if err != nil {
		return 0, 0, err
	}
	defer query.Close()
	confidenceV := query.Exec(self.ctx)
	if confidenceV.Err != nil {
		return 0, 0, errors.Wrapf(confidenceV.Err, "error evaluating query:%v", query.Statement())
	}

	if len(confidenceP.Value.(promql.Vector)) == 0 || len(confidenceV.Value.(promql.Vector)) == 0 {
		return 0, 0, errors.Errorf("no result for confidence query:%v", query.Statement())

	}

	// Use the smaller confidence of volume or value.
	confidence := confidenceP.Value.(promql.Vector)[0].V
	if confidence > confidenceV.Value.(promql.Vector)[0].V {
		confidence = confidenceV.Value.(promql.Vector)[0].V
	}

	return result[0].V, confidence, nil
}

func (self *Aggregator) median(values []float64) float64 {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	price := values[len(values)/2]
	return price
}

// valuesAt returns the value from all sources for a given symbol with the confidence level.
// 100% confidence is when all apis have returned a value within the last tracker interval.
// For every missing value the calculation subtracts some confidence level.
// Confidence is calculated actualDataPointCount/maxDataPointCount.
//
// Example for 1h.
// avg(count_over_time(indexTracker_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval)).
func (self *Aggregator) valuesAtWithConfidence(symbol string, at time.Time) ([]float64, float64, error) {
	// Get the tracker interval.
	// It is the same for all endpoints of a given symbol.
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`last_over_time(`+index.IntervalMetricName+`{symbol="`+format.SanitizeMetricName(symbol)+`"}[3h])`, // No tracker should query slower then this.
		at,
	)
	if err != nil {
		return nil, 0, err
	}
	defer query.Close()
	trackerInterval := query.Exec(self.ctx)
	if trackerInterval.Err != nil {
		return nil, 0, errors.Wrapf(trackerInterval.Err, "error evaluating query:%v", query.Statement())
	}
	if len(trackerInterval.Value.(promql.Vector)) == 0 {
		return nil, 0, errors.Errorf("no values for tracker interval at:%v, query:%v", at, query.Statement())
	}

	lookBack := time.Duration(trackerInterval.Value.(promql.Vector)[0].V + 1e+9) // Pull interval + 1 sec to avoid races. Interval is in nanosecond granularity
	var prices []float64
	pricesVector, err := self.valuesAt(symbol, at, lookBack)
	if err != nil {
		return nil, 0, err
	}

	for _, price := range pricesVector {
		prices = append(prices, price.V)
	}

	// Confidence level.
	query, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+format.SanitizeMetricName(symbol)+`"
			}[`+lookBack.String()+`]) /
			(`+strconv.Itoa(int(lookBack.Nanoseconds()))+` / last_over_time(`+index.IntervalMetricName+`[3h])))`,
		at,
	)
	if err != nil {
		return nil, 0, err
	}
	defer query.Close()
	confidence := query.Exec(self.ctx)
	if confidence.Err != nil {
		return nil, 0, errors.Wrapf(confidence.Err, "error evaluating query:%v", query.Statement())
	}
	if len(confidence.Value.(promql.Vector)) == 0 {
		return nil, 0, errors.Errorf("no values for confidence at:%v, query:%v", at, query.Statement())
	}

	return prices, confidence.Value.(promql.Vector)[0].V, nil
}

func (self *Aggregator) valuesAt(symbol string, at time.Time, lookBack time.Duration) (promql.Vector, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`last_over_time(`+index.ValueMetricName+`{symbol="`+format.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
		at,
	)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	result := query.Exec(self.ctx)
	if result.Err != nil {
		return nil, errors.Wrapf(result.Err, "error evaluating query:%v", query.Statement())
	}

	return result.Value.(promql.Vector), nil
}

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
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/format"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/tracker/index"
)

const ComponentName = "aggregator"

type Config struct {
	LogLevel       string
	MinConfidence  float64
	ManualDataFile string
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

const (
	DefaultGranularity = 1000000
)

func (self *Aggregator) GetValueForIDWithDefaultGranularity(reqID int64, ts time.Time) (float64, error) {
	val, err := self.GetValueForID(reqID, ts)
	return val * DefaultGranularity, err
}

func (self *Aggregator) GetValueForID(reqID int64, ts time.Time) (float64, error) {
	val, err := self.getManualValue(reqID, ts)
	if err != nil {
		level.Error(self.logger).Log("msg", "get manual value", "reqID", reqID, "err", err)
	}
	if val != 0 {
		level.Warn(self.logger).Log("msg", "USING MANUAL VALUE", "reqID", reqID, "val", val)
		return val, nil
	}

	var conf float64
	switch reqID {
	case 1:
		val, conf, err = self.MedianAt("ETH/USD", ts)
	case 2:
		val, conf, err = self.MedianAt("BTC/USD", ts)
	case 3:
		val, conf, err = self.MedianAt("BNB/USD", ts)
	case 4:
		val, conf, err = self.TimeWeightedAvg("BTC/USD", ts, 24*time.Hour)
	case 5:
		val, conf, err = self.MedianAt("ETH/BTC", ts)
	case 6:
		val, conf, err = self.MedianAt("BNB/BTC", ts)
	case 7:
		val, conf, err = self.MedianAt("BNB/ETH", ts)
	case 8:
		val, conf, err = self.TimeWeightedAvg("ETH/USD", ts, 24*time.Hour)
	case 9:
		val, conf, err = self.MedianAtEOD("ETH/USD", ts)
	case 10: // For more details see https://docs.google.com/document/d/1RFCApk1PznMhSRVhiyFl_vBDPA4mP2n1dTmfqjvuTNw/edit
		val, conf, err = self.VolumWeightedAvg("AMPL/USD", time.Now().Add(-(24 * time.Hour)), time.Now(), 10*time.Minute)
	case 11:
		val, conf, err = self.MedianAt("ZEC/ETH", ts)
	case 12:
		val, conf, err = self.MedianAt("TRX/ETH", ts)
	case 13:
		val, conf, err = self.MedianAt("XRP/USD", ts)
	case 14:
		val, conf, err = self.MedianAt("XMR/ETH", ts)
	case 15:
		val, conf, err = self.MedianAt("ATOM/USD", ts)
	case 16:
		val, conf, err = self.MedianAt("LTC/USD", ts)
	case 17:
		val, conf, err = self.MedianAt("WAVES/BTC", ts)
	case 18:
		val, conf, err = self.MedianAt("REP/BTC", ts)
	case 19:
		val, conf, err = self.MedianAt("TUSD/ETH", ts)
	case 20:
		val, conf, err = self.MedianAt("EOS/USD", ts)
	case 21:
		val, conf, err = self.MedianAt("IOTA/USD", ts)
	case 22:
		val, conf, err = self.MedianAt("ETC/USD", ts)
	case 23:
		val, conf, err = self.MedianAt("ETH/PAX", ts)
	case 24:
		val, conf, err = self.TimeWeightedAvg("ETH/BTC", ts, 1*time.Hour)
	case 25:
		val, conf, err = self.MedianAt("USDC/USDT", ts)
	case 26:
		val, conf, err = self.MedianAt("XTZ/USD", ts)
	case 27:
		val, conf, err = self.MedianAt("LINK/USD", ts)
	case 28:
		val, conf, err = self.MedianAt("ZRX/BNB", ts)
	case 29:
		val, conf, err = self.MedianAt("ZEC/USD", ts)
	case 30:
		val, conf, err = self.MedianAt("XAU/USD", ts)
	case 31:
		val, conf, err = self.MedianAt("MATIC/USD", ts)
	case 32:
		val, conf, err = self.MedianAt("BAT/USD", ts)
	case 33:
		val, conf, err = self.MedianAt("ALGO/USD", ts)
	case 34:
		val, conf, err = self.MedianAt("ZRX/USD", ts)
	case 35:
		val, conf, err = self.MedianAt("COS/USD", ts)
	case 36:
		val, conf, err = self.MedianAt("BCH/USD", ts)
	case 37:
		val, conf, err = self.MedianAt("REP/USD", ts)
	case 38:
		val, conf, err = self.MedianAt("GNO/USD", ts)
	case 39:
		val, conf, err = self.MedianAt("DAI/USD", ts)
	case 40:
		val, conf, err = self.MedianAt("STEEM/BTC", ts)
	case 41:
		// ID 41 is always manual so it sholud never get here.
		// It is three month average for US PCE (monthly levels): https://www.bea.gov/data/personal-consumption-expenditures-price-index-excluding-food-and-energy
		return 0, errors.New("no manual entry for request ID 41")
	case 42:
		val, conf, err = self.MedianAtEOD("BTC/USD", ts)
	case 43:
		val, conf, err = self.MedianAt("TRB/ETH", ts)
	case 44:
		val, conf, err = self.TimeWeightedAvg("BTC/USD", ts, 1*time.Hour)
	case 45:
		val, conf, err = self.MedianAtEOD("TRB/USD", ts)
	case 46:
		val, conf, err = self.TimeWeightedAvg("ETH/USD", ts, 1*time.Hour)
	case 47:
		val, conf, err = self.MedianAt("BSV/USD", ts)
	case 48:
		val, conf, err = self.MedianAt("MAKER/USD", ts)
	case 49:
		val, conf, err = self.TimeWeightedAvg("BCH/USD", ts, 24*time.Hour)
	case 50:
		val, conf, err = self.MedianAt("TRB/USD", ts)
	case 51:
		val, conf, err = self.MedianAt("XMR/USD", ts)
	case 52:
		val, conf, err = self.MedianAt("XFT/USD", ts)
	case 53:
		val, conf, err = self.MedianAt("BTCDOMINANCE", ts)
	case 54:
		val, conf, err = self.MedianAt("WAVES/USD", ts)
	case 55:
		val, conf, err = self.MedianAt("OGN/USD", ts)
	case 56:
		val, conf, err = self.MedianAt("VIXEOD", ts)
	case 57:
		val, conf, err = self.MeanAt("DEFITVL", ts)
	case 58:
		val, conf, err = self.MeanAt("DEFIMCAP", ts)
	default:
		return 0, errors.Errorf("undeclared request ID:%v", reqID)
	}

	if conf < self.cfg.MinConfidence {
		return 0, errors.Errorf("not enough confidence - value:%v, conf:%v,confidence threshold:%v", val, conf, self.cfg.MinConfidence)
	}

	return val, err
}

func (self *Aggregator) getManualValue(reqID int64, ts time.Time) (float64, error) {
	jsonFile, err := os.Open(self.cfg.ManualDataFile)
	if err != nil {
		return 0, errors.Wrapf(err, "manual data file read Error")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]map[string]float64
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return 0, errors.Wrap(err, "unmarshal manual data file")
	}

	val := result[strconv.FormatInt(reqID, 10)]["VALUE"]
	if val != 0 {
		_timestamp := int64(result[strconv.FormatInt(reqID, 10)]["DATE"])
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
		`avg_over_time(`+index.ValueMetricName+`{symbol="`+format.SanitizeMetricName(symbol)+`"}[`+lookBack.String()+`])`,
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
				symbol="`+format.SanitizeMetricName(symbol)+`"
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
					sum_over_time(indexTracker_value{symbol="`+format.SanitizeMetricName(symbol)+`_VOLUME"}[`+aggrWindow.String()+`])
					*  on(domain)
					avg_over_time(indexTracker_value{symbol="`+format.SanitizeMetricName(symbol)+`"}[`+aggrWindow.String()+`])
				)[`+timeWindow+`:`+aggrWindow.String()+`])
			/ on(domain)
			sum_over_time(
				(
					sum_over_time(indexTracker_value{symbol="`+format.SanitizeMetricName(symbol)+`_VOLUME"}[`+aggrWindow.String()+`]
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

	// Confidence level for prices.
	q, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+format.SanitizeMetricName(symbol)+`"
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

	// Confidence level for volumes.
	q, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+format.SanitizeMetricName(symbol)+`_VOLUME"
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
			"symbol":   format.SanitizeMetricName(symbol),
		},
	).(prometheus.Gauge).Set(result[0].V)

	self.confidence.With(
		prometheus.Labels{
			"type":     "vwap",
			"interval": start.Sub(end).String(),
			"symbol":   format.SanitizeMetricName(symbol),
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

// valuesAt returns the value from all sources for a given symbol with the confidence level.
// 100% confidence is when all apis have returned a value within the last 10 minutes.
// For every missing value the calculation subtracts some confidence level.
// Confidence is calculated actualDataPointCount/maxDataPointCount.
// avg(count_over_time(indexTracker_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/indexTracker_interval))
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
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`avg(
			count_over_time(`+index.ValueMetricName+`{
				symbol="`+format.SanitizeMetricName(symbol)+`"
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

func (self *Aggregator) valuesAt(symbol string, at time.Time) (promql.Vector, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`last_over_time(`+index.ValueMetricName+`{symbol="`+format.SanitizeMetricName(symbol)+`"}[`+self.maxLookback.String()+`])`,
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

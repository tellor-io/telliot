// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package api

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
	"github.com/prometheus/common/route"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/timestamp"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/util/httputil"
	"github.com/prometheus/prometheus/util/stats"
)

type status string

const (
	statusSuccess status = "success"
	statusError   status = "error"
)

type errorType string

const (
	errorNone        errorType = ""
	errorTimeout     errorType = "timeout"
	errorCanceled    errorType = "canceled"
	errorExec        errorType = "execution"
	errorBadData     errorType = "bad_data"
	errorInternal    errorType = "internal"
	errorUnavailable errorType = "unavailable"
	errorNotFound    errorType = "not_found"
)

var (
	LocalhostRepresentations = []string{"127.0.0.1", "localhost", "::1"}
)

type apiError struct {
	typ errorType
	err error
}

func (e *apiError) Error() string {
	return fmt.Sprintf("%s: %s", e.typ, e.err)
}

type response struct {
	Status    status      `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	ErrorType errorType   `json:"errorType,omitempty"`
	Error     string      `json:"error,omitempty"`
	Warnings  []string    `json:"warnings,omitempty"`
}

type apiFuncResult struct {
	data      interface{}
	err       *apiError
	warnings  storage.Warnings
	finalizer func()
}

type apiFunc func(r *http.Request) apiFuncResult

// API can register a set of endpoints in a router and handle
// them using the provided storage and query engine.
type API struct {
	Queryable   storage.SampleAndChunkQueryable
	QueryEngine *promql.Engine
	now         func() time.Time
	logger      log.Logger
	CORSOrigin  *regexp.Regexp
}

func init() {
	jsoniter.RegisterTypeEncoderFunc("promql.Point", marshalPointJSON, marshalPointJSONIsEmpty)
}

// New returns an initialized API type.
func New(
	qe *promql.Engine,
	q storage.SampleAndChunkQueryable,
	ap storage.Appendable,
	logger log.Logger,
	CORSOrigin *regexp.Regexp,
) *API {
	a := &API{
		QueryEngine: qe,
		Queryable:   q,
		now:         time.Now,
		logger:      logger,
		CORSOrigin:  CORSOrigin,
	}

	return a
}

func setUnavailStatusOnTSDBNotReady(r apiFuncResult) apiFuncResult {
	if r.err != nil && errors.Cause(r.err.err) == tsdb.ErrNotReady {
		r.err.typ = errorUnavailable
	}
	return r
}

// Register the API's endpoints in the given router.
func (api *API) Register(r *route.Router) {
	wrap := func(f apiFunc) http.HandlerFunc {
		hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httputil.SetCORS(w, api.CORSOrigin, r)
			result := setUnavailStatusOnTSDBNotReady(f(r))
			if result.finalizer != nil {
				defer result.finalizer()
			}
			if result.err != nil {
				api.respondError(w, result.err, result.data)
				return
			}

			if result.data != nil {
				api.respond(w, result.data, result.warnings)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})
		return httputil.CompressionHandler{
			Handler: hf,
		}.ServeHTTP
	}

	r.Options("/*path", wrap(api.options))

	r.Get("/query", wrap(api.query))
	r.Post("/query", wrap(api.query))
	r.Get("/query_range", wrap(api.queryRange))
	r.Post("/query_range", wrap(api.queryRange))

	r.Get("/labels", wrap(api.labelNames))
	r.Post("/labels", wrap(api.labelNames))
	r.Get("/label/:name/values", wrap(api.labelValues))

	r.Get("/series", wrap(api.series))
	r.Post("/series", wrap(api.series))
}

type queryData struct {
	ResultType parser.ValueType  `json:"resultType"`
	Result     parser.Value      `json:"result"`
	Stats      *stats.QueryStats `json:"stats,omitempty"`
}

func invalidParamError(err error, parameter string) apiFuncResult {
	return apiFuncResult{nil, &apiError{
		errorBadData, errors.Wrapf(err, "invalid parameter %q", parameter),
	}, nil, nil}
}

func (api *API) options(r *http.Request) apiFuncResult {
	return apiFuncResult{nil, nil, nil, nil}
}

func (api *API) query(r *http.Request) (result apiFuncResult) {
	ts, err := parseTimeParam(r, "time", api.now())
	if err != nil {
		return invalidParamError(err, "time")
	}
	ctx := r.Context()
	if to := r.FormValue("timeout"); to != "" {
		var cancel context.CancelFunc
		timeout, err := parseDuration(to)
		if err != nil {
			return invalidParamError(err, "timeout")
		}

		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	qry, err := api.QueryEngine.NewInstantQuery(api.Queryable, r.FormValue("query"), ts)
	if err == promql.ErrValidationAtModifierDisabled {
		err = errors.New("@ modifier is disabled, use --enable-feature=promql-at-modifier to enable it")
	} else if err == promql.ErrValidationNegativeOffsetDisabled {
		err = errors.New("negative offset is disabled, use --enable-feature=promql-negative-offset to enable it")
	}
	if err != nil {
		return invalidParamError(err, "query")
	}

	// From now on, we must only return with a finalizer in the result (to
	// be called by the caller) or call qry.Close ourselves (which is
	// required in the case of a panic).
	defer func() {
		if result.finalizer == nil {
			qry.Close()
		}
	}()

	ctx = httputil.ContextFromRequest(ctx, r)

	res := qry.Exec(ctx)
	if res.Err != nil {
		return apiFuncResult{nil, returnAPIError(res.Err), res.Warnings, qry.Close}
	}

	// Optional stats field in response if parameter "stats" is not empty.
	var qs *stats.QueryStats
	if r.FormValue("stats") != "" {
		qs = stats.NewQueryStats(qry.Stats())
	}

	return apiFuncResult{&queryData{
		ResultType: res.Value.Type(),
		Result:     res.Value,
		Stats:      qs,
	}, nil, res.Warnings, qry.Close}
}

func (api *API) queryRange(r *http.Request) (result apiFuncResult) {
	start, err := parseTime(r.FormValue("start"))
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTime(r.FormValue("end"))
	if err != nil {
		return invalidParamError(err, "end")
	}
	if end.Before(start) {
		return invalidParamError(errors.New("end timestamp must not be before start time"), "end")
	}

	step, err := parseDuration(r.FormValue("step"))
	if err != nil {
		return invalidParamError(err, "step")
	}

	if step <= 0 {
		return invalidParamError(errors.New("zero or negative query resolution step widths are not accepted. Try a positive integer"), "step")
	}

	// For safety, limit the number of returned points per timeseries.
	// This is sufficient for 60s resolution for a week or 1h resolution for a year.
	if end.Sub(start)/step > 11000 {
		err := errors.New("exceeded maximum resolution of 11,000 points per timeseries. Try decreasing the query resolution (?step=XX)")
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	ctx := r.Context()
	if to := r.FormValue("timeout"); to != "" {
		var cancel context.CancelFunc
		timeout, err := parseDuration(to)
		if err != nil {
			return invalidParamError(err, "timeout")
		}

		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	qry, err := api.QueryEngine.NewRangeQuery(api.Queryable, r.FormValue("query"), start, end, step)
	if err == promql.ErrValidationAtModifierDisabled {
		err = errors.New("@ modifier is disabled, use --enable-feature=promql-at-modifier to enable it")
	} else if err == promql.ErrValidationNegativeOffsetDisabled {
		err = errors.New("negative offset is disabled, use --enable-feature=promql-negative-offset to enable it")
	}
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}
	// From now on, we must only return with a finalizer in the result (to
	// be called by the caller) or call qry.Close ourselves (which is
	// required in the case of a panic).
	defer func() {
		if result.finalizer == nil {
			qry.Close()
		}
	}()

	ctx = httputil.ContextFromRequest(ctx, r)

	res := qry.Exec(ctx)
	if res.Err != nil {
		return apiFuncResult{nil, returnAPIError(res.Err), res.Warnings, qry.Close}
	}

	// Optional stats field in response if parameter "stats" is not empty.
	var qs *stats.QueryStats
	if r.FormValue("stats") != "" {
		qs = stats.NewQueryStats(qry.Stats())
	}

	return apiFuncResult{&queryData{
		ResultType: res.Value.Type(),
		Result:     res.Value,
		Stats:      qs,
	}, nil, res.Warnings, qry.Close}
}

func returnAPIError(err error) *apiError {
	if err == nil {
		return nil
	}

	switch errors.Cause(err).(type) {
	case promql.ErrQueryCanceled:
		return &apiError{errorCanceled, err}
	case promql.ErrQueryTimeout:
		return &apiError{errorTimeout, err}
	case promql.ErrStorage:
		return &apiError{errorInternal, err}
	}

	return &apiError{errorExec, err}
}

func (api *API) labelNames(r *http.Request) apiFuncResult {
	start, err := parseTimeParam(r, "start", minTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", maxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}

	matcherSets, err := parseMatchersParam(r.Form["match[]"])
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	q, err := api.Queryable.Querier(r.Context(), timestamp.FromTime(start), timestamp.FromTime(end))
	if err != nil {
		return apiFuncResult{nil, &apiError{errorExec, err}, nil, nil}
	}
	defer q.Close()

	var (
		names    []string
		warnings storage.Warnings
	)
	if len(matcherSets) > 0 {
		hints := &storage.SelectHints{
			Start: timestamp.FromTime(start),
			End:   timestamp.FromTime(end),
			Func:  "series", // There is no series function, this token is used for lookups that don't need samples.
		}

		labelNamesSet := make(map[string]struct{})
		// Get all series which match matchers.
		for _, mset := range matcherSets {
			s := q.Select(false, hints, mset...)
			for s.Next() {
				series := s.At()
				for _, lb := range series.Labels() {
					labelNamesSet[lb.Name] = struct{}{}
				}
			}
			warnings = append(warnings, s.Warnings()...)
			if err := s.Err(); err != nil {
				return apiFuncResult{nil, &apiError{errorExec, err}, warnings, nil}
			}
		}

		// Convert the map to an array.
		names = make([]string, 0, len(labelNamesSet))
		for key := range labelNamesSet {
			names = append(names, key)
		}
		sort.Strings(names)
	} else {
		names, warnings, err = q.LabelNames()
		if err != nil {
			return apiFuncResult{nil, &apiError{errorExec, err}, warnings, nil}
		}
	}

	if names == nil {
		names = []string{}
	}
	return apiFuncResult{names, nil, warnings, nil}
}

func (api *API) labelValues(r *http.Request) (result apiFuncResult) {
	ctx := r.Context()
	name := route.Param(ctx, "name")

	if !model.LabelNameRE.MatchString(name) {
		return apiFuncResult{nil, &apiError{errorBadData, errors.Errorf("invalid label name: %q", name)}, nil, nil}
	}

	start, err := parseTimeParam(r, "start", minTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", maxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}

	matcherSets, err := parseMatchersParam(r.Form["match[]"])
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	q, err := api.Queryable.Querier(r.Context(), timestamp.FromTime(start), timestamp.FromTime(end))
	if err != nil {
		return apiFuncResult{nil, &apiError{errorExec, err}, nil, nil}
	}
	// From now on, we must only return with a finalizer in the result (to
	// be called by the caller) or call q.Close ourselves (which is required
	// in the case of a panic).
	defer func() {
		if result.finalizer == nil {
			q.Close()
		}
	}()
	closer := func() {
		q.Close()
	}

	var (
		vals     []string
		warnings storage.Warnings
	)
	if len(matcherSets) > 0 {
		var callWarnings storage.Warnings
		labelValuesSet := make(map[string]struct{})
		for _, matchers := range matcherSets {
			vals, callWarnings, err = q.LabelValues(name, matchers...)
			if err != nil {
				return apiFuncResult{nil, &apiError{errorExec, err}, warnings, closer}
			}
			warnings = append(warnings, callWarnings...)
			for _, val := range vals {
				labelValuesSet[val] = struct{}{}
			}
		}

		vals = make([]string, 0, len(labelValuesSet))
		for val := range labelValuesSet {
			vals = append(vals, val)
		}
	} else {
		vals, warnings, err = q.LabelValues(name)
		if err != nil {
			return apiFuncResult{nil, &apiError{errorExec, err}, warnings, closer}
		}

		if vals == nil {
			vals = []string{}
		}
	}

	sort.Strings(vals)

	return apiFuncResult{vals, nil, warnings, closer}
}

var (
	minTime = time.Unix(math.MinInt64/1000+62135596801, 0).UTC()
	maxTime = time.Unix(math.MaxInt64/1000-62135596801, 999999999).UTC()

	minTimeFormatted = minTime.Format(time.RFC3339Nano)
	maxTimeFormatted = maxTime.Format(time.RFC3339Nano)
)

func (api *API) series(r *http.Request) (result apiFuncResult) {
	if err := r.ParseForm(); err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, errors.Wrapf(err, "error parsing form values")}, nil, nil}
	}
	if len(r.Form["match[]"]) == 0 {
		return apiFuncResult{nil, &apiError{errorBadData, errors.New("no match[] parameter provided")}, nil, nil}
	}

	start, err := parseTimeParam(r, "start", minTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", maxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}

	matcherSets, err := parseMatchersParam(r.Form["match[]"])
	if err != nil {
		return invalidParamError(err, "match[]")
	}

	q, err := api.Queryable.Querier(r.Context(), timestamp.FromTime(start), timestamp.FromTime(end))
	if err != nil {
		return apiFuncResult{nil, &apiError{errorExec, err}, nil, nil}
	}
	// From now on, we must only return with a finalizer in the result (to
	// be called by the caller) or call q.Close ourselves (which is required
	// in the case of a panic).
	defer func() {
		if result.finalizer == nil {
			q.Close()
		}
	}()
	closer := func() {
		q.Close()
	}

	hints := &storage.SelectHints{
		Start: timestamp.FromTime(start),
		End:   timestamp.FromTime(end),
		Func:  "series", // There is no series function, this token is used for lookups that don't need samples.
	}

	var sets []storage.SeriesSet
	for _, mset := range matcherSets {
		// We need to sort this select results to merge (deduplicate) the series sets later.
		s := q.Select(true, hints, mset...)
		sets = append(sets, s)
	}

	set := storage.NewMergeSeriesSet(sets, storage.ChainedSeriesMerge)
	metrics := []labels.Labels{}
	for set.Next() {
		metrics = append(metrics, set.At().Labels())
	}

	warnings := set.Warnings()
	if set.Err() != nil {
		return apiFuncResult{nil, &apiError{errorExec, set.Err()}, warnings, closer}
	}

	return apiFuncResult{metrics, nil, warnings, closer}
}

// GlobalURLOptions contains fields used for deriving the global URL for local targets.
type GlobalURLOptions struct {
	ListenAddress string
	Host          string
	Scheme        string
}

// sanitizeSplitHostPort acts like net.SplitHostPort.
// Additionally, if there is no port in the host passed as input, we return the
// original host, making sure that IPv6 addresses are not surrounded by square
// brackets.
func sanitizeSplitHostPort(input string) (string, string, error) {
	host, port, err := net.SplitHostPort(input)
	if err != nil && strings.HasSuffix(err.Error(), "missing port in address") {
		var errWithPort error
		host, _, errWithPort = net.SplitHostPort(input + ":80")
		if errWithPort == nil {
			err = nil
		}
	}
	return host, port, err
}

func (api *API) respond(w http.ResponseWriter, data interface{}, warnings storage.Warnings) {
	statusMessage := statusSuccess
	var warningStrings []string
	for _, warning := range warnings {
		warningStrings = append(warningStrings, warning.Error())
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&response{
		Status:   statusMessage,
		Data:     data,
		Warnings: warningStrings,
	})
	if err != nil {
		level.Error(api.logger).Log("msg", "error marshaling json response", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if n, err := w.Write(b); err != nil {
		level.Error(api.logger).Log("msg", "error writing response", "bytesWritten", n, "err", err)
	}
}

func (api *API) respondError(w http.ResponseWriter, apiErr *apiError, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&response{
		Status:    statusError,
		ErrorType: apiErr.typ,
		Error:     apiErr.err.Error(),
		Data:      data,
	})

	if err != nil {
		level.Error(api.logger).Log("msg", "error marshaling json response", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var code int
	switch apiErr.typ {
	case errorBadData:
		code = http.StatusBadRequest
	case errorExec:
		code = http.StatusUnprocessableEntity
	case errorCanceled, errorTimeout:
		code = http.StatusServiceUnavailable
	case errorInternal:
		code = http.StatusInternalServerError
	case errorNotFound:
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if n, err := w.Write(b); err != nil {
		level.Error(api.logger).Log("msg", "error writing response", "bytesWritten", n, "err", err)
	}
}

func parseTimeParam(r *http.Request, paramName string, defaultValue time.Time) (time.Time, error) {
	val := r.FormValue(paramName)
	if val == "" {
		return defaultValue, nil
	}
	result, err := parseTime(val)
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "Invalid time value for '%s'", paramName)
	}
	return result, nil
}

func parseTime(s string) (time.Time, error) {
	if t, err := strconv.ParseFloat(s, 64); err == nil {
		s, ns := math.Modf(t)
		ns = math.Round(ns*1000) / 1000
		return time.Unix(int64(s), int64(ns*float64(time.Second))).UTC(), nil
	}
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, nil
	}

	// Stdlib's time parser can only handle 4 digit years. As a workaround until
	// that is fixed we want to at least support our own boundary times.
	// Context: https://github.com/prometheus/client_golang/issues/614
	// Upstream issue: https://github.com/golang/go/issues/20555
	switch s {
	case minTimeFormatted:
		return minTime, nil
	case maxTimeFormatted:
		return maxTime, nil
	}
	return time.Time{}, errors.Errorf("cannot parse %q to a valid timestamp", s)
}

func parseDuration(s string) (time.Duration, error) {
	if d, err := strconv.ParseFloat(s, 64); err == nil {
		ts := d * float64(time.Second)
		if ts > float64(math.MaxInt64) || ts < float64(math.MinInt64) {
			return 0, errors.Errorf("cannot parse %q to a valid duration. It overflows int64", s)
		}
		return time.Duration(ts), nil
	}
	if d, err := model.ParseDuration(s); err == nil {
		return time.Duration(d), nil
	}
	return 0, errors.Errorf("cannot parse %q to a valid duration", s)
}

func parseMatchersParam(matchers []string) ([][]*labels.Matcher, error) {
	var matcherSets [][]*labels.Matcher
	for _, s := range matchers {
		matchers, err := parser.ParseMetricSelector(s)
		if err != nil {
			return nil, err
		}
		matcherSets = append(matcherSets, matchers)
	}

OUTER:
	for _, ms := range matcherSets {
		for _, lm := range ms {
			if lm != nil && !lm.Matches("") {
				continue OUTER
			}
		}
		return nil, errors.New("match[] must contain at least one non-empty matcher")
	}
	return matcherSets, nil
}

// marshalPointJSON writes `[ts, "val"]`.
func marshalPointJSON(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	p := *((*promql.Point)(ptr))
	stream.WriteArrayStart()
	marshalTimestamp(p.T, stream)
	stream.WriteMore()
	marshalValue(p.V, stream)
	stream.WriteArrayEnd()
}

func marshalPointJSONIsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func marshalTimestamp(t int64, stream *jsoniter.Stream) {
	// Write out the timestamp as a float divided by 1000.
	// This is ~3x faster than converting to a float.
	if t < 0 {
		stream.WriteRaw(`-`)
		t = -t
	}
	stream.WriteInt64(t / 1000)
	fraction := t % 1000
	if fraction != 0 {
		stream.WriteRaw(`.`)
		if fraction < 100 {
			stream.WriteRaw(`0`)
		}
		if fraction < 10 {
			stream.WriteRaw(`0`)
		}
		stream.WriteInt64(fraction)
	}
}

func marshalValue(v float64, stream *jsoniter.Stream) {
	stream.WriteRaw(`"`)
	// Taken from https://github.com/json-iterator/go/blob/master/stream_float.go#L71 as a workaround
	// to https://github.com/json-iterator/go/issues/365 (jsoniter, to follow json standard, doesn't allow inf/nan).
	buf := stream.Buffer()
	abs := math.Abs(v)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if abs != 0 {
		if abs < 1e-6 || abs >= 1e21 {
			fmt = 'e'
		}
	}
	buf = strconv.AppendFloat(buf, v, fmt, -1, 64)
	stream.SetBuffer(buf)
	stream.WriteRaw(`"`)
}

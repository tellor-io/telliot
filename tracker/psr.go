package tracker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"sync"
	"time"
)

//PSRTracker keeps track of pre-specified requests
type PSRTracker struct {
	Requests    []PrespecifiedRequest `json:"prespecifiedRequests"`
	requestByID map[uint]*PrespecifiedRequest
}

//PrespecifiedRequest holds fields for pre-specific requests
type PrespecifiedRequest struct {
	RequestID      uint     `json:"requestID"`
	APIs           []string `json:"apis"`
	Transformation string   `json:"transformation"`
	Granularity    uint     `json:"granularity"`
}



var (
	psrLog          = util.NewLogger("tracker", "PSRTracker")
	funcs           map[string]interface{}
)

//maps requestID to a time window of values
var PSRWindows map[uint]*Window
var lastWrotePSR time.Time
var PSRWindowMutex sync.RWMutex

func GetLatestRequestValue(id uint) *TimedInt {
	PSRWindowMutex.RLock()
	defer PSRWindowMutex.RUnlock()
	w, ok := PSRWindows[id]
	if !ok {
		return nil
	}
	return w.Latest()
}

func setRequestValue(id uint, val *TimedInt) {
	PSRWindowMutex.Lock()
	_, ok := PSRWindows[id]
	if !ok {
		PSRWindows[id] = NewWindow(7 * 24 * time.Hour)
	}
	PSRWindows[id].Insert(val)
	PSRWindowMutex.Unlock()
}

func writeOutPSR() {
	PSRWindowMutex.Lock()
	for _,v := range PSRWindows {
		v.Trim()
	}
	data, err := json.MarshalIndent(PSRWindows, "", "\t")
	PSRWindowMutex.Unlock()
	if err != nil {
		psrLog.Error("failed to marshal PSR values: %s", err.Error())
		return
	}

	cfg := config.GetConfig()
	psrSavedData := filepath.Join(cfg.PSRFolder, "saved.json")
	psrSavedDataTmp := psrSavedData + ".tmp"
	err = ioutil.WriteFile(psrSavedDataTmp, data, 0644)
	if err != nil {
		psrLog.Error("failed to write out PSR values to %s: %s", psrSavedDataTmp, err.Error())
		return
	}
	//rename tmp file to old file (should be atomic on most modern OS)
	err = os.Rename(psrSavedDataTmp, psrSavedData)
	if err != nil {
		psrLog.Error("failed move new PSR save onto old: %s", err.Error())
		return
	}
	lastWrotePSR = time.Now()
}

//BuildPSRTracker creates and initializes a new tracker instance
func BuildPSRTracker() (*PSRTracker, error) {

	psr := &PSRTracker{Requests: nil, requestByID: make(map[uint]*PrespecifiedRequest)}
	if err := psr.init(); err != nil {
		return nil, err
	}
	funcs = map[string]interface{}{
		"value":   value,
		"average": average,
		"median":  median,
		"square":  square,
	}
	return psr, nil
}

//String name of this tracker
func (psr *PSRTracker) String() string {
	return "PSRTracker"
}

func (psr *PSRTracker) init() error {
	//Loop through all PSRs
	cfg := config.GetConfig()

	psrPath := filepath.Join(cfg.PSRFolder, "psr.json")
	byteValue, err := ioutil.ReadFile(psrPath)
	if err != nil {
		return fmt.Errorf("failed to read psr file @ %s: %v", psrPath, err)
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'Requests' which we defined above
	err = json.Unmarshal(byteValue, &psr)
	if err != nil {
		return err
	}
	for i := 0; i < len(psr.Requests); i++ {
		r := psr.Requests[i]
		psr.requestByID[r.RequestID] = &r
	}

	psrSavedData := filepath.Join(cfg.PSRFolder, "saved.json")

	_, err = os.Stat(psrSavedData)
	exists := true
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			return fmt.Errorf("file %s stat error: %v", psrSavedData, err)
		}
	}

	if exists {
		byteValue, err = ioutil.ReadFile(psrSavedData)
		if err != nil {
			return fmt.Errorf("failed to read psr file @ %s: %v", psrPath, err)
		}
		err = json.Unmarshal(byteValue, &PSRWindows)
		if err != nil {
			return fmt.Errorf("failed to unmarshal saved")
		}
	} else {
		PSRWindows = make(map[uint]*Window)
	}

	psrLog.Info("Initialized PSR with %d requests\n", len(psr.Requests))
	return nil
}

//Exec implements tracker API
func (psr *PSRTracker) Exec(ctx context.Context) error {
	//TODO: retrieve github updates of psr config file. For now, we'll just pull
	//PSR's as defined by psr.json file
	resultCh := make(chan *fetchResult, len(psr.Requests))
	for i := 0; i < len(psr.Requests); i++ {
		p := psr.Requests[i]
		go p.fetch(resultCh)
	}
	nerr := 0
	for i := 0; i < len(psr.Requests); i++ {
		result := <- resultCh
		if result.err != nil {
			psrLog.Error("Problem fetching PSR for id %d: %v]\n", result.r.RequestID, result.err)
			nerr++
		} else {
			setRequestValue(result.r.RequestID, result.val)
		}
	}
	if nerr > 0 {
		psrLog.Info("PSR Tracker cycle completed with %d errors", nerr)
	} else {
		psrLog.Info("PSR Tracker cycle completed succesfully")
	}
	now := time.Now()
	if now.Sub(lastWrotePSR) > 2 * time.Minute {
		writeOutPSR()
	}
	return nil
}

type fetchResult struct {
	r *PrespecifiedRequest
	val *TimedInt
	err error
}

func (r *PrespecifiedRequest) fetch(resultCh chan *fetchResult) {
	cfg := config.GetConfig()
	reqs := make([]*FetchRequest, len(r.APIs))
	argGroups := make([][]string, len(r.APIs))
	for i := 0; i < len(r.APIs); i++ {
		api := r.APIs[i]
		url, args := util.ParseQueryString(api)
		reqs[i] = &FetchRequest{queryURL: url, timeout: cfg.FetchTimeout.Duration}
		argGroups[i] = args
	}
	payloads, _ := batchFetchWithRetries(reqs)
	vals := make([]int, 0, len(payloads))
	errs := 0
	for i, pl := range payloads {
		if pl == nil {
			errs += 1
			continue
		}
		v, err := util.ParsePayload(pl, r.Granularity, argGroups[i])
		if err != nil {
			errs += 1
			continue
		}
		vals = append(vals, v)
	}
	result := &fetchResult{r: r}
	if len(vals) > 0 {
		res, err := computeTransformation(r.Transformation, vals)
		if err != nil {
			 result.err = err
		} else {
			result.val = &TimedInt{
				Created: time.Now(),
				Val:     res.Interface().(uint),
			}
		}
	} else {
		result.err = fmt.Errorf("no sucessful api hits, no value stored for id %d", r.RequestID)
	}
	resultCh <- result
}

func computeTransformation(name string, params ...interface{}) (result reflect.Value, err error) {
	f := reflect.ValueOf(funcs[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)[0]
	return result, nil
}

func value(num []int) uint {
	//fmt.Println("Calling Value", num)
	return uint(num[0])
}

func average(nums []int) uint {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	//fmt.Println("Average", sum/len(nums))
	return uint(sum / len(nums))
}

func median(num []int) uint {
	sort.Ints(num)
	//fmt.Println("Median", num[len(num)/2])
	return uint(num[len(num)/2])
}

func square(num int) int {
	//fmt.Println("Square", num*num)
	return num * num
}

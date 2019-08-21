package tracker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/cli"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
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
	sharedInstance  *PSRTracker
	funcs           map[string]interface{}
	psrWaitGroupKey = util.NewKey("tracker", "PSRFetchWaitGroup")
)

//package function to get a shared PSRInstance
func psrInstance() (*PSRTracker, error) {
	if sharedInstance == nil {
		return nil, fmt.Errorf("Missing psrInstance singleton")
	}
	return sharedInstance, nil
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
	sharedInstance = psr
	return psr, nil
}

//String name of this tracker
func (psr *PSRTracker) String() string {
	return "PSRTracker"
}

func (psr *PSRTracker) init() error {
	//Loop through all PSRs
	psrPath := cli.GetFlags().PSRPath
	psrLog.Info("Opening PSR config file at: %s\n", psrPath)

	configFile, err := os.Open(psrPath)

	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	defer configFile.Close()
	byteValue, _ := ioutil.ReadAll(configFile)
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
	psrLog.Info("Initialized PSR with %d requests\n", len(psr.Requests))
	return nil
}

//Exec implements tracker API
func (psr *PSRTracker) Exec(ctx context.Context) error {
	//TODO: retrieve github updates of psr config file. For now, we'll just pull
	//PSR's as defined by psr.json file
	var syncGroup sync.WaitGroup
	var doneGroup sync.WaitGroup
	ctx = context.WithValue(ctx, psrWaitGroupKey, &syncGroup)
	errorCh := make(chan error)
	fmt.Println("Starting")
	doneGroup.Add(1)
	go func() {
		defer doneGroup.Done()
		for {
			e := <-errorCh
			if e != nil {
				psrLog.Error("Problem in PSR fetch: %v]\n", e)
			} else {
				psrLog.Info("Finishing PSR tracker run...")
				return
			}
		}
	}()

	for i := 0; i < len(psr.Requests); i++ {
		p := psr.Requests[i]
		syncGroup.Add(1)
		psrLog.Info("Fetching PSR with id: %v\n", p.RequestID)
		go p.fetch(ctx, errorCh)
	}
	psrLog.Info("Waiting for PSR's to complete...")
	syncGroup.Wait()
	errorCh <- nil
	psrLog.Info("Waiting for exit on error reader...")
	doneGroup.Wait()
	psrLog.Info("PSR Tracker cycle complete")
	return nil
}

func (r *PrespecifiedRequest) fetch(ctx context.Context, errorCh chan error) {
	syncGroup := ctx.Value(psrWaitGroupKey).(*sync.WaitGroup)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	defer syncGroup.Done()
	cfg, err := config.GetConfig()
	if err != nil {
		errorCh <- err
		return
	}
	timeout := time.Duration(time.Duration(cfg.FetchTimeout) * time.Second)
	reqs := make([]*FetchRequest, len(r.APIs))
	argGroups := make([][]string, len(r.APIs))
	for i := 0; i < len(r.APIs); i++ {
		api := r.APIs[i]
		url, args := util.ParseQueryString(api)
		reqs[i] = &FetchRequest{queryURL: url, timeout: timeout}
		argGroups[i] = args
	}
	payloads, err := batchFetchWithRetries(reqs)
	vals := make([]int, len(payloads))
	for i := 0; i < len(vals); i++ {
		pl := payloads[i]
		if pl == nil {
			vals[i] = -1
			continue
		}
		v, err := util.ParsePayload(payloads[i], r.Granularity, argGroups[i])
		if err != nil {
			errorCh <- err
			vals[i] = -1
			continue
		}
		vals[i] = v
	}
	res, err := computeTransformation(r.Transformation, vals)
	if err != nil {
		errorCh <- err
	} else {
		y := res.Interface().(uint)
		enc := hexutil.EncodeBig(big.NewInt(int64(y)))
		DB.Put(fmt.Sprintf("%s%d", db.QueriedValuePrefix, r.RequestID), []byte(enc))
	}
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

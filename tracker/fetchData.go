package tracker

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"sort"
	"sync"
	"time"

	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
)

const API = "json(https://api.gdax.com/products/ETH-USD/ticker).price"

var waitGroupKey = util.NewKey("tracker", "fetchWaitGroup")

type RequestDataTracker struct {
}

type FetchResult struct {
	value *big.Int
	reqId uint
}

/*
var thisPSR PrespecifiedRequest
var psr PrespecifiedRequests
*/

func (b *RequestDataTracker) String() string {
	return "RequestDataTracker"
}

func (b *RequestDataTracker) Exec(ctx context.Context) error {
	//DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	var syncGroup sync.WaitGroup
	ctx = context.WithValue(ctx, waitGroupKey, syncGroup)

	/*
		funcs := map[string]interface{}{
			"value":   value,
			"average": average,
			"median":  median,
			"square":  square,
		}
	*/

	//pre-initialized PSR's
	psr, err := psrInstance()
	if err != nil {
		return err
	}

	//we will wait for all requests to complete before this tracker is done
	syncGroup.Add(len(psr.Requests))

	//we spread fetch calls across go routines and they write their values back
	//dbChan := make(chan *FetchResult)

	//handle errors in some way. For now, we just log them. But better model
	//would be to push them to a more abstract error logging/alerting mechanism
	//for monitoring activity
	//errChan := make(chan error)

	return nil

	/***
	for i := 0; i < len(psr.Requests); i++ {
		thisPSR := psr.Requests[i]
		var myFetches []int
		for i := 0; i < len(thisPSR.APIs); i++ {
			myFetches = append(myFetches, fetchAPI(thisPSR.Granularity, thisPSR.APIs[i]))
		}
		res, _ := CallPrespecifiedRequest(funcs, thisPSR.Transformation, myFetches)
		y := res[0].Interface().(uint)
		//fmt.Println(big.NewInt(int64(y)))
		enc = hexutil.EncodeBig(big.NewInt(int64(y)))
		fmt.Println("Storing Fetch Data", fmt.Sprint(thisPSR.RequestID))
		DB.Put(fmt.Sprintf("%s%d", db.QueriedValuePrefix, thisPSR.RequestID), []byte(enc))
	}
	//Loop through all those in Top50
	v, err := DB.Get(db.Top50Key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 1; i < len(v); i++ {

		id := int(v[i])
		//pull request metadata from DB
		queryMeta, err := DB.Get(fmt.Sprintf("%s%d", db.QueryMetadataPrefix, id))
		if err != nil {

		}

		if i1 > 0 {
			isPre, _, _ := checkPrespecifiedRequest(uint(i1))
			if isPre {
				fmt.Println("Prespec")
			} else {
				fmt.Println("Normal Fetch")
				//We need to go get the queryString (we should store it somewhere)
				//also we need the granularity
				fetchres := int64(fetchAPI(1000, API))
				enc = hexutil.EncodeBig(big.NewInt(fetchres))
				DB.Put(fmt.Sprint(i1), []byte(enc))
			}
		}
	}
	return nil
	**/
}

func fetchAPI(ctx context.Context, _granularity uint, queryString string, resultChan chan *FetchResult, errorChan chan error) {
	//DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	syncGroup := ctx.Value(waitGroupKey).(sync.WaitGroup)
	defer syncGroup.Done()

	cfg, err := config.GetConfig()
	if err != nil {
		errorChan <- err
		return
	}

	timeout := time.Duration(time.Duration(cfg.FetchTimeout) * time.Second)

	//TODO: need to abstract how to get URL string and JSON parsinglogic
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	url := rgx.FindStringSubmatch(queryString)[1]
	req := &FetchRequest{queryURL: url, timeout: timeout}

	valueStr, err := fetchWithRetries(req)
	if err != nil {
		errorChan <- err
		return
	}
}

/*
func oldfetchAPI(_granularity uint, queryString string) int {
	var r QueryStruct
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	url := rgx.FindStringSubmatch(queryString)[1]
	resp, _ := http.Get(url)
	input, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(input, &r)
	if err != nil {
		panic(err)
	}
	s, err := strconv.ParseFloat(r.Price, 64)
	fmt.Println(s * float64(_granularity)) //need to get granularity
	return int(s * float64(_granularity))
}

func checkPrespecifiedRequest(requestID uint) (bool, PrespecifiedRequest, error) {
	configFile, err := os.Open("../psr.json")

	if err != nil {
		fmt.Println("Error", err)
		return false, thisPSR, err
	}
	defer configFile.Close()
	byteValue, _ := ioutil.ReadAll(configFile)
	var psr PrespecifiedRequests
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &psr)
	fmt.Println(psr)
	for i := 0; i < len(psr.PrespecifiedRequests); i++ {
		if psr.PrespecifiedRequests[i].RequestID == requestID {
			thisPSR = psr.PrespecifiedRequests[i]
			fmt.Println("Id: ", psr.PrespecifiedRequests[i].RequestID)
			fmt.Println("APIs: ", psr.PrespecifiedRequests[i].APIs)
			fmt.Println("Transformation: ", psr.PrespecifiedRequests[i].Transformation)
			return true, thisPSR, nil
		}
	}
	return false, thisPSR, nil
}
*/

func value(num []int) uint {
	fmt.Println("Calling Value", num)
	return uint(num[0])
}

func average(nums []int) uint {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return uint(sum / len(nums))
}

func median(num []int) uint {
	sort.Ints(num)
	return uint(num[len(num)/2])
}

func square(num int) int {
	return num * num
}

func CallPrespecifiedRequest(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	fmt.Println("Result", result)
	return
}

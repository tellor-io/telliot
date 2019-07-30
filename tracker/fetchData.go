package tracker

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
)

//const API = "json(https://api.gdax.com/products/ETH-USD/ticker).price"

var fetchLog = util.NewLogger("tracker", "FetchDataTracker")
var waitGroupKey = util.NewKey("tracker", "fetchWaitGroup")

//RequestDataTracker fetches data from remote APIs
type RequestDataTracker struct {
}

//FetchResult contains fields after a fetch completes
type FetchResult struct {
	value *big.Int
	reqID uint
}

/*
var thisPSR PrespecifiedRequest
var psr PrespecifiedRequests
*/

func (b *RequestDataTracker) String() string {
	return "RequestDataTracker"
}

//Exec is implementation of tracker interface
func (b *RequestDataTracker) Exec(ctx context.Context) error {
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	var syncGroup sync.WaitGroup
	var doneGroup sync.WaitGroup
	doneGroup.Add(2)
	ctx = context.WithValue(ctx, waitGroupKey, &syncGroup)

	//we spread fetch calls across go routines and they write their values back
	dbChan := make(chan *FetchResult)

	go func() {
		defer doneGroup.Done()
		for {
			res := <-dbChan
			if res == nil {
				return
			}
			enc := hexutil.EncodeBig(res.value)
			DB.Put(fmt.Sprintf("%s%d", db.QueriedValuePrefix, res.reqID), []byte(enc))
		}
	}()

	//handle errors in some way. For now, we just log them. But better model
	//would be to push them to a more abstract error logging/alerting mechanism
	//for monitoring activity
	errChan := make(chan error)
	go func() {
		defer doneGroup.Done()
		for {
			e := <-errChan
			if e == nil {
				return
			}
			fetchLog.Error("Problem during fetch data: %v\n", e)
		}
	}()

	v, err := DB.Get(db.Top50Key)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i := 1; i < len(v); i++ {

		id := int(v[i])
		//pull request metadata from DB
		queryMetaBytes, err := DB.Get(fmt.Sprintf("%s%d", db.QueryMetadataPrefix, id))
		if err != nil {
			errChan <- err
			continue
		}
		fmt.Println("Value", string(queryMetaBytes))
	}

	syncGroup.Wait()
	errChan <- nil
	dbChan <- nil

	doneGroup.Wait()

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
	fmt.Println("Why doesn't this work?", v)
	for i := 0; i < len(v); i++ {
		fmt.Println(i)
		fmt.Println("Line 88", v[i])
		i1 := int(v[i])
		if i1 > 0 {
			isPre, _, _ := checkPrespecifiedRequest(uint(i1))
			if isPre {
				fmt.Println("Prespec")
			} else {
				fmt.Println("Normal Fetch")
				//We need to go get the queryString (we should store it somewhere)
				//also we need the granularity
				fetchres := int64(fetchAPI(1000, API))
				fmt.Println(big.NewInt(fetchres))
				enc = hexutil.EncodeBig(big.NewInt(fetchres))
				DB.Put(fmt.Sprint(i1), []byte(enc))
			}
		}
	}
	return nil
	**/
}

func fetchAPI(ctx context.Context, reqID uint, _granularity uint, queryString string, resultChan chan *FetchResult, errorChan chan error) {
	//DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	syncGroup := ctx.Value(waitGroupKey).(*sync.WaitGroup)
	defer syncGroup.Done()

	cfg, err := config.GetConfig()
	if err != nil {
		errorChan <- err
		return
	}

	timeout := time.Duration(time.Duration(cfg.FetchTimeout) * time.Second)
	url, args := util.ParseQueryString(queryString)
	req := &FetchRequest{queryURL: url, timeout: timeout}
	payload, err := fetchWithRetries(req)
	if err != nil {
		errorChan <- err
		return
	}

	val, err := util.ParsePayload(payload, _granularity, args)
	if err != nil {
		errorChan <- err
		return
	}
	asInt := big.NewInt(int64(val))
	resultChan <- &FetchResult{value: asInt, reqID: reqID}
}

/*
func oldfetchAPI(_granularity uint, queryString string) int {
	var r QueryStruct
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	url := rgx.FindStringSubmatch(queryString)[1]
	fmt.Println("url", url)
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

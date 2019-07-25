package tracker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

const API = "json(https://api.gdax.com/products/ETH-USD/ticker).price"

type RequestDataTracker struct {
}

type QueryStruct struct {
	Price string `price`
}

type PrespecifiedRequests struct {
	PrespecifiedRequests []PrespecifiedRequest `json:"prespecifiedRequests"`
}
type PrespecifiedRequest struct {
	RequestID      uint     `json:"requestID"`
	APIs           []string `json:"apis"`
	Transformation string   `json:"transformation"`
	Granularity    uint     `json:"granularity"`
}

var thisPSR PrespecifiedRequest

func (b *RequestDataTracker) Exec(ctx context.Context) error {
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	isPre, _, _ := checkPrespecifiedRequest(1)
	funcs := map[string]interface{}{
		"value":   value,
		"average": average,
		"median":  median,
		"square":  square,
	}
	enc := "0x"

	if isPre {
		fmt.Println("Prespec")
		var myFetches []int
		for i := 0; i < len(thisPSR.APIs); i++ {
			fmt.Println(i)
			myFetches = append(myFetches, fetchAPI(thisPSR.Granularity, thisPSR.APIs[i]))
		}
		fmt.Println("My Fetches", myFetches)
		res, _ := CallPrespecifiedRequest(funcs, thisPSR.Transformation, myFetches)
		y := res[0].Interface().(uint)
		fmt.Println(big.NewInt(int64(y)))
		enc = hexutil.EncodeBig(big.NewInt(int64(y)))
	} else {
		fmt.Println("Normal Fetch")
		fetchres := int64(fetchAPI(1000, API))
		fmt.Println(big.NewInt(fetchres))
		enc = hexutil.EncodeBig(big.NewInt(fetchres))
	}

	log.Printf("Staker Status: %v", enc)
	return DB.Put(fmt.Sprint(thisPSR.RequestID), []byte(enc))
}

func fetchAPI(_granularity uint, queryString string) int {
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
	configFile, err := os.Open("psr.json")

	if err != nil {
		fmt.Println("Error", err)
		return false, thisPSR, err
	}
	defer configFile.Close()
	fmt.Println(configFile)
	byteValue, _ := ioutil.ReadAll(configFile)
	fmt.Println(byteValue)
	var psr PrespecifiedRequests
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &psr)
	fmt.Println(psr)
	for i := 0; i < len(psr.PrespecifiedRequests); i++ {
		if psr.PrespecifiedRequests[i].RequestID == requestID {
			thisPSR = psr.PrespecifiedRequests[0]
			fmt.Println("Id: ", psr.PrespecifiedRequests[i].RequestID)
			fmt.Println("APIs: ", psr.PrespecifiedRequests[i].APIs)
			fmt.Println("Transformation: ", psr.PrespecifiedRequests[i].Transformation)
			return true, thisPSR, nil
		}
	}
	return false, thisPSR, nil
}

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

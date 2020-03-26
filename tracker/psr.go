package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"time"

	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
)

//PSRTracker keeps track of pre-specified requests
type PSRTracker struct {
	Requests    []PrespecifiedRequest `json:"prespecifiedRequests"`
}

//PrespecifiedRequest holds fields for pre-specific requests
type PrespecifiedRequest struct {
	RequestID      uint64     `json:"requestID"`
	APIs           []string `json:"apis"`
	Transformation ValueP   `json:"transformation"`
	Granularity    uint     `json:"granularity"`
	Symbol         string   `json:"symbol"`

	//generated from the API strings
	requests	   []*FetchRequest
	argGroups	   [][]string
}

var (
	psrLog = util.NewLogger("tracker", "PSRTracker")
	funcs  map[string]interface{}
)

//BuildPSRTrackers creates and initializes a new tracker instance
func BuildPSRTrackers() ([]Tracker, error) {

	psr := &PSRTracker{}
	if err := psr.init(); err != nil {
		return nil, err
	}
	trackers := make([]Tracker, len(psr.Requests))
	for i := range trackers {
		trackers[i] = psr.Requests[i]
	}
	return trackers, nil
}

func GetPSRByIDMap() (map[uint64]*PrespecifiedRequest, error) {
	result := make(map[uint64]*PrespecifiedRequest)
	psr := &PSRTracker{}
	if err := psr.init(); err != nil {
		return nil, err
	}
	for i := 0; i < len(psr.Requests); i++ {
		r := psr.Requests[i]
		result[r.RequestID] = &r
	}
	return result, nil
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

	err = EnsureValueOracle()
	if err != nil {
		return fmt.Errorf("failed to launch value oracle: %v", err)
	}

	for i := range psr.Requests {
		r := &psr.Requests[i]
		r.requests = make([]*FetchRequest, len(r.APIs))
		r.argGroups = make([][]string, len(r.APIs))
		for i := 0; i < len(r.APIs); i++ {
			api := r.APIs[i]
			url, args := util.ParseQueryString(api)
			r.requests[i] = &FetchRequest{queryURL: url, timeout: cfg.FetchTimeout.Duration}
			r.argGroups[i] = args
		}
	}

	psrLog.Info("Initialized PSR with %d requests\n", len(psr.Requests))
	return nil
}

//Exec implements tracker API
func (r PrespecifiedRequest) Exec(ctx context.Context) error {
	//TODO: retrieve github updates of psr config file. For now, we'll just pull
	//PSR's as defined by psr.json file

	//fetch all apis
	payloads, _ := batchFetchWithRetries(r.requests)

	//parse them according to that PSR's rules
	val, err := r.Transformation.Transform(&r, payloads)
	if err != nil {
		return err
	}
	bigVal := new(big.Int)
	bigVal.SetUint64(val)

	//save the value into our local data window no matter what
	setRequestValue(r.RequestID, time.Now(), bigVal)

	//if we have enough data saved in the window to make an accurate reading,
	//place the value into the DB so it can be submitted during mining
	//if value returns nil, we don't have enough saved yet
	value := r.Transformation.Value(&r)
	if value != nil {
		DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
		enc := hexutil.EncodeBig(bigVal)
		err := DB.Put(fmt.Sprintf("%s%d", db.QueriedValuePrefix, r.RequestID), []byte(enc))
		if err != nil {
			return err
		}
	}
	return nil
}

func (r PrespecifiedRequest) String() string {
	return fmt.Sprintf("%s PSR", r.Symbol)
}


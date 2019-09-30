package tracker

import (
	"context"
	"encoding/json"
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

//need a rate limit to throttle requests to 3rd party APIs
const fetchRate = time.Second / 4 //fetch requests per second

//log
var fetchLog = util.NewLogger("tracker", "FetchDataTracker")

//key to lookup a shared sync wait group in context
var waitGroupKey = util.NewKey("tracker", "fetchWaitGroup")

//RequestDataTracker fetches data from remote APIs
type RequestDataTracker struct {
}

//FetchResult contains fields after a fetch completes
type FetchResult struct {
	value *big.Int
	reqID uint
}

func (b *RequestDataTracker) String() string {
	return "FetchDataTracker"
}

//Exec is implementation of tracker interface
func (b *RequestDataTracker) Exec(ctx context.Context) error {

	//ticker to monitor to ensure we don't exceed max fetch request rate
	throttle := time.Tick(time.Duration(fetchRate))

	//wait group to synchronize concurrent fetch routines
	var syncGroup sync.WaitGroup

	//wait group to stop anonymous go routes below
	var doneGroup sync.WaitGroup
	//there are two of them below
	doneGroup.Add(2)

	//add the sync group to the context so go routines can grab it for use in separate thread
	ctx = context.WithValue(ctx, waitGroupKey, &syncGroup)

	//we spread fetch calls across go routines and this channel will receive their results
	dbChan := make(chan *FetchResult)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	go func() {
		//make sure we mark that we finished once all fetches are done
		defer doneGroup.Done()
		for {
			res := <-dbChan
			if res == nil {
				//below we send a nil entry through so we bail once we see that
				return
			}
			//encode the value in hex
			enc := hexutil.EncodeBig(res.value)
			fetchLog.Debug("Storing fetch result: %v for id: %d", res.value, res.reqID)
			//and write it to DB using value prefix and request Id
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

	//read the last top50 pulled from on-chain
	v, err := DB.Get(db.Top50Key)
	if err != nil {
		fetchLog.Error("Problem reading top50 from DB: %v\n", err)
		return err
	}

	//for each top50 request id
	for i := 0; i < len(v); i++ {

		id := int(v[i])

		//ignore PSR-range request ids
		if id <= 50 {
			fetchLog.Info("Ignoring PSR id range: %v\n", id)
			continue //PSR
		}

		//pull request metadata from DB
		fetchLog.Debug("Checking DB for request with id: %d\n", id)
		queryMetaBytes, err := DB.Get(fmt.Sprintf("%s%d", db.QueryMetadataPrefix, id))
		var meta *IDSpecifications
		if err != nil {
			fetchLog.Error("Problem reading request meta from DB: %v\n", err)
			errChan <- err
			continue
		}

		//all this needs to be moved to the top50Tracker so that fetching merely relies on
		//what top50 pulled in current cyle. Also means top50 must run BEFORE this tracker
		if queryMetaBytes == nil || len(queryMetaBytes) == 0 {
			fetchLog.Error("Did not find request metadata in DB, cannot fetch data for id: %v\n", id)
			continue
		}

		if meta == nil && queryMetaBytes != nil {
			meta = new(IDSpecifications)
			if err := json.Unmarshal(queryMetaBytes, meta); err != nil {
				top50Logger.Error("Problem unmarshalling query metadata from DB: %v\n", err)
			}
		}

		if meta != nil && len(meta.QueryString) > 0 {
			syncGroup.Add(1)
			<-throttle
			go fetchAPI(ctx, uint(id), uint(meta.Granularity), meta.QueryString, dbChan, errChan)
		}
	}
	syncGroup.Wait()
	errChan <- nil
	dbChan <- nil
	doneGroup.Wait()
	fetchLog.Info("All done fetching Data")
	return nil
}

func fetchAPI(ctx context.Context, reqID uint, _granularity uint, queryString string, resultChan chan *FetchResult, errorChan chan error) {
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
	fetchLog.Debug("Fetching price data from: %s\n", url)
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

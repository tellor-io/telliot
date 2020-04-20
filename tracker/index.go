package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tellor-io/TellorMiner/apiOracle"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)


var psrLog = util.NewLogger("tracker", "IndexTrackers")

var indexes map[string][]*IndexTracker

//BuildPSRTrackers creates and initializes a new tracker instance
func BuildIndexTrackers() ([]Tracker, error) {
	err := apiOracle.EnsureValueOracle()
	if err != nil {
		return nil, err
	}

	cfg := config.GetConfig()

	indexPath := filepath.Join(cfg.IndexFolder, "indexes.json")
	byteValue, err := ioutil.ReadFile(indexPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file @ %s: %v", indexPath, err)
	}
	var baseIndexes map[string][]string
	err = json.Unmarshal(byteValue, &baseIndexes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse index file: %v", err)
	}

	indexes = make(map[string][]*IndexTracker)

	//build a tracker for each unique API
	indexers := make(map[string]*IndexTracker)
	//and keep track of which APIs influence which symbols so we know what to update later
	symbolsForAPI := make(map[string][]string)

	for symbol,apis := range baseIndexes {
		for _,api := range apis {
			//did we already have a tracker for this API string?
			_,ok := indexers[api]
			if !ok {
				pathStr, args := util.ParseQueryString(api)
				var name string
				var source DataSource
				if strings.HasPrefix(pathStr, "http") {
					source = &JSONapi{&FetchRequest{queryURL: pathStr, timeout: cfg.FetchTimeout.Duration}}
					u, err := url.Parse(pathStr)
					if err != nil {
						return nil, fmt.Errorf("invalid API URL: %s: %v", pathStr, err)
					}
					name = u.Host
				} else {
					source = &JSONfile{filepath: pathStr}
					name = filepath.Base(pathStr)
				}
				indexers[api] = &IndexTracker{
					Name:       name,
					Identifier: api,
					Source:    	source,
					Args:       args,
				}
			}
			//now we definitely have one
			thisOne := indexers[api]

			//insert add it and it's more specific variant to the symbol->api map
			indexes[symbol] = append(indexes[symbol], thisOne)
			specificName := fmt.Sprintf("%s~%s", symbol, thisOne.Name)
			indexes[specificName] = append(indexes[specificName], thisOne)

			//save this for later so we can build the api->symbol map
			symbolsForAPI[api] = append(symbolsForAPI[api], symbol, specificName)
		}
	}

	//set the reverse map
	for api,symbols := range symbolsForAPI {
		indexers[api].Symbols = symbols
	}

	//make an array of trackers to be sent to Runner
	trackers := make([]Tracker, len(indexers))
	pos := 0
	for _,indexer := range indexers {
		trackers[pos] = indexer
		pos++
	}

	//start the PSR system that will feed from these indexes
	err = InitPSRs()
	if err != nil {
		return nil, fmt.Errorf("failed to initalize PSRs: %v", err)
	}

	return trackers, nil
}

type IndexTracker struct {
	Name       string
	Identifier string
	Symbols    []string
	Source	   DataSource
	Args    [][]string
}

type DataSource interface {
	Get() ([]byte, error)
}

type JSONapi struct {
	Request *FetchRequest
}

func (j *JSONapi)Get() ([]byte, error) {
	return fetchWithRetries(j.Request)
}

type JSONfile struct {
	filepath string
}

func (j *JSONfile)Get() ([]byte, error) {
	return ioutil.ReadFile(j.filepath)
}

func (i *IndexTracker) Exec(ctx context.Context) error {

	payload, err := i.Source.Get()
	if err != nil {
		return err
	}

	vals, err := util.ParsePayload(payload, i.Args)
	if err != nil {
		return err
	}

	//fmt.Printf("got value of %f for %s\n", val, i.Identifier )
	volume := 0.0
	if len(vals) >= 2 {
		volume = vals[1]
	}

	//save the value into our local data window (set 0 volume for now)
	apiOracle.SetRequestValue(i.Identifier, time.Now(), vals[0], volume)

	//update all the values that depend on these symbols
	return UpdatePSRs(ctx, i.Symbols)
}

func (i *IndexTracker) String() string {
	return fmt.Sprintf("%s on %s", strings.Join(i.Symbols, ","), i.Name)
}


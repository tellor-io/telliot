// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/benbjohnson/clock"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/apiOracle"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/yalp/jsonpath"
)

var clck clock.Clock

func init() {
	clck = clock.New()
}

var indexes map[string][]*IndexTracker

// BuildIndexTrackers creates and initializes a new tracker instance.
func BuildIndexTrackers() ([]Tracker, error) {
	err := apiOracle.EnsureValueOracle()
	if err != nil {
		return nil, err
	}

	cfg := config.GetConfig()

	indexPath := filepath.Join(cfg.ConfigFolder, "indexes.json")
	byteValue, err := ioutil.ReadFile(indexPath)
	if err != nil {
		return nil, errors.Wrapf(err, "read index file @ %s", indexPath)
	}
	var baseIndexes map[string][]TrackerInfo
	err = json.Unmarshal(byteValue, &baseIndexes)
	if err != nil {
		return nil, errors.Wrap(err, "parse index file")
	}

	indexes = make(map[string][]*IndexTracker)

	//build a tracker for each unique API
	indexers := make(map[string]*IndexTracker)
	//and keep track of which APIs influence which symbols so we know what to update later
	symbolsForAPI := make(map[string][]string)

	for symbol, apis := range baseIndexes {
		for _, api := range apis {
			// Tracker for this API already added?
			_, ok := indexers[api.URL]
			if !ok {

				// Expand any env variables with their values from the .env file.
				vars, err := godotenv.Read(cfg.EnvFile)
				// Ignore file doesn't exist errors.
				if _, ok := err.(*os.PathError); err != nil && !ok {
					return nil, errors.Wrap(err, "reading .env file")
				}
				api.URL = os.Expand(api.URL, func(key string) string {
					return vars[key]
				})

				var name string
				var source DataSource
				if strings.HasPrefix(api.URL, "http") {
					source = &JSONapi{&FetchRequest{queryURL: api.URL, timeout: cfg.FetchTimeout.Duration}}
					u, err := url.Parse(api.URL)
					if err != nil {
						return nil, errors.Wrapf(err, "invalid API URL: %s", api.URL)
					}
					name = u.Host
				} else {
					source = &JSONfile{filepath: filepath.Join(cfg.ConfigFolder, api.URL)}
					name = filepath.Base(api.URL)
				}
				indexers[api.URL] = &IndexTracker{
					Name:       name,
					Identifier: api.URL,
					Source:     source,
					JsonPath:   api.JsonPath,
				}
			}
			//now we definitely have one
			thisOne := indexers[api.URL]

			//insert add it and it's more specific variant to the symbol->api map
			indexes[symbol] = append(indexes[symbol], thisOne)
			specificName := fmt.Sprintf("%s~%s", symbol, thisOne.Name)
			indexes[specificName] = append(indexes[specificName], thisOne)

			//save this for later so we can build the api->symbol map
			symbolsForAPI[api.URL] = append(symbolsForAPI[api.URL], symbol, specificName)
		}
	}

	var sortedIndexers []string
	//set the reverse map
	for api, symbols := range symbolsForAPI {
		indexers[api].Symbols = symbols
		sortedIndexers = append(sortedIndexers, api)
	}

	// sort the Indexer array so we return the same order every time
	sort.Strings(sortedIndexers)

	//make an array of trackers to be sent to Runner
	trackers := make([]Tracker, len(indexers))
	for idx, api := range sortedIndexers {
		trackers[idx] = indexers[api]
	}

	//start the PSR system that will feed from these indexes
	err = InitPSRs()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize PSRs")
	}

	return trackers, nil
}

type TrackerInfo struct {
	URL         string `json:"URL"`
	TrackerType string `json:"type"`
	JsonPath    string `json:"jsonPath"`
}
type IndexTracker struct {
	Name       string
	Identifier string
	Symbols    []string
	Source     DataSource
	JsonPath   string
}

type DataSource interface {
	Get() ([]byte, error)
}

type JSONapi struct {
	Request *FetchRequest
}

func (j *JSONapi) Get() ([]byte, error) {
	return fetchWithRetries(j.Request)
}

type JSONfile struct {
	filepath string
}

func (j *JSONfile) Get() ([]byte, error) {
	return ioutil.ReadFile(j.filepath)
}

func (i *IndexTracker) Exec(ctx context.Context) error {
	payload, err := i.Source.Get()
	if err != nil {
		return err
	}
	vals := make([]float64, 0)

	var decodedPayload interface{}
	err = json.Unmarshal(payload, &decodedPayload)
	if err != nil {
		return err
	}
	match, err := jsonpath.Read(decodedPayload, i.JsonPath)
	if err != nil {
		return err
	}
	arr, ok := match.([]interface{})
	if !ok {
		return errors.Wrap(err, "JsonPath match need to be an array")
	}
	for _, a := range arr {
		val, _ := strconv.ParseFloat(fmt.Sprintf("%v", a), 64)
		vals = append(vals, val)
	}

	fmt.Fprintln(os.Stdout, fmt.Sprintf("got value of %f for %s\n", vals, i.Identifier))
	volume := 0.0
	if len(vals) >= 2 {
		volume = vals[1]
	}

	//save the value into our local data window (set 0 volume for now)
	apiOracle.SetRequestValue(i.Identifier, clck.Now(), apiOracle.PriceInfo{Price: vals[0], Volume: volume})
	//update all the values that depend on these symbols
	return UpdatePSRs(ctx, i.Symbols)
}

func (i *IndexTracker) String() string {
	return fmt.Sprintf("%s on %s", strings.Join(i.Symbols, ","), i.Name)
}

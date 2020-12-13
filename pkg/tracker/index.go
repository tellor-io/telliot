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

// parseIndexFile parses indexes.json file and returns a *IndexTracker,
// for every URL in index file, also a map[string][]string that describes which APIs
// influence which symbols.
func parseIndexFile() (trackersPerURL map[string]*IndexTracker, symbolsForAPI map[string][]string, err error) {
	cfg := config.GetConfig()

	// Load index file.
	indexFilePath := filepath.Join(cfg.ConfigFolder, "indexes.json")
	byteValue, err := ioutil.ReadFile(indexFilePath)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "read index file @ %s", indexFilePath)
	}
	// Parse to json.
	baseIndexes := make(map[string][]IndexObject)
	err = json.Unmarshal(byteValue, &baseIndexes)
	if err != nil {
		return nil, nil, errors.Wrap(err, "parse index file")
	}
	// Keep track of tracker per symbol.
	indexes = make(map[string][]*IndexTracker)
	// Build a tracker for each unique URL.
	trackersPerURL = make(map[string]*IndexTracker)
	// Keep track of which APIs influence which symbols so we know what to update later.
	symbolsForAPI = make(map[string][]string)

	for symbol, apis := range baseIndexes {
		for _, api := range apis {
			// Tracker for this API already added?
			_, ok := trackersPerURL[api.URL]
			if !ok {

				// Expand any env variables with their values from the .env file.
				vars, err := godotenv.Read(cfg.EnvFile)
				// Ignore file doesn't exist errors.
				if _, ok := err.(*os.PathError); err != nil && !ok {
					return nil, nil, errors.Wrap(err, "reading .env file")
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
						return nil, nil, errors.Wrapf(err, "invalid API URL: %s", api.URL)
					}
					name = u.Host
				} else {
					source = &JSONfile{filepath: filepath.Join(cfg.ConfigFolder, api.URL)}
					name = filepath.Base(api.URL)
				}
				trackersPerURL[api.URL] = &IndexTracker{
					Name:       name,
					Identifier: api.URL,
					Source:     source,
					JSONPath:   api.JSONPath,
				}
			}
			// Now we definitely have one.
			thisOne := trackersPerURL[api.URL]

			// Insert add it and it's more specific variant to the symbol -> api map.
			indexes[symbol] = append(indexes[symbol], thisOne)
			specificName := fmt.Sprintf("%s~%s", symbol, thisOne.Name)
			indexes[specificName] = append(indexes[specificName], thisOne)

			// Save this for later so we can build the api->symbol map.
			symbolsForAPI[api.URL] = append(symbolsForAPI[api.URL], symbol, specificName)
		}
	}
	return

}

// BuildIndexTrackers creates and initializes a new tracker instance.
func BuildIndexTrackers() ([]Tracker, error) {
	err := apiOracle.EnsureValueOracle()
	if err != nil {
		return nil, err
	}

	// Load trackers from the index file,
	// and build a tracker for each unique URL, symbol
	indexers, symbolsForAPI, err := parseIndexFile()
	if err != nil {
		return nil, err
	}

	var sortedIndexers []string
	// Set the reverse map.
	for api, symbols := range symbolsForAPI {
		indexers[api].Symbols = symbols
		sortedIndexers = append(sortedIndexers, api)
	}

	// Sort the Indexer array so we return the same order every time.
	sort.Strings(sortedIndexers)

	// Make an array of trackers to be sent to Runner.
	trackers := make([]Tracker, len(indexers))
	for idx, api := range sortedIndexers {
		trackers[idx] = indexers[api]
	}

	// Start the PSR system that will feed from these indexes.
	err = InitPSRs()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize PSRs")
	}

	return trackers, nil
}

// IndexObject will be used in parsing index file.
type IndexObject struct {
	URL      string `json:"URL"`
	JSONPath string `json:"JSONPath"`
}
type IndexTracker struct {
	Name       string
	Identifier string
	Symbols    []string
	Source     DataSource
	JSONPath   string
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

	vals, err := i.parsePayload(payload)
	if err != nil {
		return err
	}

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

// parsePayload parses the input JSON payload to a slice of float64
// The input JSON will get queried using JSONPath query language if
// the JSONPath expression is not empty.
func (i *IndexTracker) parsePayload(payload []byte) (vals []float64, err error) {

	var decodedPayload, result interface{}
	err = json.Unmarshal(payload, &decodedPayload)
	if err != nil {
		return
	}

	// Query the json payload using JSONPath expression if needed.
	result = decodedPayload
	if len(strings.TrimSpace(i.JSONPath)) > 0 {
		if err != nil {
			return
		}
		result, err = jsonpath.Read(decodedPayload, i.JSONPath)
		if err != nil {
			return
		}
	}

	// Expect result to be a slice of float or a single float value.
	var resultList []interface{}
	switch result := result.(type) {
	case []interface{}:
		resultList = result
	default:
		resultList = []interface{}{result}
	}

	// Parse each item of slice to a float.
	vals = make([]float64, 0)
	for _, a := range resultList {
		val, err := strconv.ParseFloat(fmt.Sprintf("%v", a), 64)
		if err != nil {
			return nil, errors.Wrap(err, "JSON value needs to be a valid float")
		}
		vals = append(vals, val)
	}
	return
}

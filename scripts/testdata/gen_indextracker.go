// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/tracker"
	"github.com/tellor-io/telliot/pkg/tracker/index"
	"github.com/tellor-io/telliot/pkg/util"
)

type IndexTrackerTestData struct {
	URL      string
	Param    string
	Payload  string
	Expected []float64
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmsgprefix)

	f, err := ioutil.ReadFile(filepath.Join("configs", "config.json"))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Config{}

	if err := json.Unmarshal(f, &cfg); err != nil {
		log.Fatal(err)
	}

	err = util.SetupLoggingConfig(cfg.Logger)
	if err != nil {
		log.Fatal(err)
	}
	DB, err := db.Open(filepath.Join(os.TempDir(), "testdata_gen"))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := index.BuildIndexTrackers(&cfg, DB); err != nil {
		log.Fatal(err)
	}
	indexes := tracker.GetIndexes()
	testdata := make(map[string][]*IndexTrackerTestData)
	for symbol, indexers := range indexes {
		// Skip on duplicates.
		if strings.Contains(symbol, "~") {
			continue
		}
		var testdataForSymbol = make([]*IndexTrackerTestData, 0)
		for _, indexer := range indexers {
			payload, err := indexer.Source.Get()
			if err != nil {
				log.Fatal(err)
			}
			expected, err := indexer.ParsePayload(payload)
			if err != nil {
				log.Fatal(err)
			}
			testdataForSymbol = append(testdataForSymbol, &IndexTrackerTestData{
				URL:      indexer.Identifier,
				Param:    indexer.Param,
				Payload:  string(payload),
				Expected: expected,
			})
		}
		testdata[symbol] = testdataForSymbol
	}
	testdataJSON, err := json.MarshalIndent(testdata, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filepath.Join("test", "tracker", "testdata", "test_api.json"), testdataJSON, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

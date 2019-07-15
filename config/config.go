package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//Config holds global config info derived from config.json
type Config struct {
	ContractAddress   string   `json:"contractAddress"`
	NodeURL           string   `json:"nodeURL"`
	PrivateKey        string   `json:"privateKey"`
	DatabaseURL       string   `json:"databaseURL"`
	PublicAddress     string   `json:"publicAddress"`
	TrackerSleepCycle uint     `json:"trackerCycle"` //in seconds
	Trackers          []string `json:"trackerTypes"`
}

var config *Config

func init() {
	path := os.Getenv("CONFIG")
	if len(path) == 0 {
		panic("Missing env CONFIG variable")
	}
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		panic(err.Error())
	}
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	fmt.Printf("config: %+v", config)
}

//GetConfig returns a shared instance of config
func GetConfig() (*Config, error) {
	if config == nil {
		return nil, fmt.Errorf("Config was not initialized")
	}
	return config, nil
}

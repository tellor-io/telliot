package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Contract_Address string `json:"contract_Address"`
	NodeURL          string `json:"nodeURL"`
	Private_Key      string `json:"private_Key"`
	DatabaseURL      string `json:"databaseURL"`
	Public_address   string `json:"public_address"`
}

func main() {
	var config Config
	configFile, err := os.Open("config.json")
	fmt.Println(configFile)
	defer configFile.Close()
	if err != nil {
		return
	}
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	fmt.Println(config)
	fmt.Println(config.DatabaseURL)
}

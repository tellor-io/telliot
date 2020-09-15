package config

import (
	"os"
	"path/filepath"
	"testing"
)

func createEnvFile(t *testing.T) {
	f, err := os.Create(".env")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString("ETH_PRIVATE_KEY=\"0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\"")
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
}

func TestConfig(t *testing.T) {
	//Creating a mock .ENV file to go around this issue with godotenv:
	//https://github.com/joho/godotenv/issues/43
	createEnvFile(t)
	dir, err := filepath.Abs("../config.json")
	if err != nil {
		t.Fatal("Error reading config directory")
	}
	err = ParseConfig(dir)
	if err != nil {
		println(err)
		t.Fatal("")
	}
	cfg := GetConfig()
	if len(cfg.ContractAddress) == 0 {
		t.Fatal("Config did not parse correctly")
	}
}

package cli

import (
	"log"
	"testing"
)

func TestFlags(t *testing.T) {
	flags := GetFlags()
	if len(flags.LoggingConfigPath) == 0 {
		t.Fatal("Missing logging config path")
	}
	if len(flags.ConfigPath) == 0 {
		t.Fatal("Missing config path")
	}
}

func TestDefaultPaths(t *testing.T) {
	flags := GetFlags()
	if len(flags.LoggingConfigPath) == 0 {
		t.Fatal("Default logging path not found")
	} else {
		log.Println("Logging path", flags.LoggingConfigPath)
	}
	if len(flags.PSRPath) == 0 {
		t.Fatal("Default PSR path not found")
	} else {
		log.Println("PSR Path", flags.PSRPath)
	}
}

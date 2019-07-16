package cli

import (
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

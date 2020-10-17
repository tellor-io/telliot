// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"testing"

	"github.com/tellor-io/TellorMiner/pkg/config"
)

func TestDB(t *testing.T) {
	config.OpenTestConfig(t)
	db, cleanup := OpenTestDB(t)
	defer t.Cleanup(cleanup)

	err := db.Put("sample", []byte("sample_value"))
	if err != nil {
		t.Error(err)
	}
	b, err := db.Has("sample")
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("Expected key to be present in DB")
	}
	v, err := db.Get("sample")
	if err != nil {
		t.Error(err)
	}
	s := string(v)
	if s != "sample_value" {
		t.Error("Get value doesn't match original: " + s + " != sample_value")
	}
	t.Log("Retrieved " + s)
}

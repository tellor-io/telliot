// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"testing"

	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestDB(t *testing.T) {
	cfg, err := config.OpenTestConfig("../..")
	testutil.Ok(t, err)
	db, cleanup, err := OpenTestDB(cfg)
	testutil.Ok(t, err)
	defer func() {
		testutil.Ok(t, cleanup())
	}()

	err = db.Put("sample", []byte("sample_value"))
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

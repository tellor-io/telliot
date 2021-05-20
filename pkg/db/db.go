// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"github.com/tellor-io/telliot/pkg/format"
)

const ComponentName = "db"

type Config struct {
	LogLevel string
	Path     string
	// Connect to this remote DB.
	RemoteHost    string
	RemotePort    uint
	RemoteTimeout format.Duration
}

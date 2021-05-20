// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/tellor-io/telliot/pkg/cli"
)

func main() {
	//lint:ignore faillint it should print to console
	fmt.Printf(cli.VersionMessage, cli.GitTag, cli.GitHash)
	ctx := kong.Parse(cli.Cli(), kong.Name("Telliot"),
		kong.Description("The official Tellor cli tool"),
		kong.UsageOnError())
	err := ctx.Run(*ctx)
	ctx.FatalIfErrorf(err)
}

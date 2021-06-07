// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/tellor-io/telliot/pkg/cli"
)

func main() {
	// Don't show the version message when it's an help command.
	shouldShowVersionMessage := true
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			shouldShowVersionMessage = false
			break
		}
	}
	if shouldShowVersionMessage {
		//lint:ignore faillint it should print to console
		fmt.Printf(cli.VersionMessage, cli.GitTag, cli.GitHash)
	}
	ctx := kong.Parse(cli.CLI, kong.Name("telliot"),
		kong.Description("The official Tellor cli tool"),
		kong.UsageOnError())

	ctx.FatalIfErrorf(ctx.Run(*ctx))
}

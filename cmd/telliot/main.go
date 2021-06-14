// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	"github.com/google/go-github/v35/github"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/cli"
)

var GitTag string
var GitHash string

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
		fmt.Printf(cli.VersionMessage, GitTag, GitHash)

		newRelease, err := checkNewVersion(GitTag)
		if err != nil {
			log.Printf("ERROR checking for a new release:%v", err.Error())
		}
		if newRelease != "" {
			log.Printf("THERE IS A NEW RELEASE: %v", newRelease)
		}
	}
	ctx := kong.Parse(&cli.CLI, kong.Name("telliot"),
		kong.Description("The official Tellor cli tool"),
		kong.UsageOnError())

	ctx.FatalIfErrorf(ctx.Run(*ctx))
}

func checkNewVersion(current string) (string, error) {
	if current == "" { // Can be empty when using `go run`.
		return "", nil
	}
	ctx, cncl := context.WithTimeout(context.Background(), 2*time.Second)
	defer cncl()
	client := github.NewClient(nil)
	release, resp, err := client.Repositories.GetLatestRelease(ctx, "tellor-io", "telliot")

	if err != nil {
		return "", errors.Wrap(err, "checking for a new release")
	}

	if resp.StatusCode/100 != 2 {
		rbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.Errorf("bad response status %v", resp.Status)
		}
		return "", errors.Errorf("bad response status %v from %q", resp.Status, string(rbody))
	}

	parts := strings.Split(current, "-")

	if len(parts) == 0 {
		return "", errors.New("failed to process current tag name")
	}

	if parts[0] != release.GetTagName() {
		return release.GetHTMLURL(), nil
	}

	return "", nil
}

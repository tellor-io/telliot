// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structtag"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/cmd/telliot/cli"
	"github.com/tellor-io/telliot/pkg/config"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type argument struct {
	Optional bool
	Help     string
}
type commandInfo struct {
	Arguments map[string]argument
	Help      string
}

type cfgDoc struct {
	Name    string
	Help    string
	Default interface{}
}

func main() {
	app := kingpin.New(filepath.Base(os.Args[0]), "Telliot config docs generator.")
	app.HelpFlag.Short('h')
	// outputFile := app.Flag("output", "Output file for the generated doc.").String()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	if _, err := app.Parse(os.Args[1:]); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	cliDocs := make(map[string]commandInfo)
	cli := cli.Cli()
	if err := generateCommand("", reflect.ValueOf(cli).Elem(), cliDocs); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "cli docs", "err", err)
		os.Exit(1)
	}

	var (
		envDocs map[string]string
		err     error
	)
	if envDocs, err = genEnvDocs(); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "env docs", "err", err)
		os.Exit(1)
	}

	cfgDocs := make(map[string]interface{})
	cfg := config.DefaultConfig()
	if err := genCfgDocs(reflect.ValueOf(cfg), cfgDocs); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "cli", "err", err)
		os.Exit(1)
	}
	spew.Dump(cfgDocs)
	spew.Dump(envDocs)
	spew.Dump(cliDocs)
	logger.Log("msg", "success")
}

func generateCommand(parent string, cli reflect.Value, docs map[string]commandInfo) error {
	for i := 0; i < cli.NumField(); i++ {
		v := cli.Field(i)
		t := cli.Type().Field(i)
		switch v.Kind() {
		case reflect.Struct:

			// If there is no child struct fields then v is a leaf.
			leafFound := true
			if v.Type().NumField() > 0 {
				// Checking the first field to know if it's a leaf.
				v0 := v.Type().Field(0)

				tags, err := structtag.Parse(string(v0.Tag))
				if err != nil {
					return errors.Wrapf(err, "%s: failed to parse tag %q", v.Type().Field(i).Name, v.Type().Field(i).Tag)
				}
				_, err = tags.Get("cmd")
				leafFound = err != nil
			}
			tags, err := structtag.Parse(string(t.Tag))
			if err != nil {
				return errors.Wrapf(err, "%s: failed to parse tag %q", v.Type().Field(i).Name, v.Type().Field(i).Tag)
			}
			tag, err := tags.Get("help")
			if err != nil {
				return errors.Wrapf(err, "help tag missing: %s", t.Name)
			}
			if leafFound {
				// v is a leaf in the cmd tree.
				cmdName := strings.ToLower(t.Name)
				if len(parent) > 0 {
					cmdName = fmt.Sprintf("%s %s", parent, cmdName)
				}
				docs[cmdName] = commandInfo{
					Arguments: getArguments(v),
					Help:      tag.Value(),
				}

			} else {
				parentName := strings.ToLower(t.Name)
				if len(parent) > 0 {
					parentName = fmt.Sprintf("%s %s", parent, parentName)
				}
				// Add top level command too.
				docs[parentName] = commandInfo{
					Arguments: map[string]argument{},
					Help:      tag.Value(),
				}
				if err := generateCommand(parentName, v, docs); err != nil {
					return errors.Wrapf(err, "%s", t.Name)
				}
			}

		case reflect.Ptr:
			return errors.New("nil pointers are not allowed in configuration")
		case reflect.Interface:

		}
	}
	return nil
}

func getArguments(cmd reflect.Value) map[string]argument {
	out := make(map[string]argument)
	for i := 0; i < cmd.NumField(); i++ {
		// v := cmd.Field(i)
		t := cmd.Type().Field(i)
		tags, err := structtag.Parse(string(t.Tag))
		if err != nil {
			fmt.Printf("err: %v\n", errors.Wrapf(err, "%s: failed to parse tag %q", t.Name, t.Tag))
			continue
		}
		_, err = tags.Get("arg")
		if err != nil {
			fmt.Printf("err: %v\n", errors.Wrapf(err, "help tag missing: %s", t.Name))
			continue
		}
		_, optionalErr := tags.Get("optional")
		helpText := ""
		help, _ := tags.Get("help")
		if help != nil {
			helpText = help.Value()
		}
		out[t.Name] = argument{
			Optional: optionalErr == nil,
			Help:     helpText,
		}

	}
	return out
}

func genCfgDocs(cfg reflect.Value, cfgDocs map[string]interface{}) error {
	for i := 0; i < cfg.NumField(); i++ {
		v := cfg.Field(i)
		t := cfg.Type().Field(i)
		switch v.Kind() {
		case reflect.Struct:
			cfgDocs[t.Name] = make(map[string]interface{})
			if err := genCfgDocs(v, cfgDocs[t.Name].(map[string]interface{})); err != nil {
				return err
			}
		default:
			doc := cfgDoc{
				Name:    t.Name,
				Default: v.Interface(),
			}
			tags, _ := structtag.Parse(string(t.Tag))
			if tags != nil {
				help, _ := tags.Get("help")
				if help != nil {
					doc.Help = help.Value()
				}
			}
			cfgDocs[t.Name] = doc
		}
	}
	return nil
}

func genEnvDocs() (map[string]string, error) {
	docs := make(map[string]string)
	bytes, err := ioutil.ReadFile("configs/.env.example")
	if err != nil {
		return nil, err
	}
	envExamples := strings.Split(string(bytes), "\n")
	for _, env := range envExamples {
		comment := strings.TrimSpace(strings.Split(env, "#")[1])
		name := strings.TrimSpace(strings.Split(env, "=")[0])
		docs[name] = comment
	}
	return docs, nil
}

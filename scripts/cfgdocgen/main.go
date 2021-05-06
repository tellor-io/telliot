// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/fatih/structtag"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/cmd/telliot/cli"
	"github.com/tellor-io/telliot/pkg/config"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type argument struct {
	Name     string
	Help     string
	Optional bool
}

type cliDoc struct {
	Name      string
	Help      string
	Arguments []argument
}

type cfgDoc struct {
	Name     string
	Help     string
	Default  interface{}
	Required bool
}

type envDoc struct {
	Name     string
	Help     string
	Required bool
}

var logger log.Logger

func main() {
	app := kingpin.New(filepath.Base(os.Args[0]), "Telliot config docs generator.")
	app.HelpFlag.Short('h')
	outputFile := app.Flag("output", "Output file for the generated doc.").String()

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	if _, err := app.Parse(os.Args[1:]); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	// Generating cli docs from the cli struct.
	cliDocsMap := make(map[string]cliDoc)
	cli := cli.Cli()
	if err := genCliDocs("", reflect.ValueOf(cli).Elem(), cliDocsMap); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "cli docs", "err", err)
		os.Exit(1)
	}
	cliDocs := make([]cliDoc, 0)
	keys := []string{}
	for k := range cliDocsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		cliDocs = append(cliDocs, cliDocsMap[k])
	}

	// Generating env docs from the .env.example file.
	var (
		envDocs []envDoc
		err     error
	)
	if envDocs, err = genEnvDocs(); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "env docs", "err", err)
		os.Exit(1)
	}

	// Generating config docs from the default config object.
	cfgDocsMap := make(map[string]cfgDoc)
	cfg := config.DefaultConfig()
	if err := genCfgDocs("", reflect.ValueOf(cfg), cfgDocsMap); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "cli", "err", err)
		os.Exit(1)
	}
	cfgDocs := make([]cfgDoc, 0)
	keys = []string{}
	for k := range cfgDocsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		cfgDocs = append(cfgDocs, cfgDocsMap[k])
	}
	tmpl := template.Must(template.ParseFiles("scripts/cfgdocgen/configuration.md"))
	outf, err := os.Open(*outputFile)
	if err != nil {
		level.Error(logger).Log("msg", "failed to open output file, redirecting to stdout", "err", err, "output", *outputFile)
		outf = os.Stdout
	}
	err = tmpl.Execute(outf,
		struct {
			CliDocs []cliDoc
			EnvDocs []envDoc
			CfgDocs []cfgDoc
		}{
			CliDocs: cliDocs,
			EnvDocs: envDocs,
			CfgDocs: cfgDocs,
		})
	if err != nil {
		level.Error(logger).Log("msg", "failed to execute template", "err", err)
		os.Exit(1)
	}
	logger.Log("msg", "success")
}

func genCliDocs(parent string, cli reflect.Value, docs map[string]cliDoc) error {
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
				docs[cmdName] = cliDoc{
					Name:      cmdName,
					Arguments: getArguments(v),
					Help:      tag.Value(),
				}

			} else {
				parentName := strings.ToLower(t.Name)
				if len(parent) > 0 {
					parentName = fmt.Sprintf("%s %s", parent, parentName)
				}
				// Add top level command too.
				docs[parentName] = cliDoc{
					Name:      parentName,
					Arguments: []argument{},
					Help:      tag.Value(),
				}
				if err := genCliDocs(parentName, v, docs); err != nil {
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

func getArguments(cmd reflect.Value) []argument {
	out := make([]argument, 0)
	for i := 0; i < cmd.NumField(); i++ {
		// v := cmd.Field(i)
		t := cmd.Type().Field(i)
		tags, err := structtag.Parse(string(t.Tag))
		if err != nil {
			level.Debug(logger).Log("msg", "failed to parse tag", "field", t.Name, "tag", t.Tag, "err", err)
			continue
		}
		_, err = tags.Get("arg")
		if err != nil {
			level.Debug(logger).Log("msg", "help tag missing", "field", t.Name, "err", err)
			continue
		}
		_, optionalErr := tags.Get("optional")
		helpText := ""
		help, _ := tags.Get("help")
		if help != nil {
			helpText = help.Value()
		}
		out = append(out, argument{
			Name:     t.Name,
			Optional: optionalErr == nil,
			Help:     helpText,
		})

	}
	return out
}

func genCfgDocs(parent string, cfg reflect.Value, cfgDocs map[string]cfgDoc) error {
	for i := 0; i < cfg.NumField(); i++ {
		v := cfg.Field(i)
		t := cfg.Type().Field(i)
		switch v.Kind() {
		case reflect.Struct:
			parentName := t.Name
			if len(parent) > 0 {
				parentName = fmt.Sprintf("%s.%s", parent, parentName)
			}
			if err := genCfgDocs(parentName, v, cfgDocs); err != nil {
				return err
			}
		default:
			name := t.Name
			if len(parent) > 0 {
				name = fmt.Sprintf("%s.%s", parent, name)
			}
			doc := cfgDoc{
				Name:    name,
				Default: v.Interface(),
			}
			tags, _ := structtag.Parse(string(t.Tag))
			if tags != nil {
				help, _ := tags.Get("help")
				if help != nil {
					doc.Help = help.Value()
				}
			}
			cfgDocs[name] = doc
		}
	}
	return nil
}

func genEnvDocs() ([]envDoc, error) {
	docs := make([]envDoc, 0)
	bytes, err := ioutil.ReadFile("configs/.env.example")
	if err != nil {
		return nil, err
	}
	envExamples := strings.Split(string(bytes), "\n")
	for _, env := range envExamples {
		var (
			help     string
			required bool
		)
		comment := strings.TrimSpace(strings.Split(env, "#")[1])
		help = comment

		parts := strings.Fields(comment)
		if len(parts) > 0 && parts[0] == "required" {
			required = true
			help = strings.TrimSpace(strings.TrimPrefix(comment, "required"))
		}

		name := strings.TrimSpace(strings.Split(env, "=")[0])
		docs = append(docs, envDoc{
			Name:     name,
			Help:     help,
			Required: required,
		})
	}
	return docs, nil
}

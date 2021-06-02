// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/fatih/structtag"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/cli"
	"github.com/tellor-io/telliot/pkg/config"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type cliOutput struct {
	Name      string
	CmdOutput string
}

type cfgDoc struct {
	Name     string
	Help     string
	Default  interface{}
	Required bool
}

func (c *cfgDoc) String() string {
	return fmt.Sprintf("(Required: %v) %s - Default: %v", c.Required, c.Help, c.Default)
}

type envDoc struct {
	Name     string
	Help     string
	Required bool
}

func main() {
	app := kingpin.New(filepath.Base(os.Args[0]), "Telliot config docs generator.")
	app.HelpFlag.Short('h')
	outputFile := app.Flag("output", "Output file for the generated doc.").String()
	cliBin := app.Flag("cli-bin", "Cli binary for generating command outputs.").Required().String()

	var err error
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	if _, err = app.Parse(os.Args[1:]); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	// Generating cli docs from the cli struct.
	cliDocsMap := make(map[string]string)
	cli := cli.Cli()
	if err = NewCliDocsGenerator(logger, *cliBin).genCliDocs("", reflect.ValueOf(cli).Elem(), cliDocsMap); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "cli docs", "err", err)
		os.Exit(1)
	}
	cliDocs := make([]cliOutput, 0)
	keys := []string{}
	for k := range cliDocsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		cliDocs = append(cliDocs, cliOutput{
			Name:      k,
			CmdOutput: cliDocsMap[k],
		})
	}
	// Generating env docs from the .env.example file.
	var (
		envDocs []envDoc
	)
	if envDocs, err = genEnvDocs(); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "env docs", "err", err)
		os.Exit(1)
	}

	// Generating config docs from the default config object.
	cfgDocsMap := make(map[string]interface{})
	cfgMap := make(map[string]interface{})
	cfg := config.DefaultConfig()
	if err := genCfgDocs(reflect.ValueOf(cfg), cfgDocsMap, cfgMap); err != nil {
		level.Error(logger).Log("msg", "failed to generate", "type", "cli", "err", err)
		os.Exit(1)
	}
	// Converto to json
	cfgDocs, err := json.MarshalIndent(cfgDocsMap, "", "\t")
	if err != nil {
		level.Error(logger).Log("msg", "marshaling config docs to json", "err", err)
		os.Exit(1)
	}
	defCfg, err := json.MarshalIndent(cfgMap, "", "\t")
	if err != nil {
		level.Error(logger).Log("msg", "marshaling default config to json", "err", err)
		os.Exit(1)
	}

	tmpl := template.Must(template.ParseFiles("scripts/cfgdocgen/configuration.md"))
	outf, err := os.Create(*outputFile)
	if err != nil {
		level.Error(logger).Log("msg", "failed to open output file, redirecting to stdout", "err", err, "output", *outputFile)
		outf = os.Stdout
	}
	err = tmpl.Execute(outf,
		struct {
			CliDocs []cliOutput
			EnvDocs []envDoc
			CfgDocs string
			Cfg     string
		}{
			CliDocs: cliDocs,
			EnvDocs: envDocs,
			CfgDocs: string(cfgDocs),
			Cfg:     string(defCfg),
		})
	if err != nil {
		level.Error(logger).Log("msg", "failed to execute template", "err", err)
		os.Exit(1)
	}
	logger.Log("msg", "success")
}

func NewCliDocsGenerator(logger log.Logger, cliBin string) *cliDocsGenerator {
	return &cliDocsGenerator{logger, cliBin}
}

type cliDocsGenerator struct {
	logger log.Logger
	cliBin string
}

func (self *cliDocsGenerator) cmdOutput(args string) string {
	_args := strings.Split(args, " ")
	_args = append(_args, "--help")
	cmd := exec.Command(self.cliBin, _args...)
	stdout, err := cmd.Output()
	if err != nil {
		level.Error(self.logger).Log("msg", "failed to execute telliot command", "err", err, "args", args, "cli", self.cliBin)
		os.Exit(1)
	}
	return string(stdout)

}

func (self *cliDocsGenerator) genCliDocs(parent string, cli reflect.Value, docs map[string]string) error {
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

			if leafFound {
				// v is a leaf in the cmd tree.
				cmdName := strings.ToLower(t.Name)
				if len(parent) > 0 {
					cmdName = fmt.Sprintf("%s %s", parent, cmdName)
				}
				docs[cmdName] = self.cmdOutput(cmdName)
			} else {
				parentName := strings.ToLower(t.Name)
				if len(parent) > 0 {
					parentName = fmt.Sprintf("%s %s", parent, parentName)
				}
				// Add top level command too.
				docs[parentName] = self.cmdOutput(parentName)
				if err := self.genCliDocs(parentName, v, docs); err != nil {
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

func genCfgDocs(cfg reflect.Value, cfgDocs map[string]interface{}, defCfg map[string]interface{}) error {
	for i := 0; i < cfg.NumField(); i++ {
		v := cfg.Field(i)
		t := cfg.Type().Field(i)
		switch v.Kind() {
		case reflect.Struct:
			cfgDocs[t.Name] = make(map[string]interface{})
			childDoc := (cfgDocs[t.Name]).(map[string]interface{})
			childCfg := (cfgDocs[t.Name]).(map[string]interface{})
			if err := genCfgDocs(v, childDoc, childCfg); err != nil {
				return err
			}
		default:
			name := t.Name
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
			cfgDocs[name] = doc.String()
			defCfg[name] = doc.Default
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

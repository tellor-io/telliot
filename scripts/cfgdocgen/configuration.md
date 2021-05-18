---
description: Telliot tweaks and settings to keep your rig running smoothly.
---

# Configuration reference

## CLI reference

Telliot commands and config file options are as the following:

#### Required Flags <a id="docs-internal-guid-d1a57725-7fff-a753-9236-759dd3f42eed"></a>

* `--config` \(path to your config file.\)

#### Telliot Commands
{{range .CliDocs}}

* `{{ .Name }}` {{/*Ranging over the current cli arguments and writing name, optional tag and help for each argument if any, we also will write the cli help after this loop*/}}{{range .Arguments }}\({{.Name}}{{if .Optional}}\(optional\){{end}}{{if .Help}}: {{ .Help }}{{end}}\){{end}}  \({{ .Help }}\)
{{end}}
#### .env file options:

{{range .EnvDocs}}
* `{{ .Name }}` {{if .Required }}\(required\){{end}} - {{ .Help }}
{{end}}

#### Config file options:
{{range .CfgDocs}}
* `{{ .Name }}` - {{if .Required }}\(required\){{end}} {{if .Default }}\(default: {{.Default}}\) - {{end}}{{ .Help }}
{{end}}
### LogConfig file options

The logging.config file consists of two fields: \* component \* level

The component is the package.component combination.

E.G. the Runner component in the tracker package would be: tracker.Runner

To turn on logging, add the component and the according level. Note the default level is "INFO", so to turn down the number of logs, enter "WARN" or "ERROR"

DEBUG - logs everything in INFO and additional developer logs

INFO - logs most information about the mining operation

WARN - logs all warnings and errors

ERROR - logs only serious errors

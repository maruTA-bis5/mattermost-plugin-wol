// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "net.bis5.mattermost.plugins.wol",
  "name": "Wake-on-LAN Plugin",
  "description": "This plugin generates the Wake on LAN magic packet.",
  "version": "0.2.0",
  "min_server_version": "5.18.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "TriggerWord",
        "display_name": "Trigger Word",
        "type": "text",
        "help_text": "",
        "placeholder": "",
        "default": "wol"
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}

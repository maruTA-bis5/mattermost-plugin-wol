/*
 * Mattermost Wake-on-LAN Plugin
 * Copyright 2020 Takayuki Maruyama
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (p *Plugin) executeHelp(fields []string) (*model.CommandResponse, *model.AppError) {
	if len(fields) >= 3 {
		return p.executeHelpSpecifiedCommand(fields[2])
	}
	trigger := p.getTriggerWord()
	text := fmt.Sprintf("Usage:\n/%s command [arguments]\n", trigger)
	text += "\nThe commands are:\n"
	text += "- add\n    - Register alias name for MAC address\n"
	text += "- list\n    - Get a list of registered MAC address & alias names\n"
	text += "- remove\n    - Remove alias name\n"
	text += "- wake\n    - Send Wake on LAN magic packet\n"
	text += fmt.Sprintf("\nUse `/%s help <command>` for more information about the command.", trigger)
	return p.newCommandResponse(text)
}

func (p *Plugin) executeHelpSpecifiedCommand(command string) (*model.CommandResponse, *model.AppError) {
	switch command {
	case "add":
		return p.newCommandResponse(p.createAddHelpText())
	case "list":
		return p.newCommandResponse(p.createListHelpText())
	case "remove":
		return p.newCommandResponse(p.createRemoveHelpText())
	case "wake":
		return p.newCommandResponse(p.createWakeHelpText())
	}
	return p.executeHelp([]string{})
}
func (p *Plugin) createAddHelpText() string {
	return "`add <MAC Address> <alias>`\nMACアドレスの別名を登録します"
}
func (p *Plugin) createListHelpText() string {
	return "`list`\nあなたが登録した別名とMACアドレスの一覧を表示します"
}
func (p *Plugin) createRemoveHelpText() string {
	return "`remove <alias>`\n別名の登録を削除します"
}
func (p *Plugin) createWakeHelpText() string {
	return "`wake <MAC Address or alias>`\n指定したMACアドレスまたは別名が示すアドレスに対してマジックパケットを送信します"
}

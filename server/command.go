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
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

func (p *Plugin) createCommand(triggerWord string) *model.Command {
	return &model.Command{
		Trigger:          triggerWord,
		AutoComplete:     true,
		AutoCompleteDesc: "available commands: wake, add, list, remove, help",
		AutoCompleteHint: "[command]",
	}
}

// ExecuteCommand hook
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	fields := strings.Fields(args.Command)
	var subCommand = ""
	if len(fields) > 1 {
		subCommand = fields[1]
	}

	switch subCommand {
	case "wake":
		return p.executeWake(c, args, fields)
	case "add":
		return p.executeAdd(c, args, fields)
	case "list":
		return p.executeList(c, args, fields)
	case "remove":
		return p.executeRemove(c, args, fields)
	default:
		return p.executeHelp(fields)
	}
}

func (p *Plugin) newCommandResponse(text string) (*model.CommandResponse, *model.AppError) {
	return &model.CommandResponse{
		ResponseType: model.CommandResponseTypeEphemeral,
		Text:         text,
	}, nil
}

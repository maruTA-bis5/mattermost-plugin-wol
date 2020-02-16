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

	"github.com/mattermost/mattermost-server/v5/mlog"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) executeList(c *plugin.Context, args *model.CommandArgs, fields []string) (*model.CommandResponse, *model.AppError) {
	// fetch
	entries, err := p.loadEntries(args.UserId)
	if err != nil {
		p.API.LogError("Could not load entries for user", mlog.String("userID", args.UserId), mlog.Err(err))
		return p.newCommandResponse("Could load alias entries.")
	}
	if len(entries) == 0 {
		return p.newCommandResponse("No alias entries found.")
	}
	text := "### Available aliases\n"
	for _, e := range entries {
		text += fmt.Sprintf("- %s (%s)\n", e.Name, e.MacAddress)
	}
	// print
	return p.newCommandResponse(text)
}

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
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

func (p *Plugin) executeRemove(c *plugin.Context, args *model.CommandArgs, fields []string) (*model.CommandResponse, *model.AppError) {
	if len(fields) <= 2 {
		return p.newCommandResponse("削除する別名を指定してください")
	}
	alias := fields[2]
	// fetch
	entries, err := p.loadEntries(args.UserId)
	if err != nil {
		p.API.LogError("Could not load entries for user", "userID", args.UserId, "error", err)
		return p.newCommandResponse("Could load alias entries.")
	}

	// remove
	_, found := p.findEntry(entries, alias)
	if !found {
		return p.newCommandResponse("別名[" + alias + "]が見つかりません")
	}
	delete(entries, alias)

	err = p.storeEntries(args.UserId, entries)
	if err != nil {
		p.API.LogError("store error", "error", err)
		return p.newCommandResponse("store error. err=" + err.Error())
	}

	return p.newCommandResponse("別名[" + alias + "]を削除しました")
}

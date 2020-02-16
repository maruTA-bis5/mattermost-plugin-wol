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
	"net"

	"github.com/mattermost/mattermost-server/v5/mlog"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) executeAdd(c *plugin.Context, args *model.CommandArgs, fields []string) (*model.CommandResponse, *model.AppError) {
	if len(fields) <= 3 {
		return p.executeHelpSpecifiedCommand("add")
	}
	macAddress, err := net.ParseMAC(fields[2])
	if err != nil {
		p.API.LogInfo("Could not parse mac address", mlog.String("address", fields[2]))
		return p.newCommandResponse("エラー: 不正なMACアドレスが指定されました")
	}
	alias := fields[3]
	entry := Entry{
		Name:       alias,
		MacAddress: macAddress.String(),
	}
	// store
	entries, err := p.loadEntries(args.UserId)
	if err != nil {
		p.API.LogError("Could not load entries for user", mlog.String("userID", args.UserId), mlog.Err(err))
		return p.newCommandResponse("Could not load alias entriee")
	}
	entries[alias] = entry
	err = p.storeEntries(args.UserId, entries)
	if err != nil {
		p.API.LogError("Could not store entries", mlog.String("userID", args.UserId), mlog.Err(err))
		return p.newCommandResponse("Could not store entries. err=" + err.Error())
	}

	return p.newCommandResponse("[" + macAddress.String() + "]の別名として[" + alias + "]を登録しました")
}

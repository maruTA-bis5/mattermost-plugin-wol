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
	"github.com/mdlayher/wol"
)

func (p *Plugin) executeWake(c *plugin.Context, args *model.CommandArgs, fields []string) (*model.CommandResponse, *model.AppError) {
	if len(fields) <= 2 {
		return p.newCommandResponse("MACアドレスまたは別名を指定してください")
	}
	aliasOrMacAddress := fields[2]
	target, err := net.ParseMAC(fields[2])
	if err != nil {
		entry, result := p.findEntryForUserForName(args.UserId, aliasOrMacAddress)
		if !result {
			return p.newCommandResponse("エラー: 不正なMACアドレスが指定されました") // XXX MACアドレスだけじゃないんだけど
		}
		target, err = net.ParseMAC(entry.MacAddress)
	}
	client, err := wol.NewClient()
	if err != nil {
		p.API.LogError("Could not initialize wol client", mlog.Err(err))
		return p.newCommandResponse("エラー: wol.NewClient() failed")
	}
	defer client.Close()
	address := "255.255.255.255:7" // 仮
	var password []byte
	client.WakePassword(address, target, password)
	return p.newCommandResponse("MACアドレス[" + target.String() + "]のPCを起動するためのマジックパケットを送出しました")
}

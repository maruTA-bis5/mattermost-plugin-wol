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
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/mdlayher/wol"
)

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
	default:
		return p.executeHelp(c, args, fields)
	}
}

func (p *Plugin) executeWake(c *plugin.Context, args *model.CommandArgs, fields []string) (*model.CommandResponse, *model.AppError) {
	if len(fields) <= 2 {
		// TODO
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         "MACアドレスを指定して下さい",
		}, nil
	}
	target, err := net.ParseMAC(fields[2])
	if err != nil {
		// TODO
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         "エラー: 不正なMACアドレスが指定されました",
		}, nil
	}
	client, err := wol.NewClient()
	if err != nil {
		// TODO
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         "エラー: wol.NewClient() failed",
		}, nil
	}
	defer client.Close()
	address := "255.255.255.255:7" // 仮
	var password []byte
	client.WakePassword(address, target, password)
	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         "MACアドレス[" + target.String() + "]のPCを起動するためのマジックパケットを送出しました",
	}, nil
}

func (p *Plugin) executeHelp(c *plugin.Context, args *model.CommandArgs, fields []string) (*model.CommandResponse, *model.AppError) {
	// TODO
	// wake 指定したアドレスに対してWake-on-LANのためのマジックパケットを送出します
	// ※例: /wol wake ff:ff:ff:ff:ff:ff",
	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         "TODO help text. `/wol wake [MAC ADDRESS]`",
	}, nil
}

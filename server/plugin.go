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
	"sync"

	"github.com/mattermost/mattermost-server/v6/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	activated bool
}

// OnActivate hook
func (p *Plugin) OnActivate() error {
	triggerWord := p.getConfiguration().TriggerWord
	err := p.API.RegisterCommand(p.createCommand(triggerWord))
	if err != nil {
		p.API.LogError("Could not register command", "error", err)
		return err
	}
	p.setActivated(true)
	return nil
}

// OnDeactivate hook
func (p *Plugin) OnDeactivate() error {
	p.setActivated(false)
	return nil
}

func (p *Plugin) isActivated() bool {
	return p.activated
}

func (p *Plugin) setActivated(activated bool) {
	p.activated = activated
}

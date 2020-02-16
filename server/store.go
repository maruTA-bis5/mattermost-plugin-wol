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
	"github.com/mattermost/mattermost-server/v5/mlog"
)

// Entry to send magic packet
type Entry struct {
	Name       string
	MacAddress string
}

func (p *Plugin) findEntry(entries map[string]Entry, alias string) (Entry, bool) {
	entry, ok := entries[alias]
	return entry, ok
}

func (p *Plugin) findEntryForUserForName(userID, alias string) (Entry, bool) {
	entries, err := p.loadEntries(userID)
	if err != nil {
		p.API.LogError("Could not load entries for user", mlog.String("userID", userID), mlog.Err(err))
		return Entry{}, false
	}
	return p.findEntry(entries, alias)
}

func (p *Plugin) loadEntries(userID string) (map[string]Entry, error) {
	var entries map[string]Entry
	exists, err := p.Helpers.KVGetJSON(userID, &entries)
	if !exists {
		entries = make(map[string]Entry)
	}
	return entries, err
}

func (p *Plugin) storeEntries(userID string, entries map[string]Entry) error {
	return p.Helpers.KVSetJSON(userID, entries)
}

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
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-server/v6/model"
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
		p.API.LogError("Could not load entries for user", "userID", userID, "error", err)
		return Entry{}, false
	}
	return p.findEntry(entries, alias)
}

func (p *Plugin) loadEntries(userID string) (map[string]Entry, *model.AppError) {
	var entries map[string]Entry
	bytes, err := p.API.KVGet(userID)
	if err != nil {
		p.API.LogWarn("loadEntries: error", "err", err)
		entries = make(map[string]Entry)
	}
	if bytes == nil && err == nil { // not found
		entries = make(map[string]Entry)
	} else {
		err := json.Unmarshal(bytes, &entries)
		if err != nil {
			return nil, model.NewAppError("loadEntries", "Could not unmarshal entries", nil, err.Error(), http.StatusInternalServerError).Wrap(err)
		}
	}
	return entries, err
}

func (p *Plugin) storeEntries(userID string, entries map[string]Entry) *model.AppError {
	json, err := json.Marshal(entries)
	if err != nil {
		return model.NewAppError("storeEntries", "Could not marshal entries", nil, err.Error(), http.StatusInternalServerError).Wrap(err)
	}
	setErr := p.API.KVSet(userID, json)
	return setErr
}

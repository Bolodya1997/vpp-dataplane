// Copyright (C) 2019 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpplink

import (
	"fmt"

	"git.fd.io/govpp.git/adapter"
	"git.fd.io/govpp.git/adapter/statsclient"
)

func Statsclientfunc(sc *statsclient.StatsClient) (ifNames adapter.NameStat, dumpStats []adapter.StatEntry, err error) {

	dumpStatsNames, err := sc.DumpStats("/if/names")
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, nil, err
	}
	ifNames = dumpStatsNames[0].Data.(adapter.NameStat)

	dumpStats, err = sc.DumpStats("/if/")
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, nil, err
	}
	return ifNames, dumpStats, nil
}

// SPDX-License-Identifier: Apache-2.0
//
// The OpenSearch Contributors require contributions made to
// this file be licensed under the Apache-2.0 license or a
// compatible open source license.
//
// Modifications Copyright OpenSearch Contributors. See
// GitHub history for details.

// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package opensearchapi

import (
	"strconv"
	"strings"
	"time"
)

// CatNodesParams represents possible parameters for the CatNodesReq
type CatNodesParams struct {
	Bytes                   string
	FullID                  *bool
	H                       []string
	IncludeUnloadedSegments *bool
	Local                   *bool
	MasterTimeout           time.Duration
	ClusterManagerTimeout   time.Duration
	Sort                    []string
	Time                    string
	V                       *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
}

func (r CatNodesParams) get() map[string]string {
	params := make(map[string]string)

	if r.Bytes != "" {
		params["bytes"] = r.Bytes
	}

	if r.FullID != nil {
		params["full_id"] = strconv.FormatBool(*r.FullID)
	}

	if len(r.H) > 0 {
		params["h"] = strings.Join(r.H, ",")
	}

	if r.IncludeUnloadedSegments != nil {
		params["include_unloaded_segments"] = strconv.FormatBool(*r.IncludeUnloadedSegments)
	}

	if r.Local != nil {
		params["local"] = strconv.FormatBool(*r.Local)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.ClusterManagerTimeout != 0 {
		params["cluster_manager_timeout"] = formatDuration(r.ClusterManagerTimeout)
	}

	if len(r.Sort) > 0 {
		params["s"] = strings.Join(r.Sort, ",")
	}

	if r.Time != "" {
		params["time"] = r.Time
	}

	if r.V != nil {
		params["v"] = strconv.FormatBool(*r.V)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	params["format"] = "json"

	return params
}

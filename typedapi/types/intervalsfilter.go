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


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


package types

// IntervalsFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_types/query_dsl/fulltext.ts#L74-L86
type IntervalsFilter struct {
	After          *Intervals `json:"after,omitempty"`
	Before         *Intervals `json:"before,omitempty"`
	ContainedBy    *Intervals `json:"contained_by,omitempty"`
	Containing     *Intervals `json:"containing,omitempty"`
	NotContainedBy *Intervals `json:"not_contained_by,omitempty"`
	NotContaining  *Intervals `json:"not_containing,omitempty"`
	NotOverlapping *Intervals `json:"not_overlapping,omitempty"`
	Overlapping    *Intervals `json:"overlapping,omitempty"`
	Script         *Script    `json:"script,omitempty"`
}

// NewIntervalsFilter returns a IntervalsFilter.
func NewIntervalsFilter() *IntervalsFilter {
	r := &IntervalsFilter{}

	return r
}

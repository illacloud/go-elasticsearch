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

// CompletionSuggestOption type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_global/search/_types/suggester.ts#L73-L84
type CompletionSuggestOption struct {
	CollateMatch *bool                  `json:"collate_match,omitempty"`
	Contexts     map[string][]Context   `json:"contexts,omitempty"`
	Fields       map[string]interface{} `json:"fields,omitempty"`
	Id_          *string                `json:"_id,omitempty"`
	Index_       *string                `json:"_index,omitempty"`
	Routing_     *string                `json:"_routing,omitempty"`
	Score        *float64               `json:"score,omitempty"`
	Score_       *float64               `json:"_score,omitempty"`
	Source_      interface{}            `json:"_source,omitempty"`
	Text         string                 `json:"text"`
}

// NewCompletionSuggestOption returns a CompletionSuggestOption.
func NewCompletionSuggestOption() *CompletionSuggestOption {
	r := &CompletionSuggestOption{
		Contexts: make(map[string][]Context, 0),
		Fields:   make(map[string]interface{}, 0),
	}

	return r
}

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

import (
	"encoding/json"
	"fmt"
)

// MultiTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_types/aggregations/Aggregate.ts#L464-L468
type MultiTermsBucket struct {
	Aggregations            map[string]Aggregate `json:"-"`
	DocCount                int64                `json:"doc_count"`
	DocCountErrorUpperBound *int64               `json:"doc_count_error_upper_bound,omitempty"`
	Key                     []FieldValue         `json:"key"`
	KeyAsString             *string              `json:"key_as_string,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s MultiTermsBucket) MarshalJSON() ([]byte, error) {
	type opt MultiTermsBucket
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Aggregations {
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewMultiTermsBucket returns a MultiTermsBucket.
func NewMultiTermsBucket() *MultiTermsBucket {
	r := &MultiTermsBucket{
		Aggregations: make(map[string]Aggregate, 0),
	}

	return r
}

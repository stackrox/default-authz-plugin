/*
Copyright 2019 StackRox Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package jsonutil

import "encoding/json"

// ToRaw converts an arbitrary Go object into its most generic form, i.e., any string-based type gets converted to
// `string`, any number type gets converted to `float64`, any object/struct type gets converted to
// `map[string]interface{}`, and every list type gets converted to `[]interface{}`.
func ToRaw(x interface{}) (interface{}, error) {
	jsonBytes, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	var raw interface{}
	if err := json.Unmarshal(jsonBytes, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

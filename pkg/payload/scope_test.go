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

package payload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// TestAccessScope_MarshalUnmarshal tests that access scopes are marshalled as expected, and that the resulting JSON
// can be unmarshalled to obtain the original struct.
func TestAccessScope_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		scope        AccessScope
		expectedJSON string
	}{
		{
			scope: AccessScope{
				Verb: "view",
				Noun: "Deployment",
			},
			expectedJSON: `{"verb":"view","noun":"Deployment"}`,
		},
		{
			scope: AccessScope{
				Verb: "view",
				Noun: "Deployment",
				Attributes: NounAttributes{
					Cluster: Cluster{
						ID:   "0",
						Name: "cluster0",
					},
				},
			},
			expectedJSON: `{"verb":"view","noun":"Deployment","attributes":{"cluster":{"id":"0","name":"cluster0"}}}`,
		},
		{
			scope: AccessScope{
				Verb: "view",
				Noun: "Deployment",
				Attributes: NounAttributes{
					Cluster: Cluster{
						ID:   "0",
						Name: "cluster0",
					},
					Namespace: "ns1",
				},
			},
			expectedJSON: `{"verb":"view","noun":"Deployment","attributes":{"cluster":{"id":"0","name":"cluster0"},"namespace":"ns1"}}`,
		},
		{
			scope: AccessScope{
				Verb: "view",
				Noun: "Deployment",
				Attributes: NounAttributes{
					Namespace: "ns1",
				},
			},
			expectedJSON: `{"verb":"view","noun":"Deployment","attributes":{"namespace":"ns1"}}`,
		},
	}

	for i, testCase := range cases {
		tc := testCase
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			jsonBytes, err := json.Marshal(tc.scope)
			if err != nil {
				t.Fatalf("Could not marshal access scope to JSON: %v", err)
			}
			jsonBytesFromPtr, err := json.Marshal(&tc.scope)
			if err != nil {
				t.Fatalf("Could not marshal access scope to JSON: %v", err)
			}

			if !bytes.Equal(jsonBytes, jsonBytesFromPtr) {
				t.Fatalf("JSON marshalling is not consistent for pointer vs non-pointer")
			}
			jsonStr := string(jsonBytes)
			if jsonStr != tc.expectedJSON {
				t.Fatalf("Marshalling access scope %+v did not yield expected JSON result %q, but %q", tc.scope, tc.expectedJSON, jsonStr)
			}

			var unmarshalled AccessScope
			if err := json.Unmarshal(jsonBytes, &unmarshalled); err != nil {
				t.Fatalf("Could not unmarshal previously marshalled JSON object: %v", err)
			}

			if !reflect.DeepEqual(tc.scope, unmarshalled) {
				t.Fatalf("Marshalling and unmarshalling %+v produced distinct object %+v", tc.scope, unmarshalled)
			}
		})
	}
}

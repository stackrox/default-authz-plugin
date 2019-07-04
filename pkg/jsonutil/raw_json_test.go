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

import (
	"reflect"
	"testing"
)

func TestToRaw_Struct(t *testing.T) {
	testObj := struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}{
		Foo: "qux",
		Bar: 42,
	}

	raw, err := ToRaw(testObj)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expected := map[string]interface{}{
		"foo": "qux",
		"bar": float64(42),
	}

	if !reflect.DeepEqual(raw, expected) {
		t.Fatalf("Object %#v is not equal to expected %#v", raw, expected)
	}
}

func TestToRaw_Slice(t *testing.T) {
	testSlice := []string{"foo", "bar"}

	raw, err := ToRaw(testSlice)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expected := []interface{}{"foo", "bar"}

	if !reflect.DeepEqual(raw, expected) {
		t.Fatalf("Object %#v is not equal to expected %#v", raw, expected)
	}
}

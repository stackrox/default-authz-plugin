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

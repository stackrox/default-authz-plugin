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

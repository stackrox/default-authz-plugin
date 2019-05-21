package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// LoadServerConfig loads the server configuration (in JSON format) from the given file.
func LoadServerConfig(path string) (*ServerConfig, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg ServerConfig
	if err := json.Unmarshal(contents, &cfg); err != nil {
		return nil, fmt.Errorf("parsing server config JSON: %v", err)
	}
	if err := ValidateServerConfig(&cfg); err != nil {
		return nil, fmt.Errorf("validating server config: %v", err)
	}
	return &cfg, nil
}

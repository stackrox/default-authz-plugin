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

package config

import (
	"errors"
	"fmt"
	"strings"
)

// ValidateServerConfig checks the given config for consistency, and returns an error if it is invalid.
func ValidateServerConfig(cfg *ServerConfig) error {
	if cfg.DisableTLS != (cfg.TLS == nil) {
		return errors.New("either set `disableTLS` to `true`, or add a `tls` configuration section")
	}
	if cfg.TLS != nil {
		if cfg.TLS.CertFile == "" {
			return errors.New("TLS config is missing a `certFile`")
		}
		if cfg.TLS.KeyFile == "" {
			return errors.New("TLS config is missing a `keyFile`")
		}
	}

	var authMechanisms []string
	if cfg.Auth.AllowAnonymous {
		authMechanisms = append(authMechanisms, "allowAnonymous")
	}
	if cfg.Auth.HtpasswdFile != "" {
		authMechanisms = append(authMechanisms, "htpasswdFile")
	}
	if cfg.Auth.ClientCACertFile != "" {
		if cfg.TLS == nil {
			return errors.New("client certificate authentication only works with enabled TLS")
		}
		authMechanisms = append(authMechanisms, "clientCACertFile")
	}

	if len(authMechanisms) != 1 {
		return fmt.Errorf("exactly one authentication mechanism must be specified, got: [%s]", strings.Join(authMechanisms, ", "))
	}

	if cfg.Port < 0 || cfg.Port > 65535 {
		return fmt.Errorf("port %d is not a valid TCP port number", cfg.Port)
	}

	return nil
}

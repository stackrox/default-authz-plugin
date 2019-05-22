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

package config

import (
	"errors"
	"fmt"
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

	if cfg.AllowAnonymous != (cfg.Auth == nil) {
		return errors.New("either set `allowAnonymous` to `true`, or add an `auth` configuration section")
	}
	if cfg.Auth != nil {
		if (cfg.Auth.HtpasswdFile == "") == (cfg.Auth.ClientCACertFile == "") {
			return errors.New("auth configuration must specify either a `htpasswdFile`, or a `clientCACertFile`")
		}
		if cfg.Auth.ClientCACertFile != "" && cfg.TLS == nil {
			return errors.New("client certificate authentication only works with enabled TLS")
		}
	}

	if cfg.Port < 0 || cfg.Port > 65535 {
		return fmt.Errorf("port %d is not a valid TCP port number", cfg.Port)
	}

	return nil
}

package config

// TLSConfig is the (optional) TLS configuration for servers serving HTTPS.
type TLSConfig struct {
	// KeyFile is the file containing the PEM-encoded server private key.
	KeyFile string
	// CertFile is the file containing the PEM-encoded server certificate (chain).
	CertFile string
}

// AuthConfig controls authentication behavior of the server. Exactly one of the fields must have a non-zero value.
type AuthConfig struct {
	// AllowAnonymous allows anonymous access if set to true.
	AllowAnonymous bool
	// ClientCACertFile makes the server require client certificates signed by the CA certificate stored as PEM in the
	// given file.
	ClientCACertFile string
	// HtpasswdFile allows client to authenticate via basic auth, checking their passwords against the given .htpasswd
	// file (bcrypt only).
	HtpasswdFile string
}

// ServerConfig is the main HTTP(S) server configuration data structure.
type ServerConfig struct {
	// The address to bind to. The empty value (default) binds to all local interfaces.
	BindAddress string
	// The port to bind to. If set to 0 (default), binds to port 80 if TLS is disabled, or port 443 if TLS is
	// enabled.
	Port int

	// The TLS configuration. If this is not set, `DisableTLS` must be true.
	TLS *TLSConfig `json:"tls"`
	// Whether to disable TLS. For security reasons, this needs to be set explicitly if TLS is not enabled.
	DisableTLS bool

	// The authentication configuration.
	Auth AuthConfig
}

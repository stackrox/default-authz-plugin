package config

type TLSConfig struct {
	KeyFile  string
	CertFile string
}

type AuthConfig struct {
	ClientCACertFile string
	HtpasswdFile     string
}

type ServerConfig struct {
	BindAddress string
	Port        int

	TLS        *TLSConfig `json:"tls"`
	DisableTLS bool

	Auth           *AuthConfig
	AllowAnonymous bool
}

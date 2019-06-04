package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/stackrox/default-authz-plugin/server/config"
)

func createTLSConfig(serverTLSConf *config.TLSConfig, authConf config.AuthConfig) (*tls.Config, error) {
	if serverTLSConf == nil {
		return nil, nil
	}

	cert, err := tls.LoadX509KeyPair(serverTLSConf.CertFile, serverTLSConf.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("loading server key pair: %v", err)
	}

	tlsConf := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		Certificates:             []tls.Certificate{cert},
		ClientAuth:               tls.NoClientCert,
	}

	if authConf.ClientCACertFile != "" {
		clientCACertsPEM, err := ioutil.ReadFile(authConf.ClientCACertFile)
		if err != nil {
			return nil, fmt.Errorf("reading client CA certs: %v", err)
		}

		clientCAPool := x509.NewCertPool()
		if ok := clientCAPool.AppendCertsFromPEM(clientCACertsPEM); !ok {
			return nil, fmt.Errorf("no client CA certificates found in %s", authConf.ClientCACertFile)
		}

		tlsConf.ClientCAs = clientCAPool
		tlsConf.ClientAuth = tls.RequireAndVerifyClientCert
	}

	return tlsConf, nil
}

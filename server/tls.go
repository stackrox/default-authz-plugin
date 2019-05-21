package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/stackrox/sample-authz-plugin/server/config"
	"io/ioutil"
)

func createTLSConfig(serverTlsConf *config.TLSConfig, authConf *config.AuthConfig) (*tls.Config, error) {
	if serverTlsConf == nil {
		return nil, nil
	}

	cert, err := tls.LoadX509KeyPair(serverTlsConf.CertFile, serverTlsConf.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("loading server key pair: %v", err)
	}

	tlsConf := &tls.Config{
		MinVersion: tls.VersionTLS12,
		PreferServerCipherSuites: true,
		Certificates: []tls.Certificate{cert},
		ClientAuth: tls.NoClientCert,
	}

	if authConf != nil && authConf.ClientCACertFile != "" {
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

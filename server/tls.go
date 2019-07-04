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

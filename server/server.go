package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stackrox/default-authz-plugin/server/config"
)

// ServeFunc abstracts the `ListenAndServe()` or `ListenAndServeTLS("", "")` mechanism of an `http.Server`.
type ServeFunc func() error

// Create creates an HTTP server from the given config.
func Create(config *config.ServerConfig, handler http.Handler) (ServeFunc, error) {
	tlsConf, err := createTLSConfig(config.TLS, config.Auth)
	if err != nil {
		return nil, err
	}

	port := config.Port
	if port == 0 {
		if tlsConf != nil {
			port = 443
		} else {
			port = 80
		}
	}

	effectiveHandler := handler

	if config.Auth.HtpasswdFile != "" {
		effectiveHandler, err = createBasicAuthHandler(config.Auth.HtpasswdFile, handler)
		if err != nil {
			return nil, err
		}
	}

	server := &http.Server{
		Addr:      fmt.Sprintf("%s:%d", config.BindAddress, port),
		TLSConfig: tlsConf,
		Handler:   effectiveHandler,
		ErrorLog:  log.New(os.Stderr, "http server", log.LstdFlags),
	}

	if tlsConf != nil {
		return func() error {
			log.Println("Listening w/ TLS on", server.Addr)
			return server.ListenAndServeTLS("", "")
		}, nil
	}

	return func() error {
		log.Println("Listening in plaintext mode on", server.Addr)
		return server.ListenAndServe()
	}, nil
}

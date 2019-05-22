package main

import (
	"errors"
	"flag"
	"log"
	"net/http"

	"github.com/stackrox/sample-authz-plugin/server"
	"github.com/stackrox/sample-authz-plugin/server/config"
)

var (
	serverConfigFlag = flag.String("server-config", "", "Server configuration file (JSON format)")
)

func main() {
	if err := mainCmd(); err != nil {
		log.Fatal(err)
	}
}

func mainCmd() error {
	flag.Parse()

	if *serverConfigFlag == "" {
		return errors.New("must specify a server configuration file")
	}

	serverCfg, err := config.LoadServerConfig(*serverConfigFlag)
	if err != nil {
		return err
	}

	httpServeFunc, err := server.Create(serverCfg, http.DefaultServeMux)
	if err != nil {
		return err
	}

	return httpServeFunc()
}

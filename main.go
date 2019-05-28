package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/stackrox/sample-authz-plugin/rules"
	"github.com/stackrox/sample-authz-plugin/rules/engines"
	"github.com/stackrox/sample-authz-plugin/server"
	"github.com/stackrox/sample-authz-plugin/server/config"

	// Ensure the rule engines are registered.
	_ "github.com/stackrox/sample-authz-plugin/rules/engines/constant"
)

var (
	serverConfigFlag = flag.String("server-config", "", "Server configuration file (JSON format)")
	engineFlag       = flag.String("engine", "deny_all", "Rule engine (<name>[:<params>])")
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

	engine, err := engines.GetRuleEngine(*engineFlag)
	if err != nil {
		return fmt.Errorf("failed to create rule engine: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/authorize", rules.NewHandler(engine))

	httpServeFunc, err := server.Create(serverCfg, mux)
	if err != nil {
		return err
	}

	return httpServeFunc()
}

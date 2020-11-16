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

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/stackrox/default-authz-plugin/rules"
	"github.com/stackrox/default-authz-plugin/rules/engines"
	"github.com/stackrox/default-authz-plugin/server"
	"github.com/stackrox/default-authz-plugin/server/config"

	// Ensure the rule engines are registered.
	_ "github.com/stackrox/default-authz-plugin/rules/engines/constant"
	_ "github.com/stackrox/default-authz-plugin/rules/engines/gval"
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
	mux.Handle("/authorize", rules.NewHandler(engine, strings.EqualFold(os.Getenv("LOGLEVEL"), "debug")))

	httpServeFunc, err := server.Create(serverCfg, mux)
	if err != nil {
		return err
	}

	return httpServeFunc()
}

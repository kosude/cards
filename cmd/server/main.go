/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package main

import (
	"net"
	"net/http"
	"os"
	"strconv"

	"gitlab.com/kosude/cards/internal/config"
	"gitlab.com/kosude/cards/internal/logger"
)

func main() {
	// init logger
	log := logger.New()

	// load environment configuration...
	cfg, err := config.LoadYaml(os.Getenv("API_ENV"), log)
	if err != nil {
		log.Errorf("Failed to load application configuration: %v", err)
		os.Exit(2)
	}
	// ...and update logger verbosity from this
	log.SetVerbose(cfg.DeployType == "development")
	if cfg.DeployType == "development" {
		log.Warnf("Running in a development environment")
	}

	// build HTTP server
	addr := net.JoinHostPort("::", strconv.Itoa(cfg.ServerPort))
	server := &http.Server{
		Addr: addr,
		// Handler: buildHandler(),
	}

	// start HTTP server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Errorf("%v", err)
		os.Exit(1)
	}
}

/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package main

import (
	"net"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	echomware "github.com/labstack/echo/v4/middleware"
	"gitlab.com/kosude/cards/cmd/server/router"
	"gitlab.com/kosude/cards/internal/config"
	"gitlab.com/kosude/cards/internal/logger"
)

func main() {
	// init logger
	log := logger.New()

	apiEnv := os.Getenv("API_ENV")
	if apiEnv == "" {
		log.Errorf("Variable API_ENV is not set")
		os.Exit(2)
	}
	// load environment configuration...
	cfg, err := config.LoadYaml(os.Getenv("API_ENV"), log)
	if err != nil {
		log.Errorf("Failed to load application configuration: %v", err)
		os.Exit(3)
	}
	// ...and update logger verbosity from this
	log.SetVerbose(cfg.DeployType == "development")
	if cfg.DeployType == "development" {
		log.Warnf("Running in a development environment")
	}

	addr := net.JoinHostPort("::", strconv.Itoa(cfg.ServerPort))

	e := echo.New()
	e.HideBanner = true

	setupRoutes(e, cfg, &log)
	setupMiddleware(e, &log)

	log.Errorf("%v", e.Start(addr))
}

// Set up all route handlers on the given Echo instance
func setupRoutes(e *echo.Echo, cfg *config.Config, log *logger.Logger) {
	base := e.Group(cfg.RouteBase)
	v1 := base.Group("/v1")

	router.InitRoutes(v1, log)
}

// Set up middleware configuration on the given Echo instance
func setupMiddleware(e *echo.Echo, log *logger.Logger) {
	// logging middleware; defer to our logging system
	e.Use(echomware.RequestLoggerWithConfig(echomware.RequestLoggerConfig{
		LogMethod: true,
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v echomware.RequestLoggerValues) error {
			log.Infof("%v uri: %v, status: %v", v.Method, v.URI, v.Status)
			return nil
		},
	}))
}

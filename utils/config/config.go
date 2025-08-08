/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package config

import (
	"os"
	"strings"

	"gitlab.com/kosude/cards/utils/logger"
	"gopkg.in/yaml.v3"
)

// default config values
const default_serverPort = 8100
const default_routeBase = ""

// An API env configuration
type Config struct {
	ServerPort int    `yaml:"server_port"`
	DeployType string `yaml:"deployment_type"`
	RouteBase  string `yaml:"route_base"`
}

// Return a configuration struct which is derived from a specified configuration YAML file
func LoadYaml(file string, logger logger.Logger) (*Config, error) {
	c := Config{
		ServerPort: default_serverPort,
		RouteBase:  default_routeBase,
	}

	// read yaml file into c
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// trim ending slash from the route base, if present
	c.RouteBase = strings.TrimSuffix(c.RouteBase, "/")

	return &c, err
}

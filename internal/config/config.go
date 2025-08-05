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

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/kosude/cards/internal/logger"
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
	println(c.RouteBase)

	// validate final environment struct
	if err = c.validate(); err != nil {
		return nil, err
	}

	return &c, err
}

// Validate the config struct
func (c *Config) validate() error {
	// any required fields validated here
	return validation.ValidateStruct(c,
		// deployment_type is required
		validation.Field(
			&c.DeployType, validation.Required,
		),
		// deployment_type is within an enum
		validation.Field(
			&c.DeployType,
			validation.NewStringRule(
				func(v string) bool {
					// must be either "development" or "production".
					if v != "development" && v != "production" {
						return false
					}
					return true
				}, "deployment_type must be one of ['development', 'production']")),
	)
}

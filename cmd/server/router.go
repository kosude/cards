/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package main

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/slash"
)

func buildRouter(routeBase string) http.Handler {
	router := routing.New()

	// common handlers and middleware
	router.Use(
		slash.Remover(http.StatusMovedPermanently),
	)

	base := router.Group(routeBase)
	api := base.Group("/v1") // TODO consider later revisions to 'v1'

	api.Get("", func(c *routing.Context) error {
		return c.Write("API root")
	})

	return router
}

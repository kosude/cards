/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package router

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/kosude/cards/controllers/card"
	"gitlab.com/kosude/cards/internal/logger"
)

// Set all generation routes under the /card/ group
func SetCardRoutes(base *echo.Group, log *logger.Logger) {
	g := base.Group("/card")

	g.GET("/languages", func(c echo.Context) error {
		return card.Languages(c, log)
	})
}

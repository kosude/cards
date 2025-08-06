/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package router

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/kosude/cards/controllers/dyn"
)

// Set all routes under the /dyn/ group
func SetDynRoutes(base *echo.Group) {
	g := base.Group("/dyn")

	g.GET("/languages", dyn.Languages)
}

/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package router

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/kosude/cards/internal/logger"
)

// Initialise all routes for the base Echo group
func InitRoutes(grp *echo.Group, log *logger.Logger) {
	SetDynRoutes(grp, log)
}

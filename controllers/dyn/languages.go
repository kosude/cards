/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package dyn

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Generate dynamic programming language overview card
func Languages(c echo.Context) error {
	return c.String(http.StatusOK, "Languages")
}

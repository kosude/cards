/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package card

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/kosude/cards/internal/logger"
	"gitlab.com/kosude/cards/internal/render"
	"gitlab.com/kosude/cards/internal/render/card"
	"gitlab.com/kosude/cards/internal/render/component"
	"gitlab.com/kosude/cards/internal/render/style"
	"gitlab.com/kosude/cards/internal/render/style/colourthemes"
)

// Generate dynamic programming language overview card
func Languages(c echo.Context, log *logger.Logger) error {
	// init card service
	if err := render.ParseTemplates(log); err != nil {
		return err
	}

	// configure card
	cr := card.New(colourthemes.Nihon(), style.Layout{
		CardWidth:    350,
		CardHeight:   270,
		BorderRadius: 4.5,
		StrokeWidth:  1,
	})

	cr.AddComponent(&component.Label{})

	svg, err := cr.RenderSVG()
	if err != nil {
		return err
	}

	c.Response().Header().Set("Content-Type", "image/svg+xml")

	return c.String(http.StatusOK, svg)
}

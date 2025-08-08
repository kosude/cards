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
	"gitlab.com/kosude/cards/render"
	"gitlab.com/kosude/cards/render/card"
	"gitlab.com/kosude/cards/render/component"
	"gitlab.com/kosude/cards/themes/colours"
	"gitlab.com/kosude/cards/themes/fonts"
	"gitlab.com/kosude/cards/themes/layouts"
	"gitlab.com/kosude/cards/utils/logger"
)

// Generate dynamic programming language overview card
func Languages(c echo.Context, log *logger.Logger) error {
	// init card service
	if err := render.ParseTemplates(log); err != nil {
		return err
	}

	// configure card
	cr := card.New(colours.Nihon(), layouts.Default(), fonts.DefaultSans())

	cr.AddComponent(component.NewLabel("Most Used Languages", component.ClassLabelHeader))

	svg, err := cr.RenderSVG()
	if err != nil {
		return err
	}

	c.Response().Header().Set("Content-Type", "image/svg+xml")

	return c.String(http.StatusOK, svg)
}

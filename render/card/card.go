/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package card

import (
	"gitlab.com/kosude/cards/render"
	"gitlab.com/kosude/cards/render/component"
	"gitlab.com/kosude/cards/render/layout"
	"gitlab.com/kosude/cards/render/style"
)

// Struct containing per-card info and components for rendering
type Card struct {
	// Slice of topmost components in the card
	components []component.IComponent

	colours   style.Colours   // Colour configuration
	layoutCfg style.LayoutCfg // Layout configuration
	fonts     style.Fonts     // Font configuration

	vertical *layout.Vertical // Topmost vertical layout manager
}

// Instantiate a new card instance
func New(col style.Colours, lytCfg style.LayoutCfg, fnt style.Fonts) *Card {
	return &Card{
		colours:   col,
		layoutCfg: lytCfg,
		fonts:     fnt,

		vertical: layout.NewVertical(lytCfg),
	}
}

// Add a component to the card's topmost component list
func (c *Card) AddComponent(comp component.IComponent) {
	c.vertical.PushComponent(comp) // update layout, and position the component correctly
	c.components = append(c.components, comp)
}

// Data to be passed to the card.xml template when rendering
type cardRenderData struct {
	render.Data

	Partials   []string // Rendered toplevel component SVg strings
	Stylesheet string   // Stylesheet partial
}

// Render a card to an SVG string
func (c *Card) RenderSVG() (string, error) {
	// initial card template data
	data := cardRenderData{
		Data: render.Data{
			Colours:   c.colours,
			LayoutCfg: c.layoutCfg,
			Fonts:     c.fonts,
		},
		Partials: []string{},
	}

	// attempt to render the stylesheet
	stylesheet, err := render.Stylesheet(c.fonts)
	if err != nil {
		return "", err
	}
	data.Stylesheet = stylesheet

	// collect partial renders from each toplevel component
	for _, comp := range c.components {
		// attempt to render this component
		str, err := comp.RenderSVG(c.colours)
		if err != nil {
			return "", err
		}
		data.Partials = append(data.Partials, str)
	}

	// correct card dimensions and position to account for stroke width
	data.Width = data.LayoutCfg.CardWidth - (data.LayoutCfg.StrokeWidth)
	data.Height = data.LayoutCfg.CardHeight - (data.LayoutCfg.StrokeWidth)
	data.PosX = data.LayoutCfg.StrokeWidth / 2
	data.PosY = data.LayoutCfg.StrokeWidth / 2

	cardTpl, err := render.FromTemplate("card.xml", data)
	if err != nil {
		return "", err
	}

	return cardTpl, nil
}

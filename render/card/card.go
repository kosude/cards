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
	"gitlab.com/kosude/cards/render/style"
)

// Struct containing per-card info and components for rendering
type Card struct {
	// Slice of topmost components in the card
	components []component.IComponent

	colours style.Colours // Colour configuration
	layout  style.Layout  // Layout configuration
	fonts   style.Fonts   // Font configuration
}

// Instantiate a new card instance
func New(col style.Colours, lyt style.Layout, fnt style.Fonts) *Card {
	return &Card{
		colours: col,
		layout:  lyt,
		fonts:   fnt,
	}
}

// Add a component to the card's topmost component list
func (c *Card) AddComponent(comp component.IComponent) {
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
			Colours: c.colours,
			Layout:  c.layout,
			Fonts:   c.fonts,
		},
		Partials:   []string{},
		Stylesheet: render.Stylesheet(c.fonts),
	}

	// collect partial renders from each toplevel component
	for _, comp := range c.components {
		// attempt to render this component
		str, err := comp.RenderSVG(c.colours, c.layout)
		if err != nil {
			// TODO better error handling here
			continue
		}
		data.Partials = append(data.Partials, str)
	}

	// correct card dimensions and position to account for stroke width
	data.Width = data.Layout.CardWidth - (data.Layout.StrokeWidth)
	data.Height = data.Layout.CardHeight - (data.Layout.StrokeWidth)
	data.PosX = data.Layout.StrokeWidth / 2
	data.PosY = data.Layout.StrokeWidth / 2

	cardTpl, err := render.FromTemplate("card.xml", data)
	if err != nil {
		return "", err
	}

	return cardTpl, nil
}

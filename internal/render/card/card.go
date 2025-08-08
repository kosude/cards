/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package card

import (
	"gitlab.com/kosude/cards/internal/render"
	"gitlab.com/kosude/cards/internal/render/component"
	"gitlab.com/kosude/cards/internal/render/style"
)

// Struct containing per-card info and components for rendering
type Card struct {
	// Slice of topmost components in the card
	components []component.IComponent

	Colours style.Colours // Colour configuration
	layout  style.Layout  // Layout configuration
}

// Instantiate a new card instance
func New(col style.Colours, lyt style.Layout) *Card {
	return &Card{
		Colours: col,
		layout:  lyt,
	}
}

// Add a component to the card's topmost component list
func (c *Card) AddComponent(comp component.IComponent) {
	c.components = append(c.components, comp)
}

// Data to be passed to the card.xml template when rendering
type cardRenderData struct {
	render.Data

	Partials []string // rendered toplevel component SVg strings

	// toplevel rect dimensions corrected at runtime to account for stroke width
	InnerCardWidth  float32
	InnerCardHeight float32

	CardPosX float32
	CardPosY float32
}

// Render a card to an SVG string
func (c *Card) RenderSVG() (string, error) {
	// initial card template data
	data := cardRenderData{
		Data: render.Data{
			Colours: c.Colours,
			Layout:  c.layout,
		},
		Partials: []string{},
	}

	// collect partial renders from each toplevel component
	for _, comp := range c.components {
		// attempt to render this component
		str, err := comp.RenderSVG()
		if err != nil {
			// TODO better error handling here
			continue
		}
		data.Partials = append(data.Partials, str)
	}

	// correct card dimensions and position to account for stroke width
	data.InnerCardWidth = data.Layout.CardWidth - (data.Layout.StrokeWidth)
	data.InnerCardHeight = data.Layout.CardHeight - (data.Layout.StrokeWidth)
	data.CardPosX = data.Layout.StrokeWidth / 2
	data.CardPosY = data.Layout.StrokeWidth / 2

	cardTpl, err := render.FromTemplate("card.xml", data)
	if err != nil {
		return "", err
	}

	return cardTpl, nil
}

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
)

// Struct containing per-card info and components for rendering
type Card struct {
	// Slice of topmost components in the card
	components []component.IComponent
}

// Instantiate a new card instance
func New() *Card {
	return &Card{}
}

// Add a component to the card's topmost component list
func (c *Card) AddComponent(comp component.IComponent) {
	c.components = append(c.components, comp)
}

// Render a card to an SVG string
func (c *Card) RenderSVG() (string, error) {
	// collect partial renders from each toplevel component
	var partials []string
	for _, comp := range c.components {
		// attempt to render this component
		str, err := comp.RenderSVG()
		if err != nil {
			// TODO better error handling here
			continue
		}
		partials = append(partials, str)
	}

	cardTplData := struct {
		Partials []string
	}{
		Partials: partials,
	}

	cardTpl, err := render.FromTemplate("card.xml", cardTplData)
	if err != nil {
		return "", err
	}

	println(cardTpl)

	return cardTpl, nil
}

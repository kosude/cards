/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package layout

import (
	"gitlab.com/kosude/cards/render/component"
	"gitlab.com/kosude/cards/render/style"
)

// Struct containing vertical layout positioning data
type Vertical struct {
	// Baseline X position
	xBaseline float32
	// Baseline Y position
	yBaseline float32

	// Y-position of next component in vertical layout
	yCursor float32
}

// Instantiate a new vertical layout with the given layout configuration
func NewVertical(cfg style.LayoutCfg) *Vertical {
	return &Vertical{
		xBaseline: cfg.CardPaddingX,
		yBaseline: cfg.CardPaddingY,

		yCursor: cfg.CardPaddingY,
	}
}

// Push a component to the vertical layout
func (v *Vertical) PushComponent(comp component.IComponent) {
	comp.SetPosition(v.xBaseline, v.yCursor)
	v.yCursor += comp.Height()
}

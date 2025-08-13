/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package component

import (
	"gitlab.com/kosude/cards/render/style"
)

// General component interface
type IComponent interface {
	// Render the component as an SVG element(s)
	RenderSVG(colours style.Colours) (string, error)

	Width() float32  // Get the component's width
	Height() float32 // Get the component's height

	// Set the component's position
	SetPosition(x float32, y float32)
	PosX() float32 // Get the component's X position
	PosY() float32 // Get the component's Y position
}

// Interface for a component that can have children elements
type IParentComponent interface {
	IComponent

	// Add child component(s) to the component
	AppendChildren(...IComponent)

	// Return a slice of nested components
	GetChildren() []IComponent
}

/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package component

import "gitlab.com/kosude/cards/render/style"

// General component interface
type IComponent interface {
	// Render the component as an SVG element(s)
	RenderSVG(colours style.Colours, layout style.Layout) (string, error)
}

// Interface for a component that can have children elements
type IParentComponent interface {
	IComponent

	// Add child component(s) to the component
	AppendChildren(...IComponent)

	// Return a slice of nested components
	GetChildren() []IComponent
}

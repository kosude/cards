/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package layout

import "gitlab.com/kosude/cards/render/component"

// Generic layout manager interface
type ILayout interface {
	// Push a component to the layout, updating both the layout and the component
	PushComponent(comp *component.IComponent)
}

/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package component

import "gitlab.com/kosude/cards/internal/render"

// Basic text label component
type Label struct {
}

// Render the component as an SVG element(s)
func (c *Label) RenderSVG() (string, error) {
	return render.FromTemplate("label.xml", nil)
}

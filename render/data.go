/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package render

import "gitlab.com/kosude/cards/render/style"

// Common data to be passed to all rendered card/component templates
type Data struct {
	Colours style.Colours // Colour configuration
	Layout  style.Layout  // Layout configuration
	Fonts   style.Fonts   // Font configuration

	Width  float32 // Width
	Height float32 // Height
	PosX   float32 // X position
	PosY   float32 // Y position
}

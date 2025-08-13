/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package style

// A struct for card layout configuration
type LayoutCfg struct {
	CardWidth    float32 // Width of the card in pixels
	CardHeight   float32 // Height of the card in pixels
	BorderRadius float32 // Border corner radius
	StrokeWidth  float32 // Border thickness

	CardPaddingX float32 // Topmost horizontal padding
	CardPaddingY float32 // Topmost vertical padding
}

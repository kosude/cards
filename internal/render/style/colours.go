/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package style

import "fmt"

// A struct containing fields for standard colour values in rendered card SVGs
type Colours struct {
	CardFill   RGBA // Card background colour
	CardStroke RGBA // Card border stroke colour
}

// A struct representing an RGBA colour value
type RGBA struct {
	R, G, B, A uint8
}

// Construct a new RGBA struct with 4 channel values
func Rgba(r uint8, g uint8, b uint8, a uint8) RGBA {
	return RGBA{R: r, G: g, B: b, A: a}
}

// Construct a new RGBA struct with 3 channel values (fully opaque)
func Rgb(r uint8, g uint8, b uint8) RGBA {
	return Rgba(r, g, b, 255)
}

// Return the colour as a 4-channel hex string
func (c RGBA) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
}

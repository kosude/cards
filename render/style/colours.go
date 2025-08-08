/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package style

import "gitlab.com/kosude/cards/utils/colour"

// A struct containing fields for standard colour values in rendered card SVGs
type Colours struct {
	CardFill   colour.RGBA // Card background colour
	CardStroke colour.RGBA // Card border stroke colour
}

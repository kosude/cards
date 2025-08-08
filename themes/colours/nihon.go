/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package colours

import (
	"gitlab.com/kosude/cards/render/style"
	"gitlab.com/kosude/cards/utils/colour"
)

// Return the 'Nihon' (red and white) colour theme
func Nihon() style.Colours {
	return style.Colours{
		CardFill:   colour.Rgb(254, 254, 254),
		CardStroke: colour.Rgb(210, 210, 210),
	}
}

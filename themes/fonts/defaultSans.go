/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package fonts

import (
	"gitlab.com/kosude/cards/render/style"
	"gitlab.com/kosude/cards/utils/font"
)

// Return the default sans font configuration
func DefaultSans() style.Fonts {
	return style.Fonts{
		Fallback: font.New("sans-serif"),
		Header:   font.New("Segoe UI", "Ubuntu"),
	}
}

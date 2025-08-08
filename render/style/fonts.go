/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package style

import "gitlab.com/kosude/cards/utils/font"

// A struct of standard fonts used in rendered card SVGs
type Fonts struct {
	Fallback font.Font // Standard font to be used in place of others that are unavailable (e.g. 'sans-serif' or 'Arial')
	Header   font.Font // Font used in header labels
}

/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package font

import (
	"fmt"
	"strings"
)

// Font type as expected in CSS
type Font struct {
	fonts []string // List of font options
}

// Create a new font list with the given options
func New(fonts ...string) Font {
	// quote each string
	quoted := make([]string, len(fonts))
	for i := range fonts {
		quoted[i] = fmt.Sprintf("\"%s\"", fonts[i])
	}

	return Font{
		fonts: quoted,
	}
}

// Format the font as a string containing comma-separated string font names
func (f Font) CSV() string {
	return strings.Join(f.fonts, ", ")
}

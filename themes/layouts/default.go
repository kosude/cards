/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package layouts

import "gitlab.com/kosude/cards/render/style"

// Return the default layout configuration
func Default() style.Layout {
	return style.Layout{
		CardWidth:    350,
		CardHeight:   270,
		BorderRadius: 4.5,
		StrokeWidth:  1,
	}
}

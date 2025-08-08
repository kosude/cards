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
}

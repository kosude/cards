/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package component

import (
	"gitlab.com/kosude/cards/render"
	"gitlab.com/kosude/cards/render/style"
)

// Basic text label component
type Label struct {
	// label text
	text string
}

// Create a new label component
func NewLabel(text string) *Label {
	return &Label{
		text: text,
	}
}

// Data to be passed to the label.xml template when rendering
type labelRenderData struct {
	render.Data

	Text string
}

// Render the component as an SVG partial
func (l *Label) RenderSVG(colours style.Colours, layout style.Layout) (string, error) {
	// label template data
	data := labelRenderData{
		Data: render.Data{
			Colours: colours,
			Layout:  layout,
		},
		Text: l.text,
	}

	// TODO dynamic positioning based on the other components in the card (do in AddComponent)
	data.PosX = data.Layout.CardPaddingX
	data.PosY = data.Layout.CardPaddingY

	labelTpl, err := render.FromTemplate("label.xml", data)
	if err != nil {
		return "", err
	}

	return labelTpl, nil
}

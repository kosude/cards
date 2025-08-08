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

// Enum type for label classes (informs styling)
type LabelClass string

const (
	ClassLabelHeader LabelClass = "label-header" // Header/title labels
)

// Basic text label component
type Label struct {
	// Label text
	text string

	// Class of the label
	class LabelClass
}

// Create a new label component
func NewLabel(text string, class LabelClass) *Label {
	return &Label{
		text:  text,
		class: class,
	}
}

// Data to be passed to the label.xml template when rendering
type labelRenderData struct {
	render.Data

	Text  string
	Class LabelClass
}

// Render the component as an SVG partial
func (l *Label) RenderSVG(colours style.Colours, layout style.Layout) (string, error) {
	// label template data
	data := labelRenderData{
		Data: render.Data{
			Colours: colours,
			Layout:  layout,
		},
		Text:  l.text,
		Class: l.class,
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

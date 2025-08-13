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

	posX float32 // X position
	posY float32 // Y position
}

// Create a new label component
func NewLabel(text string, class LabelClass) *Label {
	return &Label{
		text:  text,
		class: class,

		posX: 0,
		posY: 0,
	}
}

// Data to be passed to the label.xml template when rendering
type labelRenderData struct {
	render.Data

	Text  string
	Class LabelClass
}

// Render the label as an SVG partial
func (l *Label) RenderSVG(colours style.Colours) (string, error) {
	// label template data
	data := labelRenderData{
		Data: render.Data{
			Colours: colours,
			PosX:    l.posX,
			PosY:    l.posY,
		},
		Text:  l.text,
		Class: l.class,
	}

	labelTpl, err := render.FromTemplate("label.xml", data)
	if err != nil {
		return "", err
	}

	return labelTpl, nil
}

// Get the label's width
func (l *Label) Width() float32 {
	// TODO calculate average font character width * amount of characters
	return 0
}

// Get the label's height
func (l *Label) Height() float32 {
	// TODO calculate average font character height * amount of lines
	return 24
}

// Set the label's position
func (l *Label) SetPosition(x float32, y float32) {
	l.posX = x
	l.posY = y
}

// Get the label's X position
func (l *Label) PosX() float32 {
	return l.posX
}

// Get the label's Y position
func (l *Label) PosY() float32 {
	return l.posY
}

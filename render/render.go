/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package render

import (
	"bytes"
	"embed"
	"fmt"
	"runtime/debug"
	"text/template"

	"gitlab.com/kosude/cards/render/style"
	"gitlab.com/kosude/cards/utils/logger"
)

//go:embed */*.xml style/stylesheet.css
var tplFS embed.FS
var tpls *template.Template

// Instantiate component template data
func ParseTemplates(log *logger.Logger) (err error) {
	// handle nil value panic from template parsing
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("when parsing template (ParseTemplates): %v", r)
			fmt.Printf("Stack trace for 500 error: (TEMPLATE PARSE):\n\n%s", debug.Stack())
		}
	}()

	tpls, err = template.
		New("root").
		ParseFS(tplFS, "*/*.xml", "style/stylesheet.css")

	// print each template name (equal to filename)
	for _, t := range tpls.Templates() {
		log.Debugf("Loaded component template %s", t.Name())
	}

	return err
}

// Render a specified template by its name
func FromTemplate(tpl string, data any) (_ string, err error) {
	// handle nil value panic from template parsing
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("when executing template (FromTemplate): %v", r)
			fmt.Printf("Stack trace for 500 error: (TEMPLATE EXECUTE):\n\n%s", debug.Stack())
		}
	}()

	buf := new(bytes.Buffer)
	if err := tpls.ExecuteTemplate(buf, tpl, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Render a CSS stylesheet using data from the given configuration struct(s)
func Stylesheet(f style.Fonts) (string, error) {
	return FromTemplate("stylesheet.css",
		// data struct
		struct {
			Fonts style.Fonts
		}{
			f,
		})
}

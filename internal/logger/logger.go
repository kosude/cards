/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger instance struct
type Logger struct {
	lr *logrus.Logger
}

// Instantiate a new logger instance
func New() Logger {
	lr := logrus.New()

	// setup formatter
	lr.SetFormatter(&logrus.TextFormatter{})

	// output to stdout instead of the default stderr
	lr.SetOutput(os.Stdout)

	// by default, output info warnings and above
	lr.SetLevel(logrus.WarnLevel)

	return Logger{
		lr: lr,
	}
}

// Set the verbosity state of the logger to the value of `verbose`
func (l *Logger) SetVerbose(verbose bool) {
	if verbose {
		l.lr.SetLevel(logrus.DebugLevel)
	} else {
		l.lr.SetLevel(logrus.WarnLevel)
	}
}

// Write an error message
func (l *Logger) Errorf(fmt string, args ...any) {
	l.lr.Errorf(fmt, args...)
}

// Write a warning message
func (l *Logger) Warnf(fmt string, args ...any) {
	l.lr.Warnf(fmt, args...)
}

// Write an info message
func (l *Logger) Infof(fmt string, args ...any) {
	l.lr.Infof(fmt, args...)
}

// Write a debug message
func (l *Logger) Debugf(fmt string, args ...any) {
	l.lr.Debugf(fmt, args...)
}

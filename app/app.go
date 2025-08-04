/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Struct for routing etc.
type App struct {
	router chi.Router
}

// Configure and return a new App instance
func NewApp() App {
	r := chi.NewRouter()

	return App{
		router: r,
	}
}

// Run the app
func (a App) Run() {
	http.ListenAndServe(":3000", a.router)
}

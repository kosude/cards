/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "testing!!")
	})
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "root (working???)")
	})

	server := &http.Server{
		Addr:    net.JoinHostPort("::", "8100"),
		Handler: mux,
	}

	println("Newest version")

	log.Fatalln(server.ListenAndServe())
}

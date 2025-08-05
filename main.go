/*
 * Copyright (c) 2025 Jack Bennett.
 * All Rights Reserved.
 *
 * See the LICENCE file for more information.
 */

package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	addr := net.JoinHostPort("::", "8100")
	server := &http.Server{Addr: addr}
	log.Fatalln(server.ListenAndServe())
}

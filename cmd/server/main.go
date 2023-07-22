package main

import (
	"log"

	"github.com/Ryzhon/proglog2/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}

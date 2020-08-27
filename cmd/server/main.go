package main

import (
	"log"

	"github.com/mayankkapoor/commit-log-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}

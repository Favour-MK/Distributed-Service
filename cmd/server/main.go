package main

import (
	"log"

	"github.com/Favour-MK/Distributed-Service/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}

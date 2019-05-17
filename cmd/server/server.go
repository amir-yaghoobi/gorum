package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorum/pkg/api"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	log.Fatal(http.ListenAndServe(addr, api.Build()))
}

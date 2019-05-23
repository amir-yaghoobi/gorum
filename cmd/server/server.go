package main

import (
	"fmt"
	"gorum/pkg/auth"
	"gorum/pkg/store/mem"
	"net/http"
	"os"

	"github.com/google/logger"

	"gorum/pkg/api"
)

func main() {
	log := logger.Init("main", true, false, os.Stdout)

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	userService := auth.UserService{
		Storer: &mem.UserStore{},
	}

	sessionService := auth.SessionService{
		Storer: &mem.SessionStore{},
	}

	router, err := api.Build(userService, sessionService)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Infof("starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

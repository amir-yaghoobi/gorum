package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-pg/pg"
	"github.com/google/logger"

	"gorum/pkg/api"
	"gorum/pkg/auth"
	"gorum/pkg/store/postgres"
)

func main() {
	log := logger.Init("main", true, false, os.Stdout)

	db := initDatabase()

	userService := auth.UserService{
		Storer: postgres.NewUserStorer(db),
	}

	sessionService := auth.SessionService{
		Storer: postgres.NewSessionStorer(db),
	}

	router, err := api.Build(userService, sessionService)
	if err != nil {
		log.Fatal(err)
	}

	addr := hostAddress()
	log.Infof("starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func hostAddress() string {
	host := envWithDefault("HOST", "localhost")
	port := envWithDefault("PORT", "8080")

	return fmt.Sprintf("%s:%s", host, port)
}

func initDatabase() *pg.DB {
	host := envWithDefault("DB_HOST", "localhost")
	port := envWithDefault("DB_PORT", "5432")
	name := envWithDefault("DB_NAME", "gorum")
	user := envWithDefault("DB_USER", "postgres")
	pass := envWithDefault("DB_PASS", "")

	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		User:     user,
		Password: pass,
		Database: name,
	})
}

func envWithDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

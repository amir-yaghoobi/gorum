package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/google/logger"

	// Register migrations
	_ "gorum/pkg/store/postgres/migration"
)

const usage = `Supported commands are:
- init          creates version info table
- version       prints current database version
- up            runs all available migrations
- up [target]   runs migrations upto the target
- down          reverts last migration
- reset         revers all migrations
`

func main() {
	log := logger.Init("main", true, false, os.Stdout)

	flag.Usage = func() {
		_, _ = fmt.Fprint(flag.CommandLine.Output(), usage)
		os.Exit(2)
	}

	flag.Parse()

	db := initDatabase()

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Infof("migrated from %d to %d", oldVersion, newVersion)
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

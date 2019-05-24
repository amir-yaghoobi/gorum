package migration

import "github.com/go-pg/migrations"

func init() {
	// Disable SQL auto-discover to prevent migrations to lookup source files,
	// which are not available at runtime, for .sql files.
	migrations.DefaultCollection.DisableSQLAutodiscover(true)
}

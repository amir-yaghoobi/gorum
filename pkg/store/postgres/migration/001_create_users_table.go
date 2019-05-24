package migration

import "github.com/go-pg/migrations"

func init() {
	migrations.MustRegister(func(db migrations.DB) error {
		_, err := db.Exec(`
			CREATE TABLE users (
				id            SERIAL       PRIMARY KEY,
				name          VARCHAR(255) NOT NULL UNIQUE,
				email         VARCHAR(255) NOT NULL UNIQUE,
				secret        VARCHAR(255) NOT NULL,
				full_name     VARCHAR(255) NOT NULL,
				tag_line      VARCHAR(255),
				role          INTEGER,
				respect       INTEGER      NOT NULL DEFAULT '0',
				registered_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
			)
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}

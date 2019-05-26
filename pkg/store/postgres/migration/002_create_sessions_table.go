package migration

import "github.com/go-pg/migrations"

func init() {
	migrations.MustRegister(func(db migrations.DB) error {
		_, err := db.Exec(`
			CREATE TABLE sessions (
				id         BIGSERIAL    PRIMARY KEY,
				user_id    BIGINT       NOT NULL REFERENCES users(id),
				token      VARCHAR(255) NOT NULL UNIQUE,
				started_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
				expires_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
			)
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`DROP TABLE sessions`)
		return err
	})
}

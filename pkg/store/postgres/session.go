package postgres

import (
	"time"

	"github.com/go-pg/pg"

	"gorum/pkg/auth"
)

type session struct {
	ID        int64
	UserID    int64
	Token     string
	StartedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}

func (m *session) From(s *auth.Session) {
	m.ID = s.ID
	m.UserID = s.UserID
	m.Token = s.Token
	m.StartedAt = s.StartedAt
	m.UpdatedAt = s.UpdatedAt
	m.ExpiresAt = s.ExpiresAt
}

func (m *session) To(s *auth.Session) {
	s.ID = m.ID
	s.UserID = m.UserID
	s.Token = m.Token
	s.StartedAt = m.StartedAt
	s.UpdatedAt = m.UpdatedAt
	s.ExpiresAt = m.ExpiresAt
}

// NewSessionStorer creates a new Postgres auth.SessionStorer.
func NewSessionStorer(db *pg.DB) auth.SessionStorer {
	return &sessionStore{db: db}
}

type sessionStore struct {
	db *pg.DB
}

func (ss *sessionStore) Find(s *auth.Session, token string) error {
	model := session{}

	err := ss.db.Model(&model).Where("token = ?", token).First()
	if err != nil {
		if err == pg.ErrNoRows {
			return auth.ErrNoSession
		}

		return err
	}

	model.To(s)

	return nil
}

func (ss *sessionStore) Persist(s *auth.Session) error {
	model := session{}
	model.From(s)

	if model.ID == 0 {
		err := ss.db.Insert(&model)
		if err != nil {
			return err
		}

		s.ID = model.ID

		return nil
	}

	model.UpdatedAt = time.Now()

	err := ss.db.Update(&model)
	if err != nil {
		return err
	}

	s.UpdatedAt = model.UpdatedAt

	return nil
}

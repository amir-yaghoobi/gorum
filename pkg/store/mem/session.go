package mem

import (
	"fmt"
	"time"

	"gorum/pkg/auth"
)

// SessionStore implements an in-memory auth.SessionStorer.
type SessionStore struct {
	Sessions []auth.Session
	LastID   int64
}

// Find implements auth.SessionStorer.Find.
func (ss *SessionStore) Find(s *auth.Session, token string) error {
	found := false
	for idx := range ss.Sessions {
		if ss.Sessions[idx].Token == token {
			*s = ss.Sessions[idx]

			found = true
			break
		}
	}

	if !found {
		return auth.ErrNoSession
	}

	return nil
}

// Persist implements auth.SessionStorer.Persist.
func (ss *SessionStore) Persist(s *auth.Session) error {
	if s.ID == 0 {
		s.ID = ss.LastID + 1

		ss.Sessions = append(ss.Sessions, *s)
		ss.LastID++

		return nil
	}

	updated := false
	for idx := range ss.Sessions {
		if ss.Sessions[idx].ID == s.ID {
			s.UpdatedAt = time.Now()
			ss.Sessions[idx] = *s

			updated = true

			break
		}
	}

	if !updated {
		return fmt.Errorf("session %d does not exist", s.ID)
	}

	return nil
}

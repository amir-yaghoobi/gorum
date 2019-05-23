package auth

import (
	"errors"
	"time"

	"github.com/aslrousta/rand"
)

var (
	// ErrNoSession reports that no session could be found.
	ErrNoSession = errors.New("no session")
)

const (
	// SessionLifeTime is the maximum lifetime of a session before expired.
	SessionLifeTime = 30 * 24 * time.Hour

	tokenLength = 32
)

// Session is an open session for a user.
type Session struct {
	ID        int64
	UserID    int64
	Token     string
	StartedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}

// IsExpired retrieves whether session is expired or not.
func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}

// SessionStorer defines a storage for sessions.
type SessionStorer interface {
	// Find retrieves a session by its token.
	Find(s *Session, token string) error

	// Persist stores/updates a session.
	Persist(s *Session) error
}

// SessionService exposes operations can be performed on sessions.
type SessionService struct {
	Storer SessionStorer
}

// Start initiates a new session.
func (ss SessionService) Start(s *Session, userID int64) error {
	token, err := rand.RandomString(tokenLength, rand.All)
	if err != nil {
		return err
	}

	now := time.Now()

	s.UserID = userID
	s.Token = token
	s.StartedAt = now
	s.UpdatedAt = now
	s.ExpiresAt = now.Add(SessionLifeTime)

	return ss.Storer.Persist(s)
}

// Authenticate finds a session which corresponds to a token. It returns nil on
// success or ErrBadCredentials if token is invalid or expired.
func (ss SessionService) Authenticate(s *Session, token string) error {
	err := ss.Storer.Find(s, token)
	if err != nil {
		if err == ErrNoSession {
			return ErrBadCredentials
		}

		return err
	}

	if s.IsExpired() {
		return ErrBadCredentials
	}

	return nil
}

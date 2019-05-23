package auth_test

import (
	"testing"
	"time"

	"gorum/pkg/auth"
	"gorum/pkg/store/mem"

	. "github.com/stretchr/testify/assert"
)

func TestSession_IsExpired(t *testing.T) {
	s := auth.Session{}

	s.ExpiresAt = time.Now().Add(time.Minute)
	False(t, s.IsExpired())

	s.ExpiresAt = time.Now()
	True(t, s.IsExpired())
}

func TestSessionService_Start(t *testing.T) {
	service := auth.SessionService{
		Storer: &mem.SessionStore{},
	}

	const userID int64 = 1000

	s := auth.Session{}
	err := service.Start(&s, userID)

	NoError(t, err)
	NotZero(t, s.ID)
	Equal(t, userID, s.UserID)
	NotEmpty(t, s.Token)
	False(t, s.IsExpired())
}

func TestSessionService_Authenticate(t *testing.T) {
	service := auth.SessionService{
		Storer: &mem.SessionStore{},
	}

	const userID int64 = 1000

	s := auth.Session{}
	_ = service.Start(&s, userID)

	err := service.Authenticate(&s, "fake")
	EqualError(t, err, auth.ErrBadCredentials.Error())

	err = service.Authenticate(&s, s.Token)
	NoError(t, err)
}

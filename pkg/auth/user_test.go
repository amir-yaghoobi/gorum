package auth_test

import (
	"testing"

	"gorum/pkg/auth"
	"gorum/pkg/store/mem"

	. "github.com/stretchr/testify/assert"
)

func TestUserService_Register(t *testing.T) {
	service := auth.UserService{
		Storer: &mem.UserStore{},
	}

	u := auth.User{}
	err := service.Register(&u, "test", "fake name", "test@test.com", "12345678")

	NoError(t, err)
	NotZero(t, u.ID)
	Equal(t, auth.UserName("test"), u.Name)
	Equal(t, "fake name", u.FullName)
	Equal(t, auth.Email("test@test.com"), u.Email)
	NotEmpty(t, u.Secret)
}

func TestUserService_Authenticate(t *testing.T) {
	service := auth.UserService{
		Storer: &mem.UserStore{},
	}

	u := auth.User{}
	_ = service.Register(&u, "test", "fake name", "test@test.com", "12345678")

	err := service.Authenticate(&u, "fake", "fake")
	EqualError(t, err, auth.ErrBadCredentials.Error())

	err = service.Authenticate(&u, "test", "fake")
	EqualError(t, err, auth.ErrBadCredentials.Error())

	err = service.Authenticate(&u, "test", "12345678")
	NoError(t, err)

	err = service.Authenticate(&u, "test@test.com", "12345678")
	NoError(t, err)
}

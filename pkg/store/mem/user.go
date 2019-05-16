package mem

import (
	"fmt"
	"time"

	"gorum/pkg/auth"
)

// UserStorer implements an in-memory auth.UserStorer.
type UserStorer struct {
	Users  []auth.User
	LastID int64
}

// Find implements auth.UserStorer.Find.
func (s *UserStorer) Find(u *auth.User, id int64) error {
	return s.find(u, func(u *auth.User) bool {
		return u.ID == id
	})
}

// FindByName implements auth.UserStorer.FindByName.
func (s *UserStorer) FindByName(u *auth.User, name auth.UserName) error {
	return s.find(u, func(u *auth.User) bool {
		return u.Name == name
	})
}

// ExistsByName implements auth.UserStorer.ExistsByName.
func (s *UserStorer) ExistsByName(name auth.UserName) (bool, error) {
	exists := s.exists(func(u *auth.User) bool {
		return u.Name == name
	})
	return exists, nil
}

// FindByEmail implements auth.UserStorer.FindByEmail.
func (s *UserStorer) FindByEmail(u *auth.User, email auth.Email) error {
	return s.find(u, func(u *auth.User) bool {
		return u.Email == email
	})
}

// ExistsByEmail implements auth.UserStorer.ExistsByEmail.
func (s *UserStorer) ExistsByEmail(email auth.Email) (bool, error) {
	exists := s.exists(func(u *auth.User) bool {
		return u.Email == email
	})
	return exists, nil
}

// Persist implements auth.UserStorer.Persist.
func (s *UserStorer) Persist(u *auth.User) error {
	if u == nil {
		panic("user is nil")
	}

	if u.ID == 0 {
		u.ID = s.LastID + 1

		s.Users = append(s.Users, *u)
		s.LastID++

		return nil
	}

	updated := false
	for idx := range s.Users {
		if s.Users[idx].ID == u.ID {
			u.UpdatedAt = time.Now()
			s.Users[idx] = *u

			updated = true

			break
		}
	}

	if !updated {
		return fmt.Errorf("user %d does not exist", u.ID)
	}

	return nil
}

func (s *UserStorer) find(u *auth.User, pred func(*auth.User) bool) error {
	if u == nil {
		panic("user is nil")
	}

	found := false
	for idx := range s.Users {
		if pred(&s.Users[idx]) {
			*u = s.Users[idx]

			found = true
			break
		}
	}

	if !found {
		return auth.ErrNoUser
	}

	return nil
}

func (s *UserStorer) exists(pred func(*auth.User) bool) bool {
	for idx := range s.Users {
		if pred(&s.Users[idx]) {
			return true
		}
	}

	return false
}

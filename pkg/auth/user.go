package auth

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserName is a unique name which represents a user
type UserName string

// Email is an email address.
type Email string

// Role is role which is given to a user.
type Role int

// List of user roles.
const (
	// Member is a user who can create posts, rate them, leave comments and
	// follow other users. A member may receive special privileges based on
	// its respect points.
	Member = iota + 1

	// Admin is a user who has all the permissions of a member, in addition to
	// managing members and performing all privileged actions.
	Admin
)

// User is a registered user.
type User struct {
	ID           int64
	Name         UserName
	Email        Email
	Secret       string
	FullName     string
	TagLine      string
	Role         Role
	Respect      int
	RegisteredAt time.Time
	UpdatedAt    time.Time
}

var (
	// ErrNoUser reports that the user could not be found.
	ErrNoUser = errors.New("no user")
)

// UserStorer defines a storage for users.
type UserStorer interface {
	// Find retrieves a user by its ID.
	Find(u *User, id int64) error

	// FindByName retrieves a user by its name.
	FindByName(u *User, name UserName) error

	// ExistsByName checks if a user with a name exists.
	ExistsByName(name UserName) (bool, error)

	// FindByEmail retrieves a user by its email.
	FindByEmail(u *User, email Email) error

	// ExistsByEmail checks if a user with an email exists.
	ExistsByEmail(email Email) (bool, error)

	// Persist stores/updates a user.
	Persist(u *User) error
}

// UserService exposes operations can be performed on users.
type UserService struct {
	Storer UserStorer
}

// Register creates a new user.
func (s UserService) Register(u *User, name, fullName, email, pass string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Name = UserName(name)
	u.FullName = fullName
	u.Email = Email(email)
	u.Secret = string(hash)
	u.Role = Member
	u.RegisteredAt = time.Now()
	u.UpdatedAt = time.Now()

	return s.Storer.Persist(u)
}

// Authenticate finds a user which corresponds to a name/email and identified
// by a password. It returns nil on success or ErrBadCredentials if credentials
// are incorrect.
func (s UserService) Authenticate(u *User, principal, pass string) error {
	err := s.Storer.FindByName(u, UserName(principal))
	if err != nil {
		if err != ErrNoUser {
			return err
		}

		// TODO we can make a better guess by examining the principal to check
		//  whether it's a valid username or email, beforehand.

		// Try if user can be found by its email.
		err = s.Storer.FindByEmail(u, Email(principal))
		if err != nil {
			if err == ErrNoUser {
				return ErrBadCredentials
			}

			return err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Secret), []byte(pass))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ErrBadCredentials
		}

		return err
	}

	return nil
}

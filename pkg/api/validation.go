package api

import (
	"errors"
	"regexp"
	"unicode"
	"unicode/utf8"

	"gorum/pkg/auth"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// ErrInvalidUsername reports that username is invalid.
	ErrInvalidUsername = errors.New("invalid-username")

	// ErrDuplicateUsername reports that username is already registered.
	ErrDuplicateUsername = errors.New("duplicate-username")

	// ErrInvalidFullName reports that full name is invalid.
	ErrInvalidFullName = errors.New("invalid-full-name")

	// ErrInvalidEmail reports that email is invalid.
	ErrInvalidEmail = errors.New("invalid-email")

	// ErrDuplicateEmail reports that email is already registered.
	ErrDuplicateEmail = errors.New("duplicate-email")

	// ErrInvalidPassword reports that password is invalid.
	ErrInvalidPassword = errors.New("invalid-password")

	// ErrInvalidPasswordConfirmation reports that password and its
	// confirmation do not match.
	ErrInvalidPasswordConfirmation = errors.New("invalid-password-confirmation")
)

// ValidationError is a form validation error.
type ValidationError struct {
	Errors []error
}

// Error implements error.Error().
func (v *ValidationError) Error() string {
	return "validation error"
}

func validateUsername(name string) error {
	if utf8.RuneCountInString(name) < 6 {
		return ErrInvalidUsername
	}

	for _, r := range []rune(name) {
		if !unicode.IsLetter(r) {
			return ErrInvalidUsername
		}
	}

	exists, _ := services.User.Storer.ExistsByName(auth.UserName(name))
	if exists {
		return ErrDuplicateUsername
	}

	return nil
}

func validateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	exists, _ := services.User.Storer.ExistsByEmail(auth.Email(email))
	if exists {
		return ErrDuplicateEmail
	}

	return nil
}

func validatePassword(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return ErrInvalidPassword
	}

	return nil
}

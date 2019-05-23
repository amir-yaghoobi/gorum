package auth

import "errors"

var (
	// ErrBadCredentials reports that authentication was unsuccessful.
	ErrBadCredentials = errors.New("authentication failed")
)

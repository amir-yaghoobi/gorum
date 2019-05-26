package api

import (
	"context"
	"net/http"

	"gorum/pkg/auth"
)

type contextKey int

const (
	errorKey contextKey = iota
	userKey
)

func withError(r *http.Request, err error) *http.Request {
	ctx := context.WithValue(r.Context(), errorKey, err)
	return r.WithContext(ctx)
}

func requestError(r *http.Request) error {
	return r.Context().Value(errorKey).(error)
}

func authenticate(r *http.Request, u *auth.User) *http.Request {
	ctx := context.WithValue(r.Context(), userKey, u)
	return r.WithContext(ctx)
}

func isAuthenticated(r *http.Request) bool {
	_, ok := r.Context().Value(userKey).(*auth.User)
	return ok
}

func requestUser(r *http.Request) *auth.User {
	u, ok := r.Context().Value(userKey).(*auth.User)
	if !ok {
		return nil
	}
	return u
}

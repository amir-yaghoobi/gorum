package api

import (
	"context"
	"net/http"
)

type contextKey int

const (
	errorKey contextKey = iota
	authenticatedKey
)

func withError(r *http.Request, err error) *http.Request {
	ctx := context.WithValue(r.Context(), errorKey, err)
	return r.WithContext(ctx)
}

func requestError(r *http.Request) error {
	return r.Context().Value(authenticatedKey).(error)
}

func authenticate(r *http.Request) *http.Request {
	ctx := context.WithValue(r.Context(), authenticatedKey, true)
	return r.WithContext(ctx)
}

func isAuthenticated(r *http.Request) bool {
	_, ok := r.Context().Value(authenticatedKey).(bool)
	return ok
}

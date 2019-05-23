package api

import (
	"net/http"

	"gorum/pkg/auth"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		var session auth.Session
		err = services.Session.Authenticate(&session, sessionCookie.Value)
		if err != nil {
			if err == auth.ErrBadCredentials {
				next.ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, withError(r, err))
			return
		}

		next.ServeHTTP(w, authenticate(r))
	})
}

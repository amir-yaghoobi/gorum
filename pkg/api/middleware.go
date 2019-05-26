package api

import (
	"net/http"

	"gorum/pkg/auth"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		var s auth.Session
		err = services.Session.Authenticate(&s, cookie.Value)
		if err != nil {
			if err == auth.ErrBadCredentials {
				next.ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, withError(r, err))
			return
		}

		var u auth.User
		err = services.User.Storer.Find(&u, s.UserID)
		if err != nil {
			if err == auth.ErrNoUser {
				next.ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, withError(r, err))
			return
		}

		next.ServeHTTP(w, authenticate(r, &u))
	})
}

package api

import (
	"net/http"

	"github.com/gorilla/schema"

	"gorum/pkg/auth"
)

func buildAuthRoutes() {
	services.Router.
		Path("/register").
		Methods(http.MethodGet).
		HandlerFunc(getRegister).
		Name("registration")

	services.Router.
		Path("/register").
		Methods(http.MethodPost).
		HandlerFunc(postRegister).
		Name("register")

	services.Router.
		Path("/login").
		Methods(http.MethodGet).
		HandlerFunc(getLogin).
		Name("login")

	services.Router.
		Path("/login").
		Methods(http.MethodPost).
		HandlerFunc(postLogin).
		Name("post-login")
}

func getRegister(w http.ResponseWriter, r *http.Request) {
	if isAuthenticated(r) {
		redirect(w, r, "home")
		return
	}

	view(w, "registration", nil)
}

func postRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		viewError(w, "registration", err)
		return
	}

	body := struct {
		Username             string `schema:"username"`
		Email                string `schema:"email"`
		Password             string `schema:"password"`
		PasswordConfirmation string `schema:"password_confirmation"`
	}{}

	decoder := schema.NewDecoder()

	err = decoder.Decode(&body, r.Form)
	if err != nil {
		viewError(w, "registration", err)
		return
	}

	// TODO validate registration form

	var u auth.User
	err = services.User.Register(&u, body.Username, body.Email, body.Password)
	if err != nil {
		viewError(w, "registration", err)
		return
	}

	redirect(w, r, "login")
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	if isAuthenticated(r) {
		redirect(w, r, "home")
		return
	}

	view(w, "login", nil)
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		viewError(w, "login", err)
		return
	}

	body := struct {
		Username   string `schema:"username"`
		Password   string `schema:"password"`
		RememberMe bool   `schema:"remember_me"`
	}{}

	decoder := schema.NewDecoder()

	err = decoder.Decode(&body, r.Form)
	if err != nil {
		viewError(w, "login", err)
		return
	}

	// TODO validate login form

	var u auth.User
	err = services.User.Authenticate(&u, body.Username, body.Password)
	if err != nil {
		viewError(w, "login", err)
		return
	}

	var s auth.Session
	err = services.Session.Start(&s, u.ID)
	if err != nil {
		viewError(w, "login", err)
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    s.Token,
		Expires:  s.ExpiresAt,
	}
	http.SetCookie(w, &cookie)

	redirect(w, r, "home")
}

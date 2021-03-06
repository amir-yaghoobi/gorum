package api

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"

	"gorum/pkg/auth"
)

type registrationForm struct {
	Username             string `schema:"username"`
	FullName             string `schema:"full_name"`
	Email                string `schema:"email"`
	Password             string `schema:"password"`
	PasswordConfirmation string `schema:"password_confirmation"`
}

func (f *registrationForm) validate() error {
	ve := &ValidationError{}

	if err := validateUsername(f.Username); err != nil {
		ve.Errors = append(ve.Errors, err)
	}

	if f.FullName == "" {
		ve.Errors = append(ve.Errors, ErrInvalidFullName)
	}

	if err := validateEmail(f.Email); err != nil {
		ve.Errors = append(ve.Errors, err)
	}

	if err := validatePassword(f.Password); err != nil {
		ve.Errors = append(ve.Errors, err)
	} else if f.Password != f.PasswordConfirmation {
		ve.Errors = append(ve.Errors, ErrInvalidPasswordConfirmation)
	}

	if len(ve.Errors) > 0 {
		return ve
	}

	return nil
}

type loginForm struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
}

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

	services.Router.
		Path("/logout").
		Methods(http.MethodGet).
		HandlerFunc(getLogout).
		Name("logout")
}

func getRegister(w http.ResponseWriter, r *http.Request) {
	if isAuthenticated(r) {
		redirect(w, r, "home")
		return
	}

	view(w, r, "registration", nil)
}

func postRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		viewError(w, r, "registration", err)
		return
	}

	form := registrationForm{}

	err = schema.NewDecoder().Decode(&form, r.Form)
	if err != nil {
		viewError(w, r, "registration", err)
		return
	}

	err = form.validate()
	if err != nil {
		render(w, r, "registration", form, err)
		return
	}

	var u auth.User
	err = services.User.Register(&u, form.Username, form.FullName, form.Email, form.Password)
	if err != nil {
		render(w, r, "registration", form, err)
		return
	}

	redirect(w, r, "login")
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	if isAuthenticated(r) {
		redirect(w, r, "home")
		return
	}

	view(w, r, "login", nil)
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		viewError(w, r, "login", err)
		return
	}

	form := loginForm{}

	err = schema.NewDecoder().Decode(&form, r.Form)
	if err != nil {
		viewError(w, r, "login", err)
		return
	}

	var u auth.User
	err = services.User.Authenticate(&u, form.Username, form.Password)
	if err != nil {
		render(w, r, "login", form, err)
		return
	}

	var s auth.Session
	err = services.Session.Start(&s, u.ID)
	if err != nil {
		render(w, r, "login", form, err)
		return
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   s.Token,
		Expires: s.ExpiresAt,
	}
	http.SetCookie(w, &cookie)

	redirect(w, r, "home")
}

func getLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		redirect(w, r, "home")
		return
	}

	var s auth.Session
	err = services.Session.Authenticate(&s, cookie.Value)
	if err != nil {
		redirect(w, r, "home")
		return
	}

	s.ExpiresAt = time.Now()
	err = services.Session.Storer.Persist(&s)
	if err != nil {
		redirect(w, r, "home")
		return
	}

	redirect(w, r, "home")
}

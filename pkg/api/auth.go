package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func buildAuthRoutes(r *mux.Router) {
	r.Path("/register").
		Methods(http.MethodGet).
		HandlerFunc(getRegister).
		Name("registration")

	r.Path("/register").
		Methods(http.MethodPost).
		HandlerFunc(postRegister).
		Name("register")

	r.Path("/login").
		Methods(http.MethodGet).
		HandlerFunc(getLogin).
		Name("login")

	r.Path("/login").
		Methods(http.MethodPost).
		HandlerFunc(postLogin).
		Name("post-login")
}

func getRegister(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "registration", nil)
	if err != nil {
		log.Println(err)
	}
}

func postRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "login", nil)
	if err != nil {
		log.Println(err)
	}
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

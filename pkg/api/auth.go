package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

}

func getLogin(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "login", nil)
	if err != nil {
		log.Println(err)
	}
}

func postLogin(w http.ResponseWriter, r *http.Request) {

}

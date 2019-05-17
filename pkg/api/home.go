package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func buildHomeRoutes(r *mux.Router) {
	r.Path("/").
		Methods(http.MethodGet).
		HandlerFunc(getHome).
		Name("home")
}

func getHome(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "home", nil)
	if err != nil {
		log.Println(err)
	}
}

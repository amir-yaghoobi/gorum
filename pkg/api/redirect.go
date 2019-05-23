package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func redirect(w http.ResponseWriter, r *http.Request, route string) {
	url, err := mux.CurrentRoute(r).Subrouter().Get(route).URL()
	if err != nil {
		services.Logger.Errorf("redirect failed: %v", err)
	}

	http.Redirect(w, r, url.String(), http.StatusSeeOther)
}

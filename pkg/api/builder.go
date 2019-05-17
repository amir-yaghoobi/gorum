package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Build creates the HTTP handler to serve the website.
func Build() http.Handler {
	r := mux.NewRouter()
	return r
}

package api

import "net/http"

func buildHomeRoutes() {
	services.Router.
		Path("/").
		Methods(http.MethodGet).
		HandlerFunc(getHome).
		Name("home")
}

func getHome(w http.ResponseWriter, r *http.Request) {
	view(w, r, "home", nil)
}

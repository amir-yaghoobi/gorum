package api

import "net/http"

func buildHomeRoutes() {
	services.Router.
		Path("/").
		Methods(http.MethodGet).
		HandlerFunc(getHome).
		Name("home")
}

func getHome(w http.ResponseWriter, _ *http.Request) {
	view(w, "home", nil)
}

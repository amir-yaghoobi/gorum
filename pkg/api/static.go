package api

import "net/http"

func serveStaticFiles() {
	staticServer := http.FileServer(http.Dir("/usr/share/static"))

	services.Router.PathPrefix("/css").Handler(staticServer)
	services.Router.PathPrefix("/font").Handler(staticServer)
	services.Router.PathPrefix("/img").Handler(staticServer)
	services.Router.PathPrefix("/js").Handler(staticServer)
}

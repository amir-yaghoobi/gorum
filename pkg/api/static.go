package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func serveStaticFiles(r *mux.Router) {
	staticServer := http.FileServer(http.Dir("/usr/share/static"))

	r.PathPrefix("/css").Handler(staticServer)
	r.PathPrefix("/font").Handler(staticServer)
	r.PathPrefix("/img").Handler(staticServer)
	r.PathPrefix("/js").Handler(staticServer)
}

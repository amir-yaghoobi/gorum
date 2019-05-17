package api

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

// templ holds all parsed templates.
var templ *template.Template

func init() {
	templ = template.New("")
	err := filepath.Walk("/usr/share/templates", func(path string, _ os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(path, ".html") {
			_, err = templ.ParseFiles(path)
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Build creates the HTTP handler to serve the website.
func Build() http.Handler {
	r := mux.NewRouter()

	serveStaticFiles(r)
	buildAuthRoutes(r)
	buildHomeRoutes(r)

	return r
}

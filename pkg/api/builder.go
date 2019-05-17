package api

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
	"github.com/gorilla/mux"
)

var (
	// templ holds all parsed templates.
	templ *template.Template

	// messages contains message translations.
	messages *ini.File
)

func init() {
	var err error

	messages, err = ini.Load("/usr/share/config/messages.ini")
	if err != nil {
		log.Fatal(err)
	}

	templ = template.New("").Funcs(template.FuncMap{
		"trans": translate,
	})

	err = filepath.Walk("/usr/share/templates", func(path string, _ os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(path, ".html") {
			_, err = templ.ParseFiles(path)
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

func translate(message string) string {
	return messages.Section("").Key(message).String()
}

// Build creates the HTTP handler to serve the website.
func Build() http.Handler {
	r := mux.NewRouter()

	serveStaticFiles(r)
	buildAuthRoutes(r)
	buildHomeRoutes(r)

	return r
}

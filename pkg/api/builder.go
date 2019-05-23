package api

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
	"github.com/google/logger"
	"github.com/gorilla/mux"

	"gorum/pkg/auth"
)

// services contains all of shared services among routes.
var services struct {
	Logger   *logger.Logger
	Template *template.Template
	User     auth.UserService
	Session  auth.SessionService
}

const (
	messagesPath  = "/usr/share/config/messages.ini"
	templatesPath = "/usr/share/templates"
)

// Build creates the HTTP handler to serve the website.
func Build(userService auth.UserService, sessionService auth.SessionService) (http.Handler, error) {
	err := initContext(userService, sessionService)
	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()

	r.Use(authMiddleware)

	serveStaticFiles(r)
	buildAuthRoutes(r)
	buildHomeRoutes(r)

	return r, nil
}

func initContext(userService auth.UserService, sessionService auth.SessionService) error {
	services.User = userService
	services.Session = sessionService
	services.Logger = logger.Init("web", true, false, os.Stdout)

	messages, err := ini.Load(messagesPath)
	if err != nil {
		return err
	}

	services.Template = template.New("").Funcs(template.FuncMap{
		"trans": func(key string) string {
			return messages.Section("").Key(key).String()
		},
	})

	return filepath.Walk(templatesPath, func(path string, _ os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(path, ".html") {
			_, err = services.Template.ParseFiles(path)
		}
		return err
	})
}

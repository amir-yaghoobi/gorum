package api

import (
	"net/http"

	"gorum/pkg/auth"
)

func view(w http.ResponseWriter, r *http.Request, view string, data interface{}) {
	render(w, r, view, data, nil)
}

func viewError(w http.ResponseWriter, r *http.Request, view string, err error) {
	render(w, r, view, nil, err)
}

func render(w http.ResponseWriter, r *http.Request, view string, data interface{}, err error) {
	vd := struct {
		Data   interface{}
		User   *auth.User
		Errors []string
	}{
		Data: data,
		User: requestUser(r),
	}

	if err != nil {
		switch e := err.(type) {
		case *ValidationError:
			for _, err := range e.Errors {
				vd.Errors = append(vd.Errors, err.Error())
			}
		default:
			vd.Errors = append(vd.Errors, err.Error())
		}
	}

	err = services.Template.ExecuteTemplate(w, view, &vd)
	if err != nil {
		services.Logger.Error(err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

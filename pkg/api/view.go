package api

import "net/http"

func view(w http.ResponseWriter, view string, data interface{}) {
	render(w, view, data, nil)
}

func viewError(w http.ResponseWriter, view string, err error) {
	render(w, view, nil, err)
}

func render(w http.ResponseWriter, view string, data interface{}, err error) {
	vd := struct {
		Data interface{}
		Error string
	}{
		Data: data,
		Error: errorMessage(err),
	}

	err = services.Template.ExecuteTemplate(w, view, &vd)
	if err != nil {
		services.Logger.Error(err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

func errorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

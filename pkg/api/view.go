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
		Data   interface{}
		Errors []string
	}{
		Data: data,
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

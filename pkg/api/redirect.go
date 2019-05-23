package api

import "net/http"

func redirect(w http.ResponseWriter, r *http.Request, route string) {
	url, err :=services.Router.Get(route).URL()
	if err != nil {
		services.Logger.Errorf("redirect failed: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url.String(), http.StatusSeeOther)
}

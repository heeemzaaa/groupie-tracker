package groupie

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, 404, "Page not found !")
		return
	}
	data := Artists
	err := Tpl.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		ErrorPage(w, 500, "Internal server error !")
		return
	}
}

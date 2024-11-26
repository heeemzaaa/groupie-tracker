package groupie

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := Artists
	err := Tpl.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		fmt.Println("Error executing the file :", err)
		return
	}
}

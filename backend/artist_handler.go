package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../../frontend/html/artist.html")
	if err != nil {
		fmt.Println("Error parsing the file :", err)
		return
	}
	r.ParseForm()
	IdStr := r.FormValue("id")
	id, err := strconv.Atoi(IdStr)

	if err != nil || id >= len(Artists) {
		fmt.Println("error getting id")
		return
	}

	data := Artists[id-1]
	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing the file :", err)
		return
	}
}

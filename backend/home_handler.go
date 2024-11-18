package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../frontend/html/index.html")
	if err != nil {
		fmt.Println("Error parsing the file :", err)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := Artists
	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing the file :", err)
		return
	}
}

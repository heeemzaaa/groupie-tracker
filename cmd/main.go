package main

import (
	"fmt"
	"net/http"

	g "groupie/backend"
)

func init() {
	_ = g.ArtistData()
}

func main() {
	var err error
	g.Tpl, err = g.Tpl.ParseGlob("../frontend/html/*.html")
	if err != nil {
		fmt.Println("Error Parsing:", err)
		return
	}
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../frontend/css"))))
	http.HandleFunc("/", g.HomeHandler)
	http.HandleFunc("/artist", g.ArtistHandler)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

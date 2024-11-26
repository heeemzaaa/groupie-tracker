package main

import (
	"fmt"
	"net/http"

	g "groupie/internal"
)

func init() {
	_ = g.ArtistData()
}

func main() {
	var err error
	g.Tpl, err = g.Tpl.ParseGlob("../templates/*.html")
	if err != nil {
		fmt.Println("Error Parsing:", err)
		return
	}
	http.HandleFunc("/my-css/", g.CssHandler)
	http.HandleFunc("/", g.HomeHandler)
	http.HandleFunc("/artist", g.ArtistHandler)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

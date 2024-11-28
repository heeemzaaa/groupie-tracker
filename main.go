package main

import (
	"fmt"
	"log"
	"net/http"

	g "groupie/internal"
)

func init() {
	_ = g.ArtistData()
}

func main() {
	var err error
	g.Tpl, err = g.Tpl.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Error Parsing:", err)
		return
	}
	http.HandleFunc("/assets/", g.HandleAssets)
	http.HandleFunc("/", g.HomeHandler)
	http.HandleFunc("/artist", g.ArtistHandler)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

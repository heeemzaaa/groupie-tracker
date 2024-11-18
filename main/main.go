package main

import (
	g "groupie/backend"
	"net/http"
)

func init() {
	_ = g.ArtistData()
}

func main() {
	http.Handle("../frontend/css/", http.StripPrefix("../frontend/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", g.HomeHandler)
	http.HandleFunc("/artist", g.ArtistHandler)
	http.ListenAndServe(":8080", nil)
}

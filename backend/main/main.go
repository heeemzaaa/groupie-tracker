package main

import (
	"net/http"

	g "groupie/backend"
)

func init() {
	_ = g.ArtistData()
}

func main() {
	http.HandleFunc("/", g.HomeHandler)
	http.HandleFunc("/artist", g.ArtistHandler)
	http.ListenAndServe(":8080", nil)
}

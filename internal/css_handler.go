package groupie

import (
	"net/http"
	"os"
)

func CssHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) < len("/my-css/") || r.URL.Path[:len("/my-css/")] != "/my-css/" {
		http.Error(w, "403 Forbidden", http.StatusForbidden)
		return
	}
	fileName := r.URL.Path[len("/my-css"):]
	filePath := "../static/css" + fileName
	cssBytes, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading CSS file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/css")

	w.Write(cssBytes)
}

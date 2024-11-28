package groupie

import (
	"fmt"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	IdStr := r.FormValue("id")
	id, err := strconv.Atoi(IdStr)

	if err != nil || id > len(Artists) {
		fmt.Println("error getting id")
		return
	}
	FetchData(Artists[id-1].LocationsAPI, &Artists[id-1].Locations)
	FetchData(Artists[id-1].ConcertDatesAPI, &Artists[id-1].ConcertDates)
	FetchData(Artists[id-1].RelationsAPI, &Artists[id-1].Relations)
	data := Artists[id-1]
	err = Tpl.ExecuteTemplate(w, "artist.html", data)
	if err != nil {
		fmt.Println("Error executing the file :", err)
		return
	}
}

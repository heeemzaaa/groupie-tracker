package groupie

import "html/template"

type Artist struct {
	Id              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsAPI    string   `json:"locations"`
	ConcertDatesAPI string   `json:"concertDates"`
	RelationsAPI    string   `json:"relations"`
	Locations       Locations
	ConcertDates    ConcertDates
	Relations       Relations
}

type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

type ConcertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

var (
	Artists       []Artist
	Slocations    []Locations
	SconcertDates []ConcertDates
	Srelations    []Relations
	Tpl           *template.Template
)

type ErrorData struct {
	StatusCode int
	Message    string
}


package groupie

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
	Locations       []string
	ConcertDates    []string
	Relations       map[string][]string
}

var Artists []Artist

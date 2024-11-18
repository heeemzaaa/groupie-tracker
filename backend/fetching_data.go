package groupie

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchData(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data : %w", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

func ArtistData() error {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	err := FetchData(url, &Artists)
	if err != nil {
		return fmt.Errorf("error fetching artists: %w", err)
	}
	return nil
}

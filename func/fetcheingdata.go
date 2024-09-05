package groupie

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Group struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type GroupInfo struct {
	Group     Group
	Locations Locations
	Date      Date
	Relations Relations
}


// ArtistsData simplifies external API fetching
func ArtistsData(w http.ResponseWriter, url string, tar interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Failed to fetch data")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		HandleError(w, fmt.Errorf("Status: %d", resp.StatusCode), http.StatusInternalServerError, "Bad response from server")
		return fmt.Errorf("Bad response")
	}

	if err := json.NewDecoder(resp.Body).Decode(tar); err != nil {
		HandleError(w, err, http.StatusNotFound, "Not Found")
		return err
	}

	return nil
}

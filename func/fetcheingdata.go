package groupie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type Group struct {
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
}

type Locations struct {
	Id int `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	Id int `json:"id"`
	Dates []string `json:"dates"`
}
type Relations struct {
	Id int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type GroupInfo struct {
	Group Group
	Locations Locations
	Date Date
	Relations Relations
}

func ArtistsData(url string, tar interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Internal Server Error 500")
	}

	defer resp.Body.Close()
	
	body, err := CheckingResp(resp)
	if err != nil {
		return fmt.Errorf("Internal Server Error 500")
	}

	err = json.Unmarshal(body, &tar)
	if err != nil {
		return fmt.Errorf("Internal Server Error 500")
	}
	return nil
}

func CheckingResp(resp *http.Response) (body []byte, err error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch data: %s", resp.Status)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

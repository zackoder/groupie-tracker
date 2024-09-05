package groupie

import (
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

var Tmp *template.Template

func init() {
	// Parse all templates once globally
	Tmp = template.Must(template.ParseGlob("template/*.html"))
}

// Simplified error handling
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, nil, http.StatusNotFound, "Not Found")
		return
	}

	if r.Method != http.MethodGet {
		HandleError(w, nil, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var artists []Group
	if err := ArtistsData(w, "https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		return
	}

	if err := Tmp.ExecuteTemplate(w, "index.html", artists); err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Template rendering error")
		return
	}
}

func ArtistsInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, nil, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	path := r.URL.Path
	if !regexp.MustCompile(`/group/\d+$`).MatchString(path) {
		HandleError(w, nil, http.StatusNotFound, "Not Found")
		return
	}

	id := strings.TrimPrefix(path, "/group/")
	var group Group
	var locations Locations
	var concertDates Date
	var relations Relations

	apiUrls := map[string]interface{}{
		"https://groupietrackers.herokuapp.com/api/artists/" + id:   &group,
		"https://groupietrackers.herokuapp.com/api/locations/" + id: &locations,
		"https://groupietrackers.herokuapp.com/api/dates/" + id:     &concertDates,
		"https://groupietrackers.herokuapp.com/api/relation/" + id:  &relations,
	}

	for url, target := range apiUrls {
		if err := ArtistsData(w, url, target); err != nil {
			return
		}
	}

	if group.Id == 0 {
		HandleError(w, nil, http.StatusNotFound, "Not Found")
		return
	}

	groupInfo := GroupInfo{
		Group:     group,
		Locations: locations,
		Date:      concertDates,
		Relations: relations,
	}

	if err := Tmp.ExecuteTemplate(w, "groupinfo.html", groupInfo); err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Template rendering error")
	}
}

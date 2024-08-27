package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./template/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.Method != http.MethodGet {
		MethodNotAllowed(w)
		return
	}

	if r.URL.Path != "/" {
		NotFounderr(w)
		return
	}

	var artists []Group
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err = ArtistsData(url, &artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, artists)
}

func ArtistsInfo(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	patren := `/group/\d+$`
	re := regexp.MustCompile(patren)
	path := r.URL.Path
	if !re.MatchString(path) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	id := strings.TrimPrefix(path, "/group/")
	var group Group
	var locations Locations
	var consertDates Date
	var relations Relations

	groupUrl := "https://groupietrackers.herokuapp.com/api/artists/" + id
	locarionUrl := "https://groupietrackers.herokuapp.com/api/locations/" + id
	consertDatesUrl := "https://groupietrackers.herokuapp.com/api/dates/" + id
	relationsUrl := "https://groupietrackers.herokuapp.com/api/relation/" + id

	if err := ArtistsData(groupUrl, &group); err != nil {
		http.Error(w, "This group does not exist", http.StatusNotFound)
		fmt.Printf("Faild to fetch data from the Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t    ---------------------------------------------------\n", groupUrl)
		return
	}
	if err := ArtistsData(locarionUrl, &locations); err != nil {
		http.Error(w, "This group does not exist", http.StatusNotFound)
		fmt.Printf("Faild to fetch data from the Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t    ---------------------------------------------------\n", locarionUrl)
		return
	}
	if err := ArtistsData(consertDatesUrl, &consertDates); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		fmt.Printf("Faild to fetch data from the Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t\t    ---------------------------------------------------\n", consertDatesUrl)
		return
	}
	if err := ArtistsData(relationsUrl, &relations); err != nil {
		http.Error(w, "This group does not exist", http.StatusInternalServerError)
		fmt.Printf("Faild to fetch data from the Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t    -------------------------------------------------\n", relationsUrl)
		return
	}

	groupInfo := GroupInfo{
		Group:     group,
		Locations: locations,
		Date:      consertDates,
		Relations: relations,
	}
	tmp.ExecuteTemplate(w, "groupinfo.html", groupInfo)
}


func JsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/js" {
		NotFounderr(w)
		return
	}
}

func StyleH(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested URL path:", r.URL.Path) 
	if strings.HasPrefix(r.URL.Path, "/css")  {
		NotFounderr(w)
		return
	}
}

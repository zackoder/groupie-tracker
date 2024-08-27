package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./template/index.html")
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.Method != http.MethodGet {
		MethodNotAllowed(w)
		return
	}

	if r.URL.Path != "/" {
		NotFounderr(w);
		return
	}

	var artists []Group
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err = ArtistsData(url, &artists);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return
	}
	fmt.Println(artists[50].Members[1])
	tmp.Execute(w, artists)
}

// func ArtistsInfo(w http.ResponseWriter, r *http.Request) {
// 	tmp, err := template.ParseGlob("./template/*.html")
// 	if err != nil {
// 		http.Error(w, "This group does not exist", http.StatusInternalServerError)
// 		return
// 	}

// 	if r.Method != http.MethodGet {
// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 		return
// 	}
// 	patren := `/group/\d+$`
// 	re := regexp.MustCompile(patren)
// 	path := r.URL.Path
// 	if !re.MatchString(path) {
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	id := strings.TrimPrefix(path, "/group/")
// 	var group  Group
// 	var locations Locations
// 	var consertDates Date
// 	var relations Relations

// 	groupUrl := "https://groupietrackers.herokuapp.com/api/artists/" + id
// 	locarionUrl := "https://groupietrackers.herokuapp.com/api/locations/" + id
// 	consertDatesUrl := "https://groupietrackers.herokuapp.com/api/dates/" + id
// 	relationsUrl := "https://groupietrackers.herokuapp.com/api/relation/" + id

// 	if err := ArtistsData(groupUrl, &group); err != nil {
// 		http.Error(w, "This group does not exist", http.StatusInternalServerError)
// 		fmt.Printf("Faild to fetch data from ther Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t    ---------------------------------------------------\n", groupUrl)
// 		return
// 	}
// 	if err := ArtistsData(locarionUrl, &locations); err != nil {
// 		http.Error(w, "This group does not exist", http.StatusInternalServerError)
// 		fmt.Printf("Faild to fetch data from ther Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t    ---------------------------------------------------\n", locarionUrl)
// 		return
// 	}
// 	if err := ArtistsData(consertDatesUrl, &consertDates); err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		fmt.Printf("Faild to fetch data from ther Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t\t    ---------------------------------------------------\n", consertDatesUrl)
// 		return
// 	}
// 	if err := ArtistsData(relationsUrl, &relations); err != nil {
// 		http.Error(w, "This group does not exist", http.StatusInternalServerError)
// 		fmt.Printf("Faild to fetch data from ther Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t    -------------------------------------------------\n", relationsUrl)
// 		return
// 	}

// 	groupInfo := GroupInfo {
// 		Group: group,
// 		Locations: locations,
// 		Date: consertDates,
// 		Relations: relations,
// 	}
// 	tmp.ExecuteTemplate(w, "groupinfo.html", groupInfo)
// }

func ArtistsInfo(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("./template/*.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
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

	var wg sync.WaitGroup
	errChan := make(chan error, 4)

	id := strings.TrimPrefix(path, "/group/")
	var group Group
	var locations Locations
	var consertDates Date
	var relations Relations

	groupUrl := "https://proupietrackers.herokuapp.com/api/artists/" + id
	locarionUrl := "https://groupietrackers.herokuapp.com/api/locations/" + id
	consertDatesUrl := "https://groupietrackers.herokuapp.com/api/dates/" + id
	relationsUrl := "https://groupietrackers.herokuapp.com/api/relation/" + id

	wg.Add(4)
	go fetchArtistsData(groupUrl, &group, &wg, errChan)
	go fetchArtistsData(locarionUrl, &locations, &wg, errChan)
	go fetchArtistsData(consertDatesUrl, &consertDates, &wg, errChan)
	go fetchArtistsData(relationsUrl, &relations, &wg, errChan)

	wg.Wait()
	close(errChan)
	
	for err := range errChan {
		if err != nil {
			http.Error(w, "An error occurred", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}

	groupInfo := GroupInfo{
		Group:     group,
		Locations: locations,
		Date:      consertDates,
		Relations: relations,
	}
	tmp.ExecuteTemplate(w, "groupinfo.html", groupInfo)
}

func fetchArtistsData(url string, targ interface{}, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	if err := ArtistsData(url, &targ); err != nil {
		errChan <- fmt.Errorf("Faild to fetch data from ther Url:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t    ---------------------------------------------------\n", targ)
	}
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/js" {
		NotFounderr(w)
		return
	}
}

func StyleH(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/css" {
		NotFounderr(w)
		return
	}
}

/*

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchArtistsData(url string, data interface{}, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	if err := ArtistsData(url, data); err != nil {
		errChan <- fmt.Errorf("failed to fetch data from the URL:\n\t    ---------------------------------------------------\n\t|  %s  |\n\t    ---------------------------------------------------\n", url)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		group       GroupType       // replace with actual type
		locations   LocationsType   // replace with actual type
		consertDates ConsertDatesType // replace with actual type
		relations   RelationsType   // replace with actual type
	)

	var wg sync.WaitGroup
	errChan := make(chan error, 4)

	wg.Add(4)
	go fetchArtistsData(groupUrl, &group, &wg, errChan)
	go fetchArtistsData(locarionUrl, &locations, &wg, errChan)
	go fetchArtistsData(consertDatesUrl, &consertDates, &wg, errChan)
	go fetchArtistsData(relationsUrl, &relations, &wg, errChan)

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			http.Error(w, "An error occurred", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}

	// If no errors, continue processing
	// Your further logic here
}

func main() {
	http.HandleFunc("/your-endpoint", handler)
	http.ListenAndServe(":8080", nil)
}
*/

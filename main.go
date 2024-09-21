package main

import (
	"fmt"
	groupie "groupie/func"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Server setup
	port := ":3000"
	http.HandleFunc("/", groupie.Index)
	http.HandleFunc("/group/", groupie.ArtistsInfo)

	// Static file serving for js but restrict folder access
	jsDir := http.StripPrefix("/js/", http.FileServer(http.Dir("./js")))
	http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
		_, err := os.ReadFile("."+r.URL.Path)
		if err != nil {
			groupie.HandleError(w,nil,http.StatusNotFound, "Not Found")
			return
		}

		jsDir.ServeHTTP(w, r)
	})

	// static file serving for css
	cssDir := http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))
	http.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.Contains(r.URL.Path, "group") {
			path = strings.TrimPrefix(r.URL.Path, "/group/")
		}
		_, err := os.ReadFile("."+path)
		if err != nil {
			groupie.HandleError(w,nil,http.StatusNotFound, "Not Found")
			return
		}
		cssDir.ServeHTTP(w, r)
	})
	

	fmt.Printf("The server is listening on http://localhost%v\n", port)
	log.Panic(http.ListenAndServe(port, nil))
}

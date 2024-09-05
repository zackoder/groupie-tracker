package main

import (
	"fmt"
	groupie "groupie/func"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Server setup
	port := ":8080"
	http.HandleFunc("/", groupie.Index)
	http.HandleFunc("/group/", groupie.ArtistsInfo)

	// Static file serving for js but restrict folder access
	jsDir := http.StripPrefix("/js/", http.FileServer(http.Dir("./js")))
	http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/js/" || filepath.Ext(r.URL.Path) == "" {
			groupie.HandleError(w, nil, http.StatusForbidden, "Forbidden")
			return
		}
		jsDir.ServeHTTP(w, r)
	})

	// static file serving for css
	cssDir := http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))
	http.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/css/" || filepath.Ext(r.URL.Path) == "" {
			groupie.HandleError(w, nil, http.StatusForbidden, "Forbidden")
			return
		}
		cssDir.ServeHTTP(w, r)
	})
	

	fmt.Printf("The server is listening on http://localhost%v\n", port)
	log.Panic(http.ListenAndServe(port, nil))
}

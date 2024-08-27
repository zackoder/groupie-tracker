package main

import (
	"fmt"
	"groupie/func"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", groupie.Index)
	http.HandleFunc("/group/", groupie.ArtistsInfo)
	http.HandleFunc("/js", groupie.JsHandler)
	http.HandleFunc("/css", groupie.StyleH)

	cssDir := http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))
    http.Handle("/css", cssDir)

	jsDir := http.StripPrefix("/js/", http.FileServer(http.Dir("./js")))
    http.Handle("/js/", jsDir)
	
	fmt.Printf("the server is listning on http://localhost%v\n", port)
	log.Panic(http.ListenAndServe(port, nil))
}

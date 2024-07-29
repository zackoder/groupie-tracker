package main

import (
	"fmt"
	"net/http"
	"github.com/zackoder/groupie-tracker/groupie"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", groupie.Index)
	http.HandleFunc("/js", groupie.JsHandler)
	http.HandleFunc("/css", groupie.StyleH)
	fmt.Printf("the server is listning on http://localhost%v\n", port)
	http.ListenAndServe(port, nil)
}


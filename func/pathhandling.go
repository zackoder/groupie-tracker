package groupie

import (
	"fmt"
	"net/http"
)


func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return 
	}
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found 404", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello")
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "js" {
		http.Error(w, "Page Not Found 404", http.StatusNotFound)
		return
	}

}

func StyleH(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "css" {
		http.Error(w, "Page Not Found 404", http.StatusNotFound)
		return 
	}
}
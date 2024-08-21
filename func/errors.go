package groupie

import "net/http"

func NotFounderr(w http.ResponseWriter) {
	http.Error(w, "Page Not Found 404", http.StatusNotFound)
}

func MethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
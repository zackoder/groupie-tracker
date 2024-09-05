package groupie

import (
	"fmt"
	"net/http"
)

type ErrorMsg struct {
	Code int
	Msg  string
}

// Centralized error handling
func HandleError(w http.ResponseWriter, err error, status int, msg string) {
	Tmp.ParseFiles("err.html")
	Message := ErrorMsg{status, msg}
	w.WriteHeader(status)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if err := Tmp.ExecuteTemplate(w, "err.html", Message); err != nil {
		fmt.Println(err.Error())
	}
}

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
	Message := ErrorMsg{status, msg}
	w.WriteHeader(status)
	_, tmperr := Tmp.ParseFiles("err.html")
	if tmperr != nil {
		htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<link rel="stylesheet" href="/css/err.css" />
			<title>%d %s</title>
		</head>
		<body>
			<div class="err">
			<img src="https://c.tenor.com/ZoVfJwzAFIAAAAAC/tenor.gif" />
			<h1>Error: %d</h1>
			<h2>%s</h2>
			</div>
		</body>
		</html>
		`, status, msg, status, msg)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(htmlContent))
	} else {
		if err := Tmp.ExecuteTemplate(w, "err.html", Message); err != nil {
			fmt.Println(err.Error())
		}
	}
}

package main

import (
	"html/template"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// params := templateParams{}
	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/about.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}

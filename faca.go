package main

import (
	"html/template"
	"net/http"
)

func facaHandler(w http.ResponseWriter, r *http.Request) {
	// params := templateParams{}
	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/first-aid-cpr-aed.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}

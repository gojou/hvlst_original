package main

import (
	"html/template"
	"net/http"
)

func babysittingHandler(w http.ResponseWriter, r *http.Request) {
	// no need to handle 404 situations, will fall throuth REGEX
	// to the indexHandler

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/babysitting.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}

}

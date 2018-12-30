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

func coursesHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/courses.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}

func facaHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/first-aid-cpr-aed.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}

func wfaHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/wilderness-first-aid.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}

}

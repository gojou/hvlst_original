// package main
package main

import (
	//	"fmt"
	"html/template"
	"net/http"
	//	"time"
	//	"strconv"

	// "google.golang.org/appengine/datastore"
	// "google.golang.org/appengine/log"

	"google.golang.org/appengine"
)

type TemplateParams struct {
	Notice    string
	Id        int
	FirstName string
	LastName  string
	EmailAddr string
	Phone     string
	Message   string
	Contacts  []Contact
}

func main() {
	http.HandleFunc("/admin/messages", messagesHandler)
	http.HandleFunc("/courses/babysitting", babysittingHandler)
	http.HandleFunc("/courses/first-aid-cpr-aed", facaHandler)
	http.HandleFunc("/courses/wilderness-first-aid", wfaHandler)
	http.HandleFunc("/courses", coursesHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	params := TemplateParams{}

	// Set the default page

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/index.html",
	))

	//Display 404 if it's odball URL
	if r.URL.Path != "/" {
		page = template.Must(template.ParseFiles(
			"static/_base.html",
			"static/404.html",
		))

	}

	if r.Method == "GET" {
		page.Execute(w, params)
		//	return
	}

}

// package main
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"strconv"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"google.golang.org/appengine"
)

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	EmailAddr string
	Phone     string
	Message   string
	Posted    time.Time
}

type templateParams struct {
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
	http.HandleFunc("/courses/babysitting", babysittingHandler)
	http.HandleFunc("/courses/first-aid-cpr-aed", facaHandler)
	http.HandleFunc("/courses/wilderness-first-aid", wfaHandler)
	http.HandleFunc("/courses", coursesHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}

func facaHandler(w http.ResponseWriter, r *http.Request) {
	params := templateParams{}
	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/first-aid-cpr-aed.html",
	))

	if r.Method == "GET" {
		page.Execute(w, params)
		return
	}
}

func babysittingHandler(w http.ResponseWriter, r *http.Request) {
	params := templateParams{}

	// no need to handle 404 situations, will fall throuth REGEX
	// to the indexHandler

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/babysitting.html",
	))

	if r.Method == "GET" {
		page.Execute(w, params)
		return
	}

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	params := templateParams{}

	// no need to handle 404 situations, will fall throuth REGEX
	// to the indexHandler

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/about.html",
	))

	if r.Method == "GET" {
		page.Execute(w, params)
		return
	}

}

func wfaHandler(w http.ResponseWriter, r *http.Request) {
	params := templateParams{}

	// no need to handle 404 situations, will fall throuth REGEX
	// to the indexHandler

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/wilderness-first-aid.html",
	))

	if r.Method == "GET" {
		page.Execute(w, params)
		return
	}

}

func coursesHandler(w http.ResponseWriter, r *http.Request) {

	params := templateParams{}

	// no need to handle 404 situations, will fall throuth REGEX
	// to the indexHandler

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/courses.html",
	))

	if r.Method == "GET" {
		page.Execute(w, params)
		return
	}

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	params := templateParams{}

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/contact.html",
	))

	if r.Method == "GET" {
		page.Execute(w, params)
		return
	}

	// It's a POST request, so handle the form submission.

	id,err := strconv.Atoi(r.FormValue("id"))
	if err!=nil {
		params.Notice = "ID must be an integer."
		id=-1;
	}
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	emailAddr := r.FormValue("emailAddr")
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	params.Id = id
	params.FirstName = firstName // Preserve the firstName field.
	params.LastName = lastName   // Preserve the lastName field.
	params.EmailAddr = emailAddr // Preserve the emailAddr field.
	params.Phone = phone         // Preserve the phone field.
	params.Message = message     // Preserve the message field.

	if (firstName == "") || (lastName == "") || (emailAddr == "") {
		params.Notice = "First name, last name, and email are required."
		page.Execute(w, params)
		return
	}

	if r.FormValue("message") == "" {
		w.WriteHeader(http.StatusBadRequest)

		params.Notice = "Please send us a message."
		page.Execute(w, params)
		return
	}

	contact := Contact{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		EmailAddr: emailAddr,
		Phone:     phone,
		Message:   message,
		Posted:    time.Now(),
	}

	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "Contact", nil)

	if _, err := datastore.Put(ctx, key, &contact); err != nil {
		log.Errorf(ctx, "datastore.Put: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't add new post. Try again?"
		params.Message = contact.Message // Preserve their message so they can try again.
		page.Execute(w, params)
		return
	}
	params.Notice = fmt.Sprintf("Thank you for your submission, %s! %s", firstName, lastName)

	// [START execute]
	page.Execute(w, params)
	// [END execute]

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	params := templateParams{}

	// Set the default page

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/index.html",
	))

	//Overwrite the default page if it's an odball URL
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

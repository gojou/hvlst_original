// package main
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"google.golang.org/appengine"
)

// type templateParams struct {
// 	Notice  string
// 	Author  string
// 	Message string
// 	Posts   []Post
// }

// type Post struct {
// 	Author  string
// 	Message string
// 	Posted  time.Time
// }

type Post struct {
	FirstName string
	LastName  string
	EmailAddr string
	Phone     string
	Message   string
	Posted    time.Time
}

type prospectParams struct {
	Notice    string
	FirstName string
	LastName  string
	EmailAddr string
	Phone     string
	Message   string
	Posts     []Post
}

func main() {
	http.HandleFunc("/courses/babysitting", babysittingHandler)
	http.HandleFunc("/courses/first-aid-cpr-aed", facaHandler)
	http.HandleFunc("/courses", coursesHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}
func facaHandler(w http.ResponseWriter, r *http.Request) {
	params := prospectParams{}

	// no need to handle 404 situations, will fall throuth REGEX
	// to the indexHandler

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
	params := prospectParams{}

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

func coursesHandler(w http.ResponseWriter, r *http.Request) {

	params := prospectParams{}

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

	params := prospectParams{}

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/contact.html",
	))

	if r.Method == "GET" {
		page.Execute(w, params)
		return
	}

	// It's a POST request, so handle the form submission.

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	emailAddr := r.FormValue("emailAddr")
	phone := r.FormValue("phone")
	message := r.FormValue("message")

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

		params.Notice = "No message provided"
		page.Execute(w, params)
		return
	}

	// post := Post{
	// 	Author:  r.FormValue("author"),
	// 	Message: r.FormValue("message"),
	// 	Posted:  time.Now(),
	// }

	post := Post{
		FirstName: firstName,
		LastName:  lastName,
		EmailAddr: emailAddr,
		Phone:     phone,
		Message:   message,
		Posted:    time.Now(),
	}

	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "Post", nil)

	// Should be redundant -- Won't get to this handler with bad URL

	// if r.URL.Path != "/" {
	// 	http.Redirect(w, r, "/", http.StatusFound)
	// 	return
	// }

	if _, err := datastore.Put(ctx, key, &post); err != nil {
		log.Errorf(ctx, "datastore.Put: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't add new post. Try again?"
		params.Message = post.Message // Preserve their message so they can try again.
		page.Execute(w, params)
		return
	}
	params.Notice = fmt.Sprintf("Thank you for your submission, %s! %s", firstName, lastName)

	// [START execute]
	page.Execute(w, params)
	// [END execute]

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	params := prospectParams{}

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/index.html",
	))

	if r.Method == "GET" {
		page.Execute(w, params)
		return
	}

	// No longer posting anything from the index page.

	// // It's a POST request, so handle the form submission.
	//
	// author := r.FormValue("author")
	// params.Author = author // Preserve the name field.
	// if author == "" {
	// 	params.Notice = "A name is required"
	// 	page.Execute(w, params)
	// 	return
	// }
	//
	// if r.FormValue("message") == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	//
	// 	params.Notice = "No message provided"
	// 	page.Execute(w, params)
	// 	return
	// }
	//
	// post := Post{
	// 	Author:  r.FormValue("author"),
	// 	Message: r.FormValue("message"),
	// 	Posted:  time.Now(),
	// }
	//
	// ctx := appengine.NewContext(r)
	// key := datastore.NewIncompleteKey(ctx, "Post", nil)
	// if r.URL.Path != "/" {
	// 	http.Redirect(w, r, "/", http.StatusFound)
	// 	return
	// }
	//
	// if _, err := datastore.Put(ctx, key, &post); err != nil {
	// 	log.Errorf(ctx, "datastore.Put: %v", err)
	//
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	params.Notice = "Couldn't add new post. Try again?"
	// 	params.Message = post.Message // Preserve their message so they can try again.
	// 	page.Execute(w, params)
	// 	return
	// }
	// params.Notice = fmt.Sprintf("Thank you for your submission, %s!", author)
	//
	// [START execute]
	// page.Execute(w, params)
	// [END execute]

}

// package main
package lifeskills

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"google.golang.org/appengine"
)

type templateParams struct {
	Notice  string
	Author  string
	Message string
	Posts   []Post
}

type Post struct {
	Author  string
	Message string
	Posted  time.Time
}

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
)

// func main() {
//
// 	http.HandleFunc("/", indexHandler)
// 	appengine.Main() // Starts the server to receive requests
// }
func init() {

	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	params := templateParams{}

	if r.Method == "GET" {
		indexTemplate.Execute(w, params)
		return
	}

	// It's a POST request, so handle the form submission.

	author := r.FormValue("author")
	params.Author = author // Preserve the name field.
	if author == "" {
		params.Notice = "A name is required"
		indexTemplate.Execute(w, params)
		return
	}

	if r.FormValue("message") == "" {
		w.WriteHeader(http.StatusBadRequest)

		params.Notice = "No message provided"
		indexTemplate.Execute(w, params)
		return
	}

	post := Post{
		Author:  r.FormValue("author"),
		Message: r.FormValue("message"),
		Posted:  time.Now(),
	}

	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "Post", nil)
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if _, err := datastore.Put(ctx, key, &post); err != nil {
		log.Errorf(ctx, "datastore.Put: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't add new post. Try again?"
		params.Message = post.Message // Preserve their message so they can try again.
		indexTemplate.Execute(w, params)
		return
	}
	params.Notice = fmt.Sprintf("Thank you for your submission, %s!", author)

	// [START execute]
	indexTemplate.Execute(w, params)
	// [END execute]

}

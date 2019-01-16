//package main
package main

import (
	"html/template"
	"net/http"
	"google.golang.org/appengine"
)

func messagesHandler(w http.ResponseWriter, r *http.Request) {

	params := TemplateParams{}

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/viewmessages.html",
	))

	if r.Method == "GET" {
		ctx:= appengine.NewContext(r)
		params.Contacts = getContacts(ctx)
		page.Execute(w, params)
		return
	}

}

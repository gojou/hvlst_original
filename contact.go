//package main
package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
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

func contactHandler(w http.ResponseWriter, r *http.Request) {

	params := TemplateParams{}

	page := template.Must(template.ParseFiles(
		"static/_base.html",
		"static/contact.html",
	))

	if r.Method == "GET" {
		ctx := appengine.NewContext(r)
		params.Contacts = getContacts(ctx)
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

		params.Notice = "Please send us a message."
		page.Execute(w, params)
		return
	}

	contact := Contact{
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
	params.Notice = fmt.Sprintf("Thank you for your submission, %s!", firstName)
	notifyEmailMsg(ctx)
	// [START execute]
	page.Execute(w, params)
	// [END execute]

}

func getContacts(ctx context.Context) []Contact {

	params := TemplateParams{}

	q := datastore.NewQuery("Contact").Order("-Posted")

	var contacts []Contact

	if _, err := q.GetAll(ctx, &contacts); err != nil {
		params.Notice = "Didn't find any messages"
	}
	return contacts
}

func notifyEmailMsg(ctx context.Context) {
	msg := &mail.Message{
		Sender:  "mark.poling@gmail.com",
		To:      []string{"mark.poling@gmail.com"},
		Subject: "HVLST has a new contact!",
		Body:    fmt.Sprintf(confirmMsg),
	}
	if err := mail.Send(ctx, msg); err != nil {
		log.Errorf(ctx, "Could not send email: %v", err)
	}
}

const confirmMsg = `
HVLST has a new contact!
`

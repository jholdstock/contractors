// Package home displays the Home page.
package home

import (
	"net/http"

	"github.com/jholdstock/contractors/lib/flight"
	"github.com/jholdstock/contractors/lib/form"
	"github.com/jholdstock/contractors/lib/router"
)

// Load the routes.
func Load() {
	router.Get("/", Index)
}

// Index displays the home page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	v := c.View.New("home/index")

	if c.Sess.Values["id"] != nil { // Logged in
		v.Vars["first_name"] = c.Sess.Values["first_name"]
	} else { // Not logged in
		form.Repopulate(r.Form, v.Vars, "email")
	}

	v.Render(w, r)
}

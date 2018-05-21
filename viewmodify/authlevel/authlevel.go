// Package authlevel adds an LoggedIn and Admin variables to the view template.
package authlevel

import (
	"net/http"

	"github.com/jholdstock/contractors/lib/flight"
	"github.com/jholdstock/contractors/lib/view"
)

// Modify sets LoggedIn and Admin variables in the template to indicate if the user is authenticated and whether they are an admin
func Modify(w http.ResponseWriter, r *http.Request, v *view.Info) {
	c := flight.Context(w, r)

	if c.Sess.Values["id"] != nil {

		v.Vars["LoggedIn"] = true

		if c.Sess.Values["admin"] == true {
			v.Vars["Admin"] = true
		}
	} else {
		v.Vars["AuthLevel"] = "anon"
	}
}

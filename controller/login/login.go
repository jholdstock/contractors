// Package login handles the user login.
package login

import (
	"net/http"

	"github.com/jholdstock/contractors/lib/flight"
	"github.com/jholdstock/contractors/middleware/acl"
	"github.com/jholdstock/contractors/model/user"

	"github.com/jholdstock/contractors/controller/home"
	"github.com/jholdstock/contractors/lib/flash"
	"github.com/jholdstock/contractors/lib/passhash"
	"github.com/jholdstock/contractors/lib/router"
	"github.com/jholdstock/contractors/lib/session"
)

// Load the routes.
func Load() {
	router.Post("/login", Store, acl.DisallowAuth)
	router.Get("/logout", Logout)
}

// Store handles the login form submission.
func Store(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Validate with required fields
	if !c.FormValid("email", "password") {
		home.Index(w, r)
		return
	}

	// Form values
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Get database result
	result, noRows, err := user.ByEmail(c.DB, email)

	// Determine if user exists
	if noRows {
		c.FlashWarning("Password is incorrect")
	} else if err != nil {
		// Display error message
		c.FlashErrorGeneric(err)
	} else if passhash.MatchString(result.Password, password) {
		if result.StatusID != 1 {
			// User inactive and display inactive message
			c.FlashNotice("Account is inactive so login is disabled.")
		} else {
			// Login successfully
			session.Empty(c.Sess)
			c.Sess.AddFlash(flash.Info{
				Message: "Login successful!",
				Class:   flash.Success})
			c.Sess.Values["id"] = result.ID
			c.Sess.Values["email"] = email
			c.Sess.Values["admin"] = result.Admin
			c.Sess.Values["first_name"] = result.FirstName
			c.Sess.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	} else {
		c.FlashWarning("Password is incorrect")
	}

	// Show the login page again
	home.Index(w, r)
}

// Logout clears the session and logs the user out.
func Logout(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// If user is authenticated
	if c.Sess.Values["id"] != nil {
		session.Empty(c.Sess)
		c.FlashNotice("Goodbye!")
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

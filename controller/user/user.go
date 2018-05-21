// Package user provides a simple CRUD application in a web page.
package user

import (
	"net/http"

	"github.com/jholdstock/contractors/lib/flight"
	"github.com/jholdstock/contractors/middleware/acl"
	"github.com/jholdstock/contractors/model/user"

	"github.com/jholdstock/contractors/lib/pagination"
	"github.com/jholdstock/contractors/lib/router"
)

var (
	uri = "/users"
)

// Load the routes.
func Load() {
	router.Get(uri, Index, acl.AdminOnly)
}

// Index displays a paginated list of all users
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	p := pagination.New(r, 10)

	items, _, err := user.PaginateAll(c.DB, p.PerPage, p.Offset)
	if err != nil {
		c.FlashErrorGeneric(err)
		items = []user.Item{}
	}

	count, err := user.CountAll(c.DB)
	if err != nil {
		c.FlashErrorGeneric(err)
	}

	p.CalculatePages(count)

	v := c.View.New("user/index")
	v.Vars["items"] = items
	v.Vars["pagination"] = p
	v.Render(w, r)
}

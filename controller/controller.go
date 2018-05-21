// Package controller loads the routes for each of the controllers.
package controller

import (
	"github.com/jholdstock/contractors/controller/debug"
	"github.com/jholdstock/contractors/controller/home"
	"github.com/jholdstock/contractors/controller/invoice"
	"github.com/jholdstock/contractors/controller/login"
	"github.com/jholdstock/contractors/controller/register"
	"github.com/jholdstock/contractors/controller/static"
	"github.com/jholdstock/contractors/controller/status"
	"github.com/jholdstock/contractors/controller/user"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	invoice.Load()
	user.Load()
}

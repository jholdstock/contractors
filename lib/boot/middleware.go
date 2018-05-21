// Package boot handles the initialization of the web components.
package boot

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/csrf"

	"github.com/jholdstock/contractors/controller/status"
	"github.com/jholdstock/contractors/lib/flight"
	"github.com/jholdstock/contractors/middleware/logrequest"
	"github.com/jholdstock/contractors/middleware/rest"

	"github.com/justinas/alice"
)

// SetUpMiddleware contains the middleware that applies to every request.
func SetUpMiddleware(h http.Handler) http.Handler {
	return alice.New( // Chain middleware, top middleware runs first
		setUpCSRF,            // Prevent CSRF
		rest.Handler,         // Support changing HTTP method sent via query string
		logrequest.Handler,   // Log every request
		context.ClearHandler, // Prevent memory leak with gorilla.sessions
	).Then(h)
}

// setUpCSRF sets up the CSRF protection.
func setUpCSRF(h http.Handler) http.Handler {
	x := flight.Xsrf()

	// Decode the string
	key, err := base64.StdEncoding.DecodeString(x.AuthKey)
	if err != nil {
		log.Fatal(err)
	}

	// Configure the middleware
	cs := csrf.Protect([]byte(key),
		csrf.ErrorHandler(http.HandlerFunc(status.InvalidToken)),
		csrf.FieldName("_token"),
		csrf.Secure(x.Secure),
	)(h)
	return cs
}

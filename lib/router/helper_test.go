package router_test

import (
	"net/http"
	"testing"

	"github.com/jholdstock/contractors/lib/router"
)

// middlewareTest1 will modify a form variable.
func middlewareTest1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		r.Form.Set("foo1", "mt1")
	})
}

// middlewareTest2 will modify a form variable.
func middlewareTest2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		r.Form.Set("foo2", "mt2")
	})
}

// TestRouteList ensure the correct routes are returned.
func TestRouteList(t *testing.T) {
	// Reset the router
	router.ResetConfig()

	// Mock the HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Test all the handlers
	router.Get("/get", handler)
	router.Post("/post", handler)
	router.Delete("/delete", handler)
	router.Patch("/patch", handler)
	router.Put("/put", handler)

	testList := []string{
		"GET	/get",
		"POST	/post",
		"DELETE	/delete",
		"PATCH	/patch",
		"PUT	/put",
	}

	list := router.RouteList()

	if len(list) != len(testList) {
		t.Fatalf("\nactual: %v\nexpected: %v", len(list), len(testList))
	}

	for i := 0; i < len(list); i++ {
		actual := list
		expected := testList
		if actual[i] != expected[i] {
			t.Fatalf("\nactual: %v\nexpected: %v", actual[i], expected[i])
		}
	}
}

// TestRouteList ensure the correct routes are NOT returned.
func TestRouteListFail(t *testing.T) {
	// Reset the router
	router.ResetConfig()

	// Mock the HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Test all the handlers
	router.Get("/get", handler)
	router.Post("/post", handler)
	router.Delete("/delete", handler)
	router.Patch("/patch", handler)
	router.Put("/put", handler)

	testList := []string{
		"GET	/get",
		"POST	/post",
		"DELETE	/delete",
		"PATCH	/patch",
		//"PUT	/put",
	}

	list := router.RouteList()

	// These should not be equal now
	if len(list) == len(testList) {
		t.Fatalf("\nactual: %v\nexpected: %v", len(list), len(testList))
	}
}

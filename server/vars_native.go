//go:build go1.7
// +build go1.7

package server

import (
	"net/http"
)

// Vars is a helper function for accessing route
// parameters from any server.Router implementation. This is the equivalent
// of using `mux.Vars(r)` with the Gorilla mux.Router.
func Vars(r *http.Request) map[string]string {
	_ = "STUB: not implemented"
	// vars doesnt exist yet, return empty map
	return nil
}

// for some reason, vars is wrong type, return empty map

// SetRouteVars will set the given value into into the request context
// with the shared 'vars' storage key.
func SetRouteVars(r *http.Request, val interface{}) { _ = "STUB: not implemented"; return }

type contextKey int

// key to set/retrieve URL params from a
// Gorilla request context.
const varsKey contextKey = 2

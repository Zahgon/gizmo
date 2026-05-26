package kit

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router is an interface to wrap different router implementations.
type Router interface {
	Handle(method string, path string, handler http.Handler)
	HandleFunc(method string, path string, handlerFunc func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	SetNotFoundHandler(handler http.Handler)
}

// RouterOption sets optional Router overrides.
type RouterOption func(Router) Router

// RouterSelect allows users to override the
// default use of the Gorilla Router.
// TODO add type and constants for names + docs
func RouterSelect(name string) RouterOption { _ = "STUB: not implemented"; return *new(RouterOption) }

// CustomRouter allows users to inject an alternate Router implementation.
func CustomRouter(r Router) RouterOption { _ = "STUB: not implemented"; return *new(RouterOption) }

// RouterNotFound will set the not found handler of the router.
func RouterNotFound(h http.Handler) RouterOption {
	_ = "STUB: not implemented"
	return *new(RouterOption)
}

// StdlibRouter is a Router implementation for the Stdlib's `http.ServeMux`.
type stdlibRouter struct {
	mux *http.ServeMux
}

// Handle will call the Stdlib's HandleFunc() methods with a check for the incoming
// HTTP method. To allow for multiple methods on a single route, use 'ANY'.
func (g *stdlibRouter) Handle(method, path string, h http.Handler) {
	_ = "STUB: not implemented"
	return
}

// HandleFunc will call the Stdlib's HandleFunc() methods with a check for the incoming
// HTTP method. To allow for multiple methods on a single route, use 'ANY'.
func (g *stdlibRouter) HandleFunc(method, path string, h func(http.ResponseWriter, *http.Request)) {
	_ = "STUB: not implemented"
	return
}

// SetNotFoundHandler will do nothing as we cannot override the stdlib not found.
func (g *stdlibRouter) SetNotFoundHandler(h http.Handler) {
	_ = "STUB: not implemented"

	// ServeHTTP will call Stdlib's ServeMux.ServerHTTP directly.
	return
}

func (g *stdlibRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return

	// GorillaRouter is a Router implementation for the Gorilla web toolkit's `mux.Router`.
}

type gorillaRouter struct {
	mux *mux.Router
}

// Handle will call the Gorilla web toolkit's Handle().Method() methods.
func (g *gorillaRouter) Handle(method, path string, h http.Handler) {
	_ = "STUB: not implemented"
	return
}

// copy the route params into a shared location
// duplicating memory, but allowing Gizmo to be more flexible with
// router implementations.

// HandleFunc will call the Gorilla web toolkit's HandleFunc().Method() methods.
func (g *gorillaRouter) HandleFunc(method, path string, h func(http.ResponseWriter, *http.Request)) {
	_ = "STUB: not implemented"
	return
}

// SetNotFoundHandler will set the Gorilla mux.Router.NotFoundHandler.
func (g *gorillaRouter) SetNotFoundHandler(h http.Handler) { _ = "STUB: not implemented"; return }

// ServeHTTP will call Gorilla mux.Router.ServerHTTP directly.
func (g *gorillaRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return

	// Vars is a helper function for accessing route
	// parameters from any server.Router implementation. This is the equivalent
	// of using `mux.Vars(r)` with the Gorilla mux.Router.
}

func Vars(r *http.Request) map[string]string { _ = "STUB: not implemented"; return nil }

// SetRouteVars will set the given value into into the request context
// with the shared 'vars' storage key.
func SetRouteVars(r *http.Request, val interface{}) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

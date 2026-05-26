package server

import (
	"net/http"

	netContext "golang.org/x/net/context"
)

// SimpleServer is a basic http Server implementation for
// serving SimpleService, JSONService or MixedService implementations.
type SimpleServer struct {
	// tracks if the Register function is already called or not
	registered bool

	cfg *Config

	// exit chan for graceful shutdown
	exit chan chan error

	// mux for routing
	mux Router
	h   http.Handler

	svc Service

	// tracks active requests
	monitor *ActivityMonitor
}

// NewSimpleServer will init the mux, exit channel and
// build the address from the given port. It will register the HealthCheckHandler
// at the given path and set up the shutDownHandler to be called on Stop().
func NewSimpleServer(cfg *Config) *SimpleServer { _ = "STUB: not implemented"; return nil }

// ServeHTTP is SimpleServer's hook for metrics and safely executing each request.
func (s *SimpleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"

	// only count non-LB requests
	return
}

// UnexpectedServerError is returned with a 500 status code when SimpleServer recovers
// from a panic in a request.
var UnexpectedServerError = []byte("unexpected server error")

// executeRequestSafely will prevent a panic in a request from bringing the server down.
func (s *SimpleServer) safelyExecuteRequest(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// log the panic for all the details later

// give the users our deepest regrets

// lookup metric name if we can

// if we cant look up the metric name, dont bother. it'll use too much memory.

// Start will start the SimpleServer at it's configured address.
// If they are configured, this will start health checks and access logging.
func (s *SimpleServer) Start() error { _ = "STUB: not implemented"; return nil }

// only add the instrument handler if the proper router is enabled

// add TLS if in the configs

// join the LB

// let the health check clean up if it needs to

// stop the listener

// Stop initiates the shutdown process and returns when
// the server completes.
func (s *SimpleServer) Stop() error { _ = "STUB: not implemented"; return nil }

// Register will accept and register SimpleServer, JSONService or MixedService implementations.
func (s *SimpleServer) Register(svcI Service) error {
	_ = "STUB: not implemented"
	// check multiple register call error
	return nil
}

// set registered to true because we called it

// quick fix for backwards compatibility

// register all simple endpoints with our wrapper

// register all JSON endpoints with our wrapper

// register all context endpoints with our wrapper

// register all context endpoints with our wrapper

// set the function handle and register it to metrics

// GetForwardedIP returns the "X-Forwarded-For" header value.
func GetForwardedIP(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// GetIP returns the IP address for the given request.
func GetIP(r *http.Request) (string, error) { _ = "STUB: not implemented"; return "", nil }

// check real ip header first

// no nginx reverse proxy?
// get IP old fashioned way

// ContextKey used to create context keys.
type ContextKey int

const (
	// UserIPKey is key to set/retrieve value from context.
	UserIPKey ContextKey = 0

	// UserForwardForIPKey is key to set/retrieve value from context.
	UserForwardForIPKey ContextKey = 1
)

// ContextWithUserIP returns new context with user ip address.
func ContextWithUserIP(ctx netContext.Context, r *http.Request) netContext.Context {
	_ = "STUB: not implemented"
	return *new(netContext.Context)
}

// ContextWithForwardForIP returns new context with forward for ip.
func ContextWithForwardForIP(ctx netContext.Context, r *http.Request) netContext.Context {
	_ = "STUB: not implemented"
	return *new(netContext.Context)
}

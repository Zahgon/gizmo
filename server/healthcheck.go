package server

import (
	"net/http"
)

// HealthCheckHandler is an interface used by SimpleServer and RPCServer
// to allow users to customize their service's health check. Start will be
// called just before server start up and the given ActivityMonitor should
// offer insite to the # of requests in flight, if needed.
// Stop will be called once the servers receive a kill signal.
type HealthCheckHandler interface {
	http.Handler
	Path() string
	Start(*ActivityMonitor) error
	Stop() error
}

// SimpleHealthCheck is a basic HealthCheckHandler implementation
// that _always_ returns with an "ok" status and shuts down immediately.
type SimpleHealthCheck struct {
	path string
}

// NewSimpleHealthCheck will return a new SimpleHealthCheck instance.
func NewSimpleHealthCheck(path string) *SimpleHealthCheck { _ = "STUB: not implemented"; return nil }

// Path will return the configured status path to server on.
func (s *SimpleHealthCheck) Path() string {
	_ = "STUB: not implemented"

	// Start will do nothing.
	return ""
}

func (s *SimpleHealthCheck) Start(monitor *ActivityMonitor) error {
	_ = "STUB: not implemented"

	// Stop will do nothing and return nil.
	return nil
}

func (s *SimpleHealthCheck) Stop() error {
	_ = "STUB: not implemented"

	// ServeHTTP will always respond with "ok-"+server.Name.
	return nil
}

func (s *SimpleHealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// NewCustomHealthCheck will return a new CustomHealthCheck with the given
// path and handler.
func NewCustomHealthCheck(path string, handler http.Handler) *CustomHealthCheck {
	_ = "STUB: not implemented"
	return nil
}

// CustomHealthCheck is a HealthCheckHandler that uses
// a custom http.Handler provided to the server via `config.CustomHealthCheckHandler`.
type CustomHealthCheck struct {
	path    string
	handler http.Handler
}

// Path will return the configured status path to server on.
func (c *CustomHealthCheck) Path() string {
	_ = "STUB: not implemented"

	// Start will do nothing.
	return ""
}

func (c *CustomHealthCheck) Start(monitor *ActivityMonitor) error {
	_ = "STUB: not implemented"

	// Stop will do nothing and return nil.
	return nil
}

func (c *CustomHealthCheck) Stop() error {
	_ = "STUB: not implemented"

	// ServeHTTP will allow the custom handler to manage the request and response.
	return nil
}

func (c *CustomHealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

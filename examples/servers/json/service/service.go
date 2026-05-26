package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/server"

	"github.com/NYTimes/gizmo/examples/nyt"
)

type (
	// JSONService will implement server.JSONService and
	// handle all requests to the server.
	JSONService struct {
		client nyt.Client
	}
	// Config is a struct to contain all the needed
	// configuration for our JSONService
	Config struct {
		Server           *server.Config
		MostPopularToken string
		SemanticToken    string
	}
)

// NewJSONService will instantiate a JSONService
// with the given configuration.
func NewJSONService(cfg *Config) *JSONService { _ = "STUB: not implemented"; return nil }

// Prefix returns the string prefix used for all endpoints within
// this service.
func (s *JSONService) Prefix() string {
	_ = "STUB: not implemented"

	// Middleware provides an http.Handler hook wrapped around all requests.
	// In this implementation, we're using a GzipHandler middleware to
	// compress our responses.
	return ""
}

func (s *JSONService) Middleware(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// JSONMiddleware provides a JSONEndpoint hook wrapped around all requests.
// In this implementation, we're using it to provide application logging and to check errors
// and provide generic responses.
func (s *JSONService) JSONMiddleware(j server.JSONEndpoint) server.JSONEndpoint {
	_ = "STUB: not implemented"
	return *new(server.JSONEndpoint)
}

// JSONEndpoints is a listing of all endpoints available in the JSONService.
func (s *JSONService) JSONEndpoints() map[string]map[string]server.JSONEndpoint {
	_ = "STUB: not implemented"
	return nil
}

type jsonErr struct {
	Err string `json:"error"`
}

func (e *jsonErr) Error() string { _ = "STUB: not implemented"; return "" }

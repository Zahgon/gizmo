package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/server"

	"github.com/NYTimes/gizmo/examples/nyt"
)

type (
	// SimpleService will implement server.SimpleService and
	// handle all requests to the server.
	SimpleService struct {
		client nyt.Client
	}
	// Config is a struct to contain all the needed
	// configuration for our SimpleService
	Config struct {
		Server           *server.Config
		MostPopularToken string `envconfig:"MOST_POPULAR_TOKEN"`
		SemanticToken    string `envconfig:"SEMANTIC_TOKEN"`
	}
)

// NewSimpleService will instantiate a SimpleService
// with the given configuration.
func NewSimpleService(cfg *Config) *SimpleService { _ = "STUB: not implemented"; return nil }

// Prefix returns the string prefix used for all endpoints within
// this service.
func (s *SimpleService) Prefix() string {
	_ = "STUB: not implemented"

	// Middleware provides an http.Handler hook wrapped around all requests.
	// In this implementation, we're using a GzipHandler middleware to
	// compress our responses.
	return ""
}

func (s *SimpleService) Middleware(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// JSONEndpoints is a listing of all endpoints available in the SimpleService.
func (s *SimpleService) Endpoints() map[string]map[string]http.HandlerFunc {
	_ = "STUB: not implemented"
	return nil
}

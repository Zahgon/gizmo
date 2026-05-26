package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/server"

	"github.com/NYTimes/gizmo/examples/nyt"
)

type (
	// MixedService will implement server.MixedService and
	// handle all requests to the server.
	MixedService struct {
		client nyt.Client
	}
	// Config is a struct to contain all the needed
	// configuration for our MixedService
	Config struct {
		Server           *server.Config
		MostPopularToken string
		SemanticToken    string
	}
)

// NewMixedService will instantiate a MixedService
// with the given configuration.
func NewMixedService(cfg *Config) *MixedService { _ = "STUB: not implemented"; return nil }

// Prefix returns the string prefix used for all endpoints within
// this service.
func (s *MixedService) Prefix() string {
	_ = "STUB: not implemented"

	// Middleware provides an http.Handler hook wrapped around all requests.
	// In this implementation, we're using a GzipHandler middleware to
	// compress our responses.
	return ""
}

func (s *MixedService) Middleware(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// JSONMiddleware provides a JSONEndpoint hook wrapped around all requests.
// In this implementation, we're using it to provide application logging and to check errors
// and provide generic responses.
func (s *MixedService) JSONMiddleware(j server.JSONEndpoint) server.JSONEndpoint {
	_ = "STUB: not implemented"
	return *new(server.JSONEndpoint)
}

// Endpoints is a listing of all endpoints available in the MixedService.
func (s *MixedService) Endpoints() map[string]map[string]http.HandlerFunc {
	_ = "STUB: not implemented"
	return nil
}

// JSONEndpoints is a listing of all JSON endpoints available in the MixedService.
func (s *MixedService) JSONEndpoints() map[string]map[string]server.JSONEndpoint {
	_ = "STUB: not implemented"
	return nil
}

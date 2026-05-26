package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/pubsub"
	"github.com/NYTimes/gizmo/pubsub/aws"
	"github.com/NYTimes/gizmo/server"
)

type (
	// JSONPubService will implement server.JSONPubService and
	// handle all requests to the server.
	JSONPubService struct {
		pub pubsub.Publisher
	}
)

// NewJSONPubService will instantiate a JSONPubService
// with the given configuration.
func NewJSONPubService(cfg aws.SNSConfig) *JSONPubService { _ = "STUB: not implemented"; return nil }

// Prefix returns the string prefix used for all endpoints within
// this service.
func (s *JSONPubService) Prefix() string {
	_ = "STUB: not implemented"

	// Middleware provides an http.Handler hook wrapped around all requests.
	// In this implementation, we're using a GzipHandler middleware to
	// compress our responses.
	return ""
}

func (s *JSONPubService) Middleware(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// JSONMiddleware provides a JSONEndpoint hook wrapped around all requests.
// In this implementation, we're using it to provide application logging and to check errors
// and provide generic responses.
func (s *JSONPubService) JSONMiddleware(j server.JSONEndpoint) server.JSONEndpoint {
	_ = "STUB: not implemented"
	return *new(server.JSONEndpoint)
}

// JSONEndpoints is a listing of all endpoints available in the JSONPubService.
func (s *JSONPubService) JSONEndpoints() map[string]map[string]server.JSONEndpoint {
	_ = "STUB: not implemented"
	return nil
}

type jsonErr struct {
	Err string `json:"error"`
}

func (e *jsonErr) Error() string { _ = "STUB: not implemented"; return "" }

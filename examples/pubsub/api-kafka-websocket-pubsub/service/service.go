package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/pubsub/kafka"
)

// StreamService offers three endpoints: one to create a new topic in
// Kafka, a second to expose the topic over a websocket and a third
// to host a web page that provides a demo.
type StreamService struct {
	port int
	cfg  *kafka.Config
}

// NewStreamService will return a new stream service instance.
// If the given config is empty, it will default to localhost.
func NewStreamService(port int, cfg *kafka.Config) *StreamService {
	_ = "STUB: not implemented"
	return nil
}

// Prefix is the string prefixed to all endpoint routes.
func (s *StreamService) Prefix() string {
	_ = "STUB: not implemented"

	// Middleware in this service will do nothing.
	return ""
}

func (s *StreamService) Middleware(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Endpoints returns the two endpoints for our stream service.
func (s *StreamService) Endpoints() map[string]map[string]http.HandlerFunc {
	_ = "STUB: not implemented"
	return nil
}

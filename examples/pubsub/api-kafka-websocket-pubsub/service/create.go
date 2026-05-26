package service

import (
	"net/http"
)

// CreateStream is a JSON endpoint for creating a new topic in Kafka.
func (s *StreamService) CreateStream(r *http.Request) (int, interface{}, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func topicName(id int64) string { _ = "STUB: not implemented"; return "" }

func createTopic(name string) error { _ = "STUB: not implemented"; return nil }

type jsonErr struct {
	Err error `json:"error"`
}

func (e jsonErr) Error() string { _ = "STUB: not implemented"; return "" }

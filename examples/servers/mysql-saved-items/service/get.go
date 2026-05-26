package service

import (
	"net/http"
)

// Get is a JSONEndpoint to return a list of saved items for the given user ID.
func (s *SavedItemsService) Get(r *http.Request) (int, interface{}, error) {
	_ = "STUB: not implemented"
	// gather the input from the request
	return 0, nil, nil
}

// do work and respond

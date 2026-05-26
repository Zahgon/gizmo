package service

import (
	"net/http"
)

// Delete is JSONEndpoint for deleting a saved item from a user's list.
func (s *SavedItemsService) Delete(r *http.Request) (int, interface{}, error) {
	_ = "STUB: not implemented"
	// gather the inputs from request
	return 0, nil, nil
}

// do work and respond

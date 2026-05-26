package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/config/mysql"
	"github.com/NYTimes/gizmo/server"
)

// SavedItemsService will keep a handle on the saved items repository and implement
// the gizmo/server.JSONService interface.
type SavedItemsService struct {
	repo SavedItemsRepo
}

// NewSavedItemsService will attempt to instantiate a new repository and service.
func NewSavedItemsService(cfg *mysql.Config) (*SavedItemsService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prefix is to implement gizmo/server.Service interface. The string will be prefixed to all endpoint
// routes.
func (s *SavedItemsService) Prefix() string {
	_ = "STUB: not implemented"

	// Middleware provides a hook to add service-wide http.Handler middleware to the service.
	// In this example we are using it to add GZIP compression to our responses.
	// This method helps satisfy the server.Service interface.
	return ""
}

func (s *SavedItemsService) Middleware(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	// wrap the response with our GZIP Middleware
	return *new(http.Handler)
}

// JSONMiddleware provides a hook to add service-wide middleware for how JSONEndpoints
// should behave. In this example, we’re using the hook to check for a header to
// identify and authorize the user. This method helps satisfy the server.JSONService interface.
func (s *SavedItemsService) JSONMiddleware(j server.JSONEndpoint) server.JSONEndpoint {
	_ = "STUB: not implemented"
	return *new(server.JSONEndpoint)
}

// wrap our endpoint with an auth check and call it

// if the endpoint returns an unexpected error, return a generic message
// and log it.

// LogWithFields will add all the request context values
// to the structured log entry along some other request info

// idKey is a type to use as a key for storing data in the request context.
type idKey int

// userIDKey can be used to store/retrieve a user ID in a request context.
const userIDKey idKey = 0

// authCheck is a JSON middleware to check the request for a valid USER_ID
// header and set it into the request context. If the header is invalid
// or does not exist, a 401 response will be returned.
func authCheck(j server.JSONEndpoint) server.JSONEndpoint {
	_ = "STUB: not implemented"
	return *new(server.JSONEndpoint)
}

// check for User ID header injected by API Gateway

// verify it's an int

// reject request if bad/no user ID

// set the ID in context if we're good

// JSONEndpoints is the most important method of the Service implementation. It provides a
// listing of all endpoints available in the service with their routes and HTTP methods.
// This method helps satisfy the server.JSONService interface.
func (s *SavedItemsService) JSONEndpoints() map[string]map[string]server.JSONEndpoint {
	_ = "STUB: not implemented"
	return nil
}

type (
	// jsonResponse is a generic struct for responding with a simple JSON message.
	jsonResponse struct {
		Message string `json:"message"`
	}
	// jsonErr is a tiny helper struct to make displaying errors in JSON better.
	jsonErr struct {
		Err string `json:"error"`
	}
)

func (e *jsonErr) Error() string { _ = "STUB: not implemented"; return "" }

var (
	// ServiceUnavailableErr is a global error that will get returned when we are experiencing
	// technical issues.
	ServiceUnavailableErr = &jsonErr{"sorry, this service is currently unavailable"}
	// UnauthErr is a global error returned when the user does not supply the proper
	// authorization headers.
	UnauthErr = &jsonErr{"please include a valid USER_ID header in the request"}
)

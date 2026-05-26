package readinglist

import (
	"context"
	"net/http"

	ocontext "golang.org/x/net/context"
)

// gRPC stub
func (s service) GetListLimit(ctx ocontext.Context, r *GetListLimitRequest) (*Links, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// go-kit endpoint.Endpoint with core business logic
func (s service) getLinks(ctx context.Context, req interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// set request defaults
}

// set maximum

// get data from the service-injected DB interface

// request decoder can be used for proto and JSON since there is no body
func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

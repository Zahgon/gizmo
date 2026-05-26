package readinglist

import (
	"context"
	"net/http"

	ocontext "golang.org/x/net/context"
)

// gRPC stub
func (s service) PutLink(ctx ocontext.Context, r *PutLinkRequest) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// go-kit endpoint.Endpoint with core business logic
func (s service) putLink(ctx context.Context, req interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// validate the request
		nil
}

// call the service-injected DB interface

// JSON request decoder
func decodePutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Protobuf request decoder
func decodePutProtoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

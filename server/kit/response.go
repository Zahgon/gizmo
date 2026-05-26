package kit

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/proto"
)

// NewProtoStatusResponse allows users to respond with a specific HTTP status code and
// a Protobuf or JSON serialized response.
func NewProtoStatusResponse(res proto.Message, code int) *ProtoStatusResponse {
	_ = "STUB: not implemented"
	return nil
}

// ProtoStatusResponse implements:
// `httptransport.StatusCoder` to allow users to respond with the given
// response with a non-200 status code.
// `proto.Marshaler` and proto.Message so it can wrap a proto Endpoint responses.
// `json.Marshaler` so it can wrap JSON Endpoint responses.
// `error` so it can be used to respond as an error within the go-kit stack.
type ProtoStatusResponse struct {
	code int
	res  proto.Message
}

// StatusCode implements httptransport.StatusCoder and will return the given HTTP code.
func (c *ProtoStatusResponse) StatusCode() int {
	_ = "STUB: not implemented"

	// Marshal is to implement proto.Marshaler. It will marshal the given message, not this
	// struct.
	return 0
}

func (c *ProtoStatusResponse) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Reset is to implement proto.Message. It uses the given message's Reset method.
		nil
}

func (c *ProtoStatusResponse) Reset() {
	_ = "STUB: not implemented"

	// String is to implement proto.Message. It uses the given message's String method.
	return
}

func (c *ProtoStatusResponse) String() string { _ = "STUB: not implemented"; return "" }

// ProtoMessage is to implement proto.Message. It uses the given message's ProtoMessage
// method.
func (c *ProtoStatusResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

// MarshalJSON is to implement json.Marshaler. It will marshal the given message, not
// this struct.
func (c *ProtoStatusResponse) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ proto.Marshaler = &ProtoStatusResponse{}

// to implement error
func (c *ProtoStatusResponse) Error() string { _ = "STUB: not implemented"; return "" }

// NewJSONStatusResponse allows users to respond with a specific HTTP status code and
// a JSON serialized response.
func NewJSONStatusResponse(res interface{}, code int) *JSONStatusResponse {
	_ = "STUB: not implemented"
	return nil
}

// JSONStatusResponse implements:
// `httptransport.StatusCoder` to allow users to respond with the given
// response with a non-200 status code.
// `json.Marshaler` so it can wrap JSON Endpoint responses.
// `error` so it can be used to respond as an error within the go-kit stack.
type JSONStatusResponse struct {
	code int
	res  interface{}
}

// StatusCode is to implement httptransport.StatusCoder
func (c *JSONStatusResponse) StatusCode() int {
	_ = "STUB: not implemented"

	// MarshalJSON is to implement json.Marshaler
	return 0
}

func (c *JSONStatusResponse) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Error is to implement error
		nil
}

func (c *JSONStatusResponse) Error() string { _ = "STUB: not implemented"; return "" }

// EncodeProtoResponse is an httptransport.EncodeResponseFunc that serializes the response
// as Protobuf. Many Proto-over-HTTP services can use it as a sensible default. If the
// response implements Headerer, the provided headers will be applied to the response.
// If the response implements StatusCoder, the provided StatusCode will be used instead
// of 200.
func EncodeProtoResponse(ctx context.Context, w http.ResponseWriter, pres interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// maybe log instead? need to avoid a second header write

// maybe log instead? need to avoid a second header write

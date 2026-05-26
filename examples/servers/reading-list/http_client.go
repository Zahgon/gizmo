package readinglist

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Client struct {
	key string
	l   log.Logger

	put endpoint.Endpoint
	get endpoint.Endpoint
}

func NewClient(host string, l log.Logger, opts ...httptransport.ClientOption) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c Client) GetLinks(ctx context.Context, limit int) (*Links, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) PutLink(ctx context.Context, url string, delete bool) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodePut(ctx context.Context, r *http.Request, req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeGet(ctx context.Context, r *http.Request, req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeGetResp(ctx context.Context, r *http.Response) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodePutResp(ctx context.Context, r *http.Response) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mustParseURL(host, path string) *url.URL { _ = "STUB: not implemented"; return nil }

func retryEndpoint(e endpoint.Endpoint, l log.Logger) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

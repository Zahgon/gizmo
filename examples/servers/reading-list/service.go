package readinglist

import (
	"context"
	"net/http"

	"cloud.google.com/go/trace"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	"google.golang.org/grpc"
)

type service struct {
	db     DB
	tracer *trace.Client
}

// ensure we implement the gRPC service
var _ ReadingListServiceServer = &service{}

func NewService(db DB) (kit.Service, error) {
	_ = "STUB: not implemented"
	return *new(kit.Service), nil
}

func (s *service) HTTPOptions() []httptransport.ServerOption {
	_ = "STUB: not implemented"

	// override the default gorilla router and select the stdlib
	return nil
}

func (s *service) HTTPRouterOptions() []kit.RouterOption { _ = "STUB: not implemented"; return nil }

// tracing HTTP requests
func (s *service) HTTPMiddleware(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// the go-kit middleware is used for checking user authentication and
// injecting the current user into the request context.
func (s *service) Middleware(ep endpoint.Endpoint) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

// reject if user is not logged in

// add the user to the request context and continue

// declare the endpoints for the HTTP server
func (s *service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	_ = "STUB: not implemented"
	return nil
}

// tracing RPC requests
func (s *service) RPCMiddleware() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func (s *service) RPCServiceDesc() *grpc.ServiceDesc { _ = "STUB: not implemented"; return nil }

func (s *service) RPCOptions() []grpc.ServerOption { _ = "STUB: not implemented"; return nil }

const userKey = "oauth-user"

func getUserFromMD(ctx context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func addUser(ctx context.Context, usr string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getUser(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

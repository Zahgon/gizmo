package api

import (
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc"

	"github.com/NYTimes/gizmo/examples/nyt"
)

type (
	// service will implement kit.Service.
	service struct {
		client nyt.Client
	}
	// Config is a struct to contain all the needed
	// configuration for our Service
	Config struct {
		MostPopularToken string `envconfig:"MOST_POPULAR_TOKEN"`
		SemanticToken    string `envconfig:"SEMANTIC_TOKEN"`
	}
)

var _ ApiServiceServer = service{}

// NewService will instantiate a Service
// with the given configuration.
func New(cfg Config) kit.Service { _ = "STUB: not implemented"; return *new(kit.Service) }

func (s service) HTTPRouterOptions() []kit.RouterOption { _ = "STUB: not implemented"; return nil }

func (s service) HTTPOptions() []httptransport.ServerOption {
	_ = "STUB: not implemented"

	// HTTPMiddleware provides an http.Handler hook wrapped around all requests.
	// In this implementation, we're using a GzipHandler middleware to
	// compress our responses.
	return nil
}

func (s service) HTTPMiddleware(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Middleware provides an kit/endpoint.Middleware hook wrapped around all requests.
func (s service) Middleware(e endpoint.Endpoint) endpoint.Endpoint {
	_ = "STUB: not implemented"

	// JSONEndpoints is a listing of all endpoints available in the Service.
	// If using Cloud Endpoints, this is not needed but handy for local dev.
	return *new(endpoint.Endpoint)
}

func (s service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (s service) RPCMiddleware() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func (s service) RPCOptions() []grpc.ServerOption { _ = "STUB: not implemented"; return nil }

func (s service) RPCServiceDesc() *grpc.ServiceDesc {
	_ = "STUB: not implemented"
	// snagged from the pb.go file
	return nil
}

package gcp

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	v1pubsub "google.golang.org/api/pubsub/v1"

	"github.com/NYTimes/gizmo/pubsub"
)

var _ pubsub.MultiPublisher = &httpPublisher{}
var _ pubsub.Publisher = &httpPublisher{}

type httpPublisher struct {
	svc   *v1pubsub.ProjectsTopicsService
	topic string
}

// NewHTTPPublisher will instantiate a new GCP MultiPublisher that utilizes the HTTP client.
// This client is useful mainly for the App Engine standard environment as the gRPC client
// counts against the socket quota for some reason.
func NewHTTPPublisher(ctx context.Context, projID, topic string, src oauth2.TokenSource) (pubsub.MultiPublisher, error) {
	_ = "STUB: not implemented"
	return *new(pubsub.MultiPublisher), nil
}

// Publish will marshal the proto message and publish it to GCP pubsub.
func (p *httpPublisher) Publish(ctx context.Context, key string, msg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishRaw will publish the message to GCP pubsub.
func (p *httpPublisher) PublishRaw(ctx context.Context, key string, m []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishMulti will publish multiple messages to GCP pubsub in a single request.
func (p *httpPublisher) PublishMulti(ctx context.Context, keys []string, messages []proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishMultiRaw will publish multiple raw byte array messages to GCP pubsub in a single request.
func (p *httpPublisher) PublishMultiRaw(ctx context.Context, keys []string, messages [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

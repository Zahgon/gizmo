package pubsubtest

import (
	"sync"

	"github.com/NYTimes/gizmo/pubsub"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type (
	// TestPublisher is a simple implementation of pubsub.Publisher meant to
	// help mock out any implementations.
	TestPublisher struct {
		// Published will contain a list of all messages that have been published.
		Published []TestPublishMsg
		pmu       sync.Mutex

		// GivenError will be returned by the TestPublisher on publish.
		// Good for testing error scenarios.
		GivenError error

		// FoundError will contain any errors encountered while marshalling
		// the protobuf struct.
		FoundError error
	}
	// TestPublishMsg is a test publish message.
	TestPublishMsg struct {
		// Key represents the message key.
		Key string
		// Body represents the message body.
		Body []byte
	}
)

var _ pubsub.Publisher = &TestPublisher{}
var _ pubsub.MultiPublisher = &TestPublisher{}

// Publish publishes the message, delegating to PublishRaw.
func (t *TestPublisher) Publish(ctx context.Context, key string, msg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishRaw publishes the raw message byte slice.
func (t *TestPublisher) PublishRaw(_ context.Context, key string, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishMulti publishes the messages, delegating to Publish.
func (t *TestPublisher) PublishMulti(ctx context.Context, keys []string, messages []proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishMultiRaw will publish multiple raw byte array messages with a context.
func (t *TestPublisher) PublishMultiRaw(ctx context.Context, keys []string, messages [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

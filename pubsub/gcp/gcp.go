package gcp // import "github.com/NYTimes/gizmo/pubsub/gcp"

import (
	"sync"
	"time"

	gpubsub "cloud.google.com/go/pubsub"
	"github.com/NYTimes/gizmo/pubsub"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// Subscriber is a Google Cloud Platform PubSub client that allows a user to
// consume messages via the pubsub.Subscriber interface.
type Subscriber struct {
	sub subscription
	ctx context.Context

	mtxStop sync.Mutex
	stopped bool
	cancel  func()

	err error
}

// NewSubscriber will instantiate a new Subscriber that wraps a pubsub.Iterator.
func NewSubscriber(ctx context.Context, projID, subscription string, opts ...option.ClientOption) (*Subscriber, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	defaultMaxMessages  = 10
	defaultMaxExtension = 60 * time.Second
)

// Start will start pulling from pubsub via a pubsub.Iterator.
func (s *Subscriber) Start() <-chan pubsub.SubscriberMessage { _ = "STUB: not implemented"; return nil }

// Err will contain any error the Subscriber has encountered while processing.
func (s *Subscriber) Err() error {
	_ = "STUB: not implemented"

	// Stop will block until the consumer has stopped consuming messages.
	return nil
}

func (s *Subscriber) Stop() error { _ = "STUB: not implemented"; return nil }

// SetReceiveSettings sets the ReceivedSettings on the google pubsub Subscription.
// Should be called before Start().
func (s *Subscriber) SetReceiveSettings(settings gpubsub.ReceiveSettings) {
	_ = "STUB: not implemented"
	return
}

// SubMessage pubsub implementation of pubsub.SubscriberMessage.
type SubMessage struct {
	msg        message
	Attributes map[string]string
}

// Message will return the data of the pubsub Message.
func (m *SubMessage) Message() []byte { _ = "STUB: not implemented"; return nil }

// ExtendDoneDeadline will call the deprecated ModifyAckDeadline for a pubsub
// Message. This likely should not be called.
func (m *SubMessage) ExtendDoneDeadline(dur time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Done will acknowledge the pubsub Message.
func (m *SubMessage) Done() error { _ = "STUB: not implemented"; return nil }

// publisher is a Google Cloud Platform PubSub client that allows a user to
// consume messages via the pubsub.MultiPublisher interface.
type publisher struct {
	topic *gpubsub.Topic
}

var _ pubsub.Publisher = &publisher{}
var _ pubsub.MultiPublisher = &publisher{}

// NewPublisher will instantiate a new GCP MultiPublisher.
func NewPublisher(ctx context.Context, cfg Config, opts ...option.ClientOption) (pubsub.MultiPublisher, error) {
	_ = "STUB: not implemented"
	return *new(pubsub.MultiPublisher), nil
}

// Update PublishSettings from cfg.PublishSettings
// but never set thresholds to 0.

// Publish will marshal the proto message and publish it to GCP pubsub.
func (p *publisher) Publish(ctx context.Context, key string, msg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishRaw will publish the message to GCP pubsub.
func (p *publisher) PublishRaw(ctx context.Context, key string, m []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishMulti will publish multiple messages to GCP pubsub in a single request.
func (p *publisher) PublishMulti(ctx context.Context, keys []string, messages []proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishMultiRaw will publish multiple raw byte array messages to GCP pubsub in a single request.
func (p *publisher) PublishMultiRaw(ctx context.Context, keys []string, messages [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// interfaces and types to make this more testable
type (
	subscription interface {
		Receive(ctx context.Context, f func(context.Context, message)) error
	}
	message interface {
		ID() string
		MsgData() []byte
		Done()
	}

	messageImpl struct {
		Msg *gpubsub.Message
	}

	subscriptionImpl struct {
		Sub *gpubsub.Subscription
	}
)

func (m messageImpl) ID() string { _ = "STUB: not implemented"; return "" }

func (m messageImpl) MsgData() []byte { _ = "STUB: not implemented"; return nil }

func (m messageImpl) Done() { _ = "STUB: not implemented"; return }

func (s subscriptionImpl) Receive(ctx context.Context, f func(context.Context, message)) error {
	_ = "STUB: not implemented"
	return nil
}

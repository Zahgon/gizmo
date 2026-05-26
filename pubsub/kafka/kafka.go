package kafka // import "github.com/NYTimes/gizmo/pubsub/kafka"

import (
	"time"

	"github.com/NYTimes/gizmo/pubsub"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

var (
	// RequiredAcks will be used in Kafka configs
	// to set the 'RequiredAcks' value.
	RequiredAcks = sarama.WaitForAll
)

// Publisher is an experimental publisher that provides an implementation for
// Kafka using the Shopify/sarama library.
type Publisher struct {
	producer sarama.SyncProducer
	topic    string
}

// NewPublisher will initiate a new experimental Kafka publisher.
func NewPublisher(cfg *Config) (pubsub.Publisher, error) {
	_ = "STUB: not implemented"
	return *new(pubsub.Publisher), nil
}

// we always want successes to return

// Publish will marshal the proto message and emit it to the Kafka topic.
func (p *Publisher) Publish(ctx context.Context, key string, m proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishRaw will emit the byte array to the Kafka topic.
func (p *Publisher) PublishRaw(_ context.Context, key string, m []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: do something with this partition/offset values

// Stop will close the pub connection.
func (p *Publisher) Stop() error { _ = "STUB: not implemented"; return nil }

type (
	// subscriber is an experimental subscriber implementation for Kafka. It is only capable of consuming a
	// single partition so multiple may be required depending on your setup.
	subscriber struct {
		cnsmr     sarama.Consumer
		topic     string
		partition int32

		offset          func() int64
		broadcastOffset func(int64)

		kerr error

		stop chan chan error
	}

	// subMessage is an SubscriberMessage implementation
	// that will broadcast the message's offset when Done().
	subMessage struct {
		message         *sarama.ConsumerMessage
		broadcastOffset func(int64)
	}
)

// Message will return the message payload.
func (m *subMessage) Message() []byte { _ = "STUB: not implemented"; return nil }

// ExtendDoneDeadline has no effect on subMessage.
func (m *subMessage) ExtendDoneDeadline(time.Duration) error {
	_ = "STUB: not implemented"

	// Done will emit the message's offset.
	return nil
}

func (m *subMessage) Done() error { _ = "STUB: not implemented"; return nil }

// NewSubscriber will initiate a the experimental Kafka consumer.
func NewSubscriber(cfg *Config, offsetProvider func() int64, offsetBroadcast func(int64)) (pubsub.Subscriber, error) {
	_ = "STUB: not implemented"
	return *new(pubsub.Subscriber), nil
}

// we always want to see errors, no matter what

// Start will start consuming message on the Kafka topic
// partition and emit any messages to the returned channel.
// On start up, it will call the offset func provider to the subscriber
// to lookup the offset to start at.
// If it encounters any issues, it will populate the Err() error
// and close the returned channel.
func (s *subscriber) Start() <-chan pubsub.SubscriberMessage { _ = "STUB: not implemented"; return nil }

// TODO: what should we do here?

// Stop willablock until the consumer has stopped consuming messages
// and return any errors seen on consumer close.
func (s *subscriber) Stop() error { _ = "STUB: not implemented"; return nil }

// close result from the partition consumer

// Err will contain any  errors that occurred during
// consumption. This method should be checked after
// a user encounters a closed channel.
func (s *subscriber) Err() error {
	_ = "STUB: not implemented"

	// GetPartitions is a helper function to look up which partitions are available
	// via the given brokers for the given topic. This should be called only on startup.
	return nil
}

func GetPartitions(brokerHosts []string, topic string) (partitions []int32, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

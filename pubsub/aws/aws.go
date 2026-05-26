package aws // import "github.com/NYTimes/gizmo/pubsub/aws"

import (
	"time"

	"github.com/NYTimes/gizmo/pubsub"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// msgAttrsKey is the context key for the SNS Message Attributes.  Its value of zero is
// arbitrary. If this package defined other context keys, they would have
// different integer values.
const msgAttrsKey key = 0

// publisher will accept AWS credentials and an SNS topic name
// and it will emit any publish events to it.
type publisher struct {
	sns   snsiface.SNSAPI
	topic string
}

// NewPublisher will initiate the SNS client.
// If no credentials are passed in with the config,
// the publisher is instantiated with the AWS_ACCESS_KEY
// and the AWS_SECRET_KEY environment variables.
func NewPublisher(cfg SNSConfig) (pubsub.Publisher, error) {
	_ = "STUB: not implemented"
	return *new(pubsub.Publisher), nil
}

// Publish will marshal the proto message and emit it to the SNS topic.
// The key will be used as the SNS message subject.
func (p *publisher) Publish(ctx context.Context, key string, m proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishRaw will emit the byte array to the SNS topic.
// The key will be used as the SNS message subject.
// You can use func WithMessageAttributes to set SNS message attributes for the message
func (p *publisher) PublishRaw(ctx context.Context, key string, m []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// WithMessageAttributes used to add SNS Message Attributes to the context
// for further usage in publishing messages to sns with provided attributes
func WithMessageAttributes(ctx context.Context, msgAttrs map[string]*sns.MessageAttributeValue) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

var (
	// defaultSQSMaxMessages is default the number of bulk messages
	// the subscriber will attempt to fetch on each
	// receive.
	defaultSQSMaxMessages int64 = 10
	// defaultSQSTimeoutSeconds is the default number of seconds the
	// SQS client will wait before timing out.
	defaultSQSTimeoutSeconds int64 = 2
	// defaultSQSSleepInterval is the default time.Duration the
	// subscriber will wait if it sees no messages
	// on the queue.
	defaultSQSSleepInterval = 2 * time.Second

	// defaultSQSDeleteBufferSize is the default limit of messages
	// allowed in the delete buffer before
	// executing a 'delete batch' request.
	defaultSQSDeleteBufferSize = 0

	defaultSQSConsumeBase64 = true
)

func defaultSQSConfig(cfg *SQSConfig) { _ = "STUB: not implemented"; return }

type (
	// subscriber is an SQS client that allows a user to
	// consume messages via the pubsub.Subscriber interface.
	subscriber struct {
		sqs sqsiface.SQSAPI

		cfg      SQSConfig
		queueURL *string

		toDelete chan *deleteRequest
		// inFlight and stopped are signals to manage delete requests
		// at shutdown.
		inFlight uint64
		stopped  uint32

		stop   chan chan error
		sqsErr error
	}

	// SQSMessage is the SQS implementation of `SubscriberMessage`.
	subscriberMessage struct {
		sub     *subscriber
		message *sqs.Message
	}

	deleteRequest struct {
		entry   *sqs.DeleteMessageBatchRequestEntry
		receipt chan error
	}
)

// incrementInflight will increment the add in flight count.
func (s *subscriber) incrementInFlight() { _ = "STUB: not implemented"; return }

// removeInfFlight will decrement the in flight count.
func (s *subscriber) decrementInFlight() { _ = "STUB: not implemented"; return }

// inFlightCount returns the number of in-flight requests currently
// running on this server.
func (s *subscriber) inFlightCount() uint64 { _ = "STUB: not implemented"; return 0 }

// NewSubscriber will initiate a new Decrypter for the subscriber
// if a key file is provided. It will also fetch the SQS Queue Url
// and set up the SQS client.
func NewSubscriber(cfg SQSConfig) (pubsub.Subscriber, error) {
	_ = "STUB: not implemented"
	return *new(pubsub.Subscriber), nil
}

// Message will decode protobufed message bodies and simply return
// a byte slice containing the message body for all others types.
func (m *subscriberMessage) Message() []byte { _ = "STUB: not implemented"; return nil }

// ExtendDoneDeadline changes the visibility timeout of the underlying SQS
// message. It will set the visibility timeout of the message to the given
// duration.
func (m *subscriberMessage) ExtendDoneDeadline(d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Done will queue up a message to be deleted. By default,
// the `SQSDeleteBufferSize` will be 0, so this will block until the
// message has been deleted.
func (m *subscriberMessage) Done() error { _ = "STUB: not implemented"; return nil }

// Start will start consuming messages on the SQS queue
// and emit any messages to the returned channel.
// If it encounters any issues, it will populate the Err() error
// and close the returned channel.
func (s *subscriber) Start() <-chan pubsub.SubscriberMessage { _ = "STUB: not implemented"; return nil }

// get messages

// we've encountered a major error
// this will set the error value and close the channel
// so the user will stop iterating and check the err

// if we didn't get any messages, lets chill out for a sec

// for each message, pass to output

func (s *subscriber) handleDeletes() { _ = "STUB: not implemented"; return }

// if the subber is stopped and this is the last request,
// flush quit!

// if buffer is full, send the request

// cleaer buffer

// clear any remainders before shutdown

func (s *subscriber) isStopped() bool { _ = "STUB: not implemented"; return false }

// Stop will block until the consumer has stopped consuming
// messages.
func (s *subscriber) Stop() error { _ = "STUB: not implemented"; return nil }

// Err will contain any errors that occurred during
// consumption. This method should be checked after
// a user encounters a closed channel.
func (s *subscriber) Err() error {
	_ = "STUB: not implemented"

	// requestRoleCredentials return the credentials from AssumeRoleProvider to assume the role
	// referenced by the roleARN. If MFASerialNumber is specified, prompt for MFA token from stdin.
	return nil
}

func requestRoleCredentials(sess *session.Session, roleARN string, MFASerialNumber string) (*credentials.Credentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

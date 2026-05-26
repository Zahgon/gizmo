package pubsubtest

import (
	"time"

	"github.com/NYTimes/gizmo/pubsub"
	"github.com/golang/protobuf/proto"
)

type (
	// TestSubscriber is a simple implementation of pubsub.Subscriber meant to
	// help mock out any implementations.
	TestSubscriber struct {
		// ProtoMessages will be marshalled into []byte and used to mock out
		// a feed if it is populated.
		ProtoMessages []proto.Message

		// JSONMessages will be marshalled into []byte and used to mock out
		// a feed if it is populated.
		JSONMessages []interface{}

		// GivenErrError will be returned by the TestSubscriber on Err().
		// Good for testing error scenarios.
		GivenErrError error

		// GivenStopError will be returned by the TestSubscriber on Stop().
		// Good for testing error scenarios.
		GivenStopError error

		// FoundError will contain any errors encountered while marshalling
		// the JSON and protobuf struct.
		FoundError error
	}
	// TestSubsMessage represents a test subscriber message.
	TestSubsMessage struct {
		Msg         []byte
		DoneTimeout time.Duration
		Doned       bool
	}
)

// Message returns the subscriber message.
func (m *TestSubsMessage) Message() []byte {
	_ = "STUB: not implemented"

	// ExtendDoneDeadline changes the underlying DoneTimeout
	return nil
}

func (m *TestSubsMessage) ExtendDoneDeadline(d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Done sets the Doned field to true.
func (m *TestSubsMessage) Done() error { _ = "STUB: not implemented"; return nil }

// Start will populate and return the test channel for the subscriber
func (t *TestSubscriber) Start() <-chan pubsub.SubscriberMessage {
	_ = "STUB: not implemented"
	return nil
}

// Err returns the GivenErrError value.
func (t *TestSubscriber) Err() error { _ = "STUB: not implemented"; return nil }

// Stop returns the GivenStopError value.
func (t *TestSubscriber) Stop() error { _ = "STUB: not implemented"; return nil }

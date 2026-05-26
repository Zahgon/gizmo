package service

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func zeroOffset() int64 { _ = "STUB: not implemented"; return 0 }
func discardOffset(offset int64) {
	_ = "STUB: not implemented"

	// Stream will init a new pubsub.Publisher and pubsub.Subscriber
	// then upgrade the current request to a websocket connection. Any messages
	// consumed from Kafka will be published to the web socket and vice versa.
	return
}

func (s *StreamService) Stream(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// start consumer, emit to ws

// start producer, emit to kafka

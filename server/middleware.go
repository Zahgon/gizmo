package server

import (
	"bytes"
	"net/http"
)

// JSONToHTTP is the middleware func to convert a JSONEndpoint to
// an http.HandlerFunc.
func JSONToHTTP(ep JSONEndpoint) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// it's JSON, so always set that content type

// prepare to grab the response from the ep

// call the func and return err or not

// CORSHandler is a middleware func for setting all headers that enable CORS.
// If an originSuffix is provided, a strings.HasSuffix check will be performed
// before adding any CORS header. If an empty string is provided, any Origin
// header found will be placed into the CORS header. If no Origin header is
// found, no headers will be added.
func CORSHandler(f http.Handler, originSuffix string) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// NoCacheHandler is a middleware func for setting the Cache-Control to no-cache.
func NoCacheHandler(f http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// JSONPHandler is a middleware func for wrapping response body with JSONP.
func JSONPHandler(f http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// using a custom ResponseWriter so we can
// capture the response of the main request,
// wrap our JSONP stuff around it
// and only write to the actual response once

// add the JSONP only if the callback exists

// if no callback, just write the bytes

var (
	jsonpStart  = []byte("/**/")
	jsonpSecond = []byte("(")
	jsonpEnd    = []byte(");")
)

type jsonpResponseWriter struct {
	w   http.ResponseWriter
	buf bytes.Buffer
}

func (w *jsonpResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (w *jsonpResponseWriter) WriteHeader(h int) { _ = "STUB: not implemented"; return }

func (w *jsonpResponseWriter) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// JSONContextToHTTP is a middleware func to convert a ContextHandler an http.Handler.
		nil
}

func JSONContextToHTTP(ep JSONContextEndpoint) ContextHandler {
	_ = "STUB: not implemented"
	return *new(ContextHandler)
}

// it's JSON, so always set that content type

// prepare to grab the response from the ep

// call the func and return err or not

// ContextToHTTP is a middleware func to convert a ContextHandler an http.Handler.
func ContextToHTTP(ep ContextHandler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// WithCloseHandler returns a Handler cancelling the context when the client
// connection close unexpectedly.
func WithCloseHandler(h ContextHandler) ContextHandler {
	_ = "STUB: not implemented"
	return *new(ContextHandler)
}

// Cancel the context if the client closes the connection

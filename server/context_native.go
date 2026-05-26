//go:build go1.7
// +build go1.7

package server

import (
	"net/http"
)

// AddIPToContext will attempt to pull an IP address out of the request and
// set it into a gorilla context.
func AddIPToContext(r *http.Request) { _ = "STUB: not implemented"; return }

// ContextFields will take a request and convert a context map to logrus Fields.
func ContextFields(r *http.Request) map[string]interface{} { _ = "STUB: not implemented"; return nil }

// gorilla.mux adds the route to context.
// we want to remove it for now

// web.varsKey for _all_ mux variables (gorilla or httprouter)

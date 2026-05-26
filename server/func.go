package server

import (
	"net/http"
)

// JSONContentType can be used for setting the Content-Type header for JSON encoding.
const JSONContentType = "application/json; charset=UTF-8"

// GetInt64Var is a helper to pull gorilla mux Vars.
// If the value is empty, it falls back to the URL
// query string.
// We are ignoring the error here bc we're assuming
// the path had a [0-9]+ descriptor on this var.
func GetInt64Var(r *http.Request, key string) int64 { _ = "STUB: not implemented"; return 0 }

// GetUInt64Var is a helper to pull gorilla mux Vars.
// If the value is empty, it falls back to the URL
// query string.
// We are ignoring the error here bc we're assuming
// the path had a [0-9]+ descriptor on this var.
func GetUInt64Var(r *http.Request, key string) uint64 { _ = "STUB: not implemented"; return 0 }

// ParseTruthyFalsy is a helper method to attempt to parse booleans in
// APIs that have no set contract on what a boolean should look like.
func ParseTruthyFalsy(flag interface{}) (result bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

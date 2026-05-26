package auth

import (
	"context"
	"crypto/rsa"
	"net/http"
	"regexp"
	"sync"
	"time"
)

// PublicKeySource is to be used by servers who need to acquire public key sets for
// verifying inbound request's JWTs.
type PublicKeySource interface {
	Get(context.Context) (PublicKeySet, error)
}

// NewReusePublicKeySource is a wrapper around PublicKeySources to only fetch a new key
// set once the current key cache has expired.
func NewReusePublicKeySource(ks PublicKeySet, src PublicKeySource) PublicKeySource {
	_ = "STUB: not implemented"
	return *new(PublicKeySource)
}

type reuseKeySource struct {
	src PublicKeySource

	mu sync.Mutex
	ks PublicKeySet
}

func (r *reuseKeySource) Get(ctx context.Context) (PublicKeySet, error) {
	_ = "STUB: not implemented"
	return *new(PublicKeySet), nil
}

// PublicKeySet contains a set of keys acquired from a JWKS that has an expiration.
type PublicKeySet struct {
	Expiry time.Time
	Keys   map[string]*rsa.PublicKey
}

// Expired will return true if the current key set is expire according to its Expiry
// field.
func (ks PublicKeySet) Expired() bool { _ = "STUB: not implemented"; return false }

// GetKey will look for the given key ID in the key set and return it, if it exists.
func (ks PublicKeySet) GetKey(id string) (*rsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JSONKey represents a public or private key in JWK format.
type JSONKey struct {
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// JSONKeyResponse represents a JWK Set object.
type JSONKeyResponse struct {
	Keys []*JSONKey `json:"keys"`
}

var reMaxAge = regexp.MustCompile("max-age=([0-9]*)")

// NewPublicKeySetFromURL will attempt to fetch a JWKS from the given URL and parse it
// into a PublicKeySet. The endpoint the URL points to must return the same format as the
// JSONKeyResponse struct.
func NewPublicKeySetFromURL(hc *http.Client, url string, defaultTTL time.Duration) (PublicKeySet, error) {
	_ = "STUB: not implemented"
	return *new(PublicKeySet), nil
}

// NewPublicKeySetFromJSON will accept a JSON payload in the format of the
// JSONKeyResponse and parse it into a PublicKeySet.
func NewPublicKeySetFromJSON(payload []byte, ttl time.Duration) (PublicKeySet, error) {
	_ = "STUB: not implemented"
	return *new(PublicKeySet), nil
}

// we only plan on using RSA

// TimeNow is used internally to determine the current time. It has been abstracted to
// this global function as a mechanism to help with testing.
var TimeNow = func() time.Time { return time.Now() }

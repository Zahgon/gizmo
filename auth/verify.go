package auth // import "github.com/NYTimes/gizmo/auth"

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2/jws"
)

// Verifier is a generic tool for verifying JWT tokens.
type Verifier struct {
	ks PublicKeySource
	df ClaimsDecoderFunc
	vf VerifyFunc

	skewAllowance int64
}

// ErrBadCreds will always be wrapped when a user's
// credentials are unexpected. This is so that we can
// distinguish between a client error from a server error
var ErrBadCreds = errors.New("bad credentials")

var defaultSkewAllowance = time.Minute * 5

// ClaimSetter is an interface for all incoming claims to implement. This ensures the
// basic format used by the `jws` package.
type ClaimSetter interface {
	BaseClaims() *jws.ClaimSet
}

// ClaimsDecoderFunc will expect to convert a JSON payload into the appropriate claims
// type.
type ClaimsDecoderFunc func(context.Context, []byte) (ClaimSetter, error)

// VerifyFunc will be called by the Verify if all other checks on the token pass.
// Developers should use this to encapsulate any business logic involved with token
// verification.
type VerifyFunc func(context.Context, interface{}) bool

// NewVerifier returns a genric Verifier that will use the given funcs and key source.
func NewVerifier(ks PublicKeySource, df ClaimsDecoderFunc, vf VerifyFunc) *Verifier {
	_ = "STUB: not implemented"
	return nil
}

// VerifyInboundKitContext is meant to be used within a go-kit stack that has populated
// the context with common headers, specficially
// kit/transport/http.ContextKeyRequestAuthorization.
func (c Verifier) VerifyInboundKitContext(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// VerifyRequest will pull the token from the "Authorization" header of the inbound
// request then decode and verify it.
func (c Verifier) VerifyRequest(r *http.Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Verify will accept an opaque JWT token, decode it and verify it.
func (c Verifier) Verify(ctx context.Context, token string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// use claims decoder func

func decodeToken(token string) (*jws.Header, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseHeader(hdr string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetAuthorizationToken will pull the Authorization header from the given request and
// attempt to retrieve the token within it.
func GetAuthorizationToken(r *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

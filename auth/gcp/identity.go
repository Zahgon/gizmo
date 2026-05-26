package gcp

import (
	"context"
	"net/http"

	"github.com/NYTimes/gizmo/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jws"
)

// IdentityConfig contains the information required for generating or verifying identity
// JWTs.
type IdentityConfig struct {
	Audience string `envconfig:"ID_AUDIENCE"`

	CertURL string `envconfig:"ID_CERT_URL"` // optional override for public key source

	Client *http.Client // optional override

	MetadataAddress string `envconfig:"ID_METADATA_ADDR"` // optional override for token and email retrieval
}

type idKeySource struct {
	cfg IdentityConfig
}

// NewDefaultIdentityVerifier will verify tokens that have the same default service
// account as the server running this verifier.
func NewDefaultIdentityVerifier(ctx context.Context, cfg IdentityConfig) (*auth.Verifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIdentityPublicKeySource fetches Google's public oauth2 certificates to be used with
// the auth.Verifier tool.
func NewIdentityPublicKeySource(ctx context.Context, cfg IdentityConfig) (auth.PublicKeySource, error) {
	_ = "STUB: not implemented"
	return *new(auth.PublicKeySource), nil
}

func (s idKeySource) Get(ctx context.Context) (auth.PublicKeySet, error) {
	_ = "STUB: not implemented"
	return *new(auth.PublicKeySet), nil
}

// NewIdentityTokenSource will use the GCP metadata services to generate GCP Identity
// tokens. More information on asserting GCP identities can be found here:
// https://cloud.google.com/compute/docs/instances/verifying-instance-identity
func NewIdentityTokenSource(cfg IdentityConfig) (oauth2.TokenSource, error) {
	_ = "STUB: not implemented"
	return *new(oauth2.TokenSource), nil
}

type idTokenSource struct {
	cfg IdentityConfig
}

func (c *idTokenSource) Token() (*oauth2.Token, error) { _ = "STUB: not implemented"; return nil, nil }

// IdentityClaimSet holds all the expected values for the various versions of the GCP
// identity token.
// More details:
// https://cloud.google.com/compute/docs/instances/verifying-instance-identity#payload
// https://developers.google.com/identity/sign-in/web/backend-auth#calling-the-tokeninfo-endpoint
type IdentityClaimSet struct {
	jws.ClaimSet

	// Email address of the default service account (only exists on GAE 2nd gen?)
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`

	// Google metadata info (appears to only exist on GCE?)
	Google map[string]interface{} `json:"google"`
}

// BaseClaims implements the auth.ClaimSetter interface.
func (s IdentityClaimSet) BaseClaims() *jws.ClaimSet {
	_ = "STUB: not implemented"

	// IdentityClaimsDecoderFunc is an auth.ClaimsDecoderFunc for GCP identity tokens.
	return nil
}

func IdentityClaimsDecoderFunc(_ context.Context, b []byte) (auth.ClaimSetter, error) {
	_ = "STUB: not implemented"
	return *new(auth.ClaimSetter), nil
}

// IdentityVerifyFunc auth.VerifyFunc wrapper around the IdentityClaimSet.
func IdentityVerifyFunc(vf func(ctx context.Context, cs IdentityClaimSet) bool) auth.VerifyFunc {
	_ = "STUB: not implemented"
	return *new(auth.VerifyFunc)
}

// Issuers contains the known Google account issuers for identity tokens.
var Issuers = map[string]bool{
	"accounts.google.com":         true,
	"https://accounts.google.com": true,
}

// ValidIdentityClaims ensures the token audience and issuers match expectations.
func ValidIdentityClaims(cs IdentityClaimSet, audience string) bool {
	_ = "STUB: not implemented"
	return false
}

// VerifyIdentityEmails is an auth.VerifyFunc that ensures IdentityClaimSets are valid
// and have the expected email and audience in their payload.
func VerifyIdentityEmails(ctx context.Context, emails []string, audience string) auth.VerifyFunc {
	_ = "STUB: not implemented"
	return *new(auth.VerifyFunc)
}

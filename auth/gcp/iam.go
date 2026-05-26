package gcp

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/NYTimes/gizmo/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jws"
	iam "google.golang.org/api/iam/v1"
)

var (
	timeNow = func() time.Time { return time.Now() }

	// docs say up to 1 hour, this plays it safe?
	// https://cloud.google.com/compute/docs/instances/verifying-instance-identity#verify_signature
	defaultTokenTTL = time.Minute * 20
)

// IAMClaimSet contains just an email for service account identification.
type IAMClaimSet struct {
	jws.ClaimSet

	// Email address of the default service account
	Email string `json:"email"`
}

// NewDefaultIAMVerifier will verify tokens that have the same default service account as
// the server running this verifier.
func NewDefaultIAMVerifier(ctx context.Context, cfg IAMConfig, clientFunc func(context.Context) *http.Client) (*auth.Verifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BaseClaims implements the auth.ClaimSetter interface.
func (s IAMClaimSet) BaseClaims() *jws.ClaimSet {
	_ = "STUB: not implemented"

	// IAMClaimsDecoderFunc is an auth.ClaimsDecoderFunc for GCP identity tokens.
	return nil
}

func IAMClaimsDecoderFunc(_ context.Context, b []byte) (auth.ClaimSetter, error) {
	_ = "STUB: not implemented"
	return *new(auth.ClaimSetter), nil
}

// IAMVerifyFunc auth.VerifyFunc wrapper around the IAMClaimSet.
func IAMVerifyFunc(vf func(ctx context.Context, cs IAMClaimSet) bool) auth.VerifyFunc {
	_ = "STUB: not implemented"
	return *new(auth.VerifyFunc)
}

// ValidIAMClaims ensures the token audience issuers matches expectations.
func ValidIAMClaims(cs IAMClaimSet, audience string) bool { _ = "STUB: not implemented"; return false }

// VerifyIAMEmails is an auth.VerifyFunc that ensures IAMClaimSets are valid
// and have the expected email and audience in their payload.
func VerifyIAMEmails(ctx context.Context, emails []string, audience string) auth.VerifyFunc {
	_ = "STUB: not implemented"
	return *new(auth.VerifyFunc)
}

type iamKeySource struct {
	cf  func(context.Context) *http.Client
	cfg IAMConfig
}

// NewIAMPublicKeySource returns a PublicKeySource that uses the Google IAM service
// for fetching public keys of a given service account. The function for returning an
// HTTP client is to allow 1st generation App Engine users to lean on urlfetch.
func NewIAMPublicKeySource(ctx context.Context, cfg IAMConfig, clientFunc func(context.Context) *http.Client) (auth.PublicKeySource, error) {
	_ = "STUB: not implemented"
	return *new(auth.PublicKeySource), nil
}

func (s iamKeySource) Get(ctx context.Context) (auth.PublicKeySet, error) {
	_ = "STUB: not implemented"
	return *

	// for the sake of GAE standard users who have to use a different *http.Client on
	// each request, we're going to init a new iam.Service on each fetch.
	// since this is cached, it should hopefully not be a huge issue
	new(auth.PublicKeySet), nil
}

// we need to fetch each key's PublicKey data since List only returns metadata.

// IAMConfig contains the information required for generating or verifying IAM JWTs.
type IAMConfig struct {
	IAMAddress string `envconfig:"IAM_ADDR"` // optional, for testing

	Audience            string `envconfig:"IAM_AUDIENCE"`
	Project             string `envconfig:"IAM_PROJECT"`
	ServiceAccountEmail string `envconfig:"IAM_SERVICE_ACCOUNT_EMAIL"`

	// JSON contains the raw bytes from a JSON credentials file.
	// This field may be nil if authentication is provided by the
	// environment and not with a credentials file, e.g. when code is
	// running on Google Cloud Platform.
	JSON []byte
}

// NewIAMTokenSource returns an oauth2.TokenSource that uses Google's IAM services
// to sign a JWT with the default service account and the given audience.
// Users should use the Identity token source if they can. This client is meant to be
// used as a bridge for users as they transition from the 1st generation App Engine
// runtime to the 2nd generation.
// This implementation can be used in the 2nd gen runtime as it can reuse an http.Client.
func NewIAMTokenSource(ctx context.Context, cfg IAMConfig) (oauth2.TokenSource, error) {
	_ = "STUB: not implemented"
	return *new(oauth2.TokenSource), nil
}

// NewContextIAMTokenSource returns an oauth2.TokenSource that uses Google's IAM services
// to sign a JWT with the default service account and the given audience.
// Users should use the Identity token source if they can. This client is meant to be
// used as a bridge for users as they transition from the 1st generation App Engine
// runtime to the 2nd generation.
// This implementation can be used in the 1st gen runtime as it allows users to pass a
// context.Context while fetching the token. The context allows the implementation to
// reuse clients while changing out the HTTP client under the hood.
func NewContextIAMTokenSource(ctx context.Context, cfg IAMConfig) (ContextTokenSource, error) {
	_ = "STUB: not implemented"
	return *new(ContextTokenSource), nil
}

// ContextTokenSource is an oauth2.TokenSource that is capable of running on the 1st
// generation App Engine environment because it can create a urlfetch.Client from the
// given context.
type ContextTokenSource interface {
	ContextToken(context.Context) (*oauth2.Token, error)
}

type iamTokenSource struct {
	cfg IAMConfig

	svc *iam.Service
}

var defaultTokenSource = google.DefaultTokenSource

func (s iamTokenSource) ContextToken(ctx context.Context) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s iamTokenSource) Token() (*oauth2.Token, error) { _ = "STUB: not implemented"; return nil, nil }

func (s iamTokenSource) newIAMToken(ctx context.Context, svc *iam.Service) (string, time.Time, error) {
	_ = "STUB: not implemented"
	return "", *new(time.Time), nil
}

// TAKEN FROM golang.org/x/oauth2 so we can add context bc GAE 1st gen + urlfetch.
// reuseCtxTokenSource is a TokenSource that holds a single token in memory
// and validates its expiry before each call to retrieve it with
// Token. If it's expired, it will be auto-refreshed using the
// new TokenSource.
type reuseTokenSource struct {
	new ContextTokenSource // called when t is expired.

	mu sync.Mutex // guards t
	t  *oauth2.Token
}

// Token returns the current token if it's still valid, else will
// refresh the current token (using r.Context for HTTP client
// information) and return the new one.
func (s *reuseTokenSource) ContextToken(ctx context.Context) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

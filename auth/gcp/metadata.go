package gcp // import "github.com/NYTimes/gizmo/auth/gcp"

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/google"
)

// GetDefaultEmail is a helper method for users on GCE or the 2nd generation GAE
// environment.
func GetDefaultEmail(ctx context.Context, addr string, hc *http.Client) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func metadataGet(ctx context.Context, addr string, hc *http.Client, suffix string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var findDefaultCredentials = google.FindDefaultCredentials

func getEmailFromCredentials(creds *google.Credentials) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

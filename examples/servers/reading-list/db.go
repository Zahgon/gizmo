package readinglist

import (
	"context"

	"cloud.google.com/go/datastore"
)

type DB interface {
	GetLinks(ctx context.Context, userID string, limit int) ([]string, error)
	PutLink(ctx context.Context, userID string, url string) error
	DeleteLink(ctx context.Context, userID string, url string) error
}

type Datastore struct {
	client *datastore.Client
}

const LinkKind = "Link"

func NewDB() (*Datastore, error) { _ = "STUB: not implemented"; return nil, nil }

type linkData struct {
	UserID string
	URL    string `datastore:",noindex"`
}

func newKey(ctx context.Context, userID, url string) *datastore.Key {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datastore) GetLinks(ctx context.Context, userID string, limit int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Datastore) DeleteLink(ctx context.Context, userID string, url string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datastore) PutLink(ctx context.Context, userID string, url string) error {
	_ = "STUB: not implemented"
	return nil
}

// run in transaction to avoid any dupes

// link already exists, just return

// put new link

// Using this to turn keys 12345, 12346 into 54321, 64321 which are easier for
// Datastore/BigTable to shard.
//
// More info: https://cloud.google.com/bigtable/docs/schema-design#row_keys_to_avoid & "Sequential numeric IDs"
func reverse(id string) string { _ = "STUB: not implemented"; return "" }

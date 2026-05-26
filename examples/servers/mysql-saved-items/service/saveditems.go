package service

import (
	"database/sql"
	"time"

	"github.com/NYTimes/gizmo/config/mysql"
	"github.com/NYTimes/sqliface"
)

type (
	// SavedItemsRepo is an interface layer between
	// our service and our database. Abstracting these methods
	// out of a pure implementation helps with testing.
	SavedItemsRepo interface {
		Get(uint64) ([]*SavedItem, error)
		Put(uint64, string) error
		Delete(uint64, string) error
	}

	// MySQLSavedItemsRepo is an implementation of the repo
	// interface built on top of MySQL.
	MySQLSavedItemsRepo struct {
		db *sql.DB
	}

	// SavedItem represents an article, blog, interactive, etc.
	// that a user wants to save for reading later.
	SavedItem struct {
		UserID    uint64    `json:"user_id"`
		URL       string    `json:"url"`
		Timestamp time.Time `json:"timestamp"`
	}
)

// NewSavedItemsRepo will attempt to connect to to MySQL and
// return a SavedItemsRepo implementation.
func NewSavedItemsRepo(cfg *mysql.Config) (SavedItemsRepo, error) {
	_ = "STUB: not implemented"
	return *new(SavedItemsRepo), nil
}

// Get will attempt to query the underlying MySQL database for saved items
// for a single user.
func (r *MySQLSavedItemsRepo) Get(userID uint64) ([]*SavedItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanItems(rows sqliface.Rows) ([]*SavedItem, error) {
	_ = "STUB: not implemented"

	// initializing so we return an empty array in case of 0
	return nil, nil
}

// Put will attempt to insert a new saved item for the user.
func (r *MySQLSavedItemsRepo) Put(userID uint64, url string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete will attempt to remove an item from a user's saved items.
func (r *MySQLSavedItemsRepo) Delete(userID uint64, url string) error {
	_ = "STUB: not implemented"
	return nil
}

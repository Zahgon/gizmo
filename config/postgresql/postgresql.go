package postgresql // import "github.com/NYTimes/gizmo/config/postgresql"

import (
	"database/sql"
)

// Config holds everything you need to
// connect and interact with a PostgreSQL DB.
type Config struct {
	User    string `envconfig:"POSTGRESQL_USER"`
	Pw      string `envconfig:"POSTGRESQL_PW"`
	Host    string `envconfig:"POSTGRESQL_HOST_NAME"`
	Port    int    `envconfig:"POSTGRESQL_PORT"`
	DBName  string `envconfig:"POSTGRESQL_DB_NAME"`
	SSLMode string `envconfig:"POSTGRESQL_SSL_MODE"`
}

const (
	// DefaultSSLMode is verify-full
	DefaultSSLMode = "verify-full"
	// DefaultPort is the default post for Postgresql connections
	DefaultPort = 5432
)

// DB will open a sql connection.
// Users must import a postgresql driver in their
// main to use this.
func (p *Config) DB() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// String will return the Postgresql connection string
func (p *Config) String() string { _ = "STUB: not implemented"; return "" }

// LoadConfigFromEnv will attempt to load a Postgresql object
// from environment variables. If not populated, nil
// is returned
func LoadConfigFromEnv() *Config { _ = "STUB: not implemented"; return nil }

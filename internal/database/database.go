package database

import (
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // The database driver in use.

	"github.com/hoanhan101/progress/internal/config"
)

// Open constructs a database connection string based on a given
// configuration.
func Open(cfg *config.Config, conn string) (*sqlx.DB, error) {
	// If the connection string is provided, use that instead of constructing a
	// new one. One use case is for Heroku deployment.
	if conn != "" {
		return sqlx.Open("postgres", conn)
	}

	// Query parameters.
	q := make(url.Values)
	q.Set("sslmode", cfg.Database.SSLMode)
	q.Set("timezone", "utc")

	// Construct url.
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.Database.User, cfg.Database.Password),
		Host:     cfg.Database.Host,
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise. This forces a round trip to the database.
func StatusCheck(db *sqlx.DB) error {
	const q = `SELECT true`
	var tmp bool
	return db.QueryRow(q).Scan(&tmp)
}

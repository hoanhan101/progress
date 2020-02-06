package database

import (
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/hoanhan101/progress/internal/config"
)

// Open knows how to open a database connection based on the configuration.
func Open(cfg *config.Config) (*sqlx.DB, error) {
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

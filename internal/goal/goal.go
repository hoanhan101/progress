package goal

import (
	// "database/sql"
	// "time"

	// "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// List gets all Goal from the database.
func List(db *sqlx.DB) ([]Goal, error) {
	goals := []Goal{}
	const q = `
SELECT
    *
FROM
    goals;
`

	if err := db.Select(&goals, q); err != nil {
		return nil, errors.Wrap(err, "selecting goals")
	}

	return goals, nil
}

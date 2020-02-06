package database

import (
	"github.com/dimiro1/darwin"
	"github.com/jmoiron/sqlx"
)

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sqlx.DB) error {
	driver := darwin.NewGenericDriver(db.DB, darwin.PostgresDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}

// migrations contains the queries needed to construct the database schema.
// Entries should never be removed from this slice once they have been ran in
// production.
var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Create goals table",
		Script: `
CREATE TABLE goals (
	goal_id      UUID,
	name         TEXT,
	date_created TIMESTAMP,

	PRIMARY KEY (goal_id)
);`,
	},
	{
		Version:     2,
		Description: "Create systems table",
		Script: `
CREATE TABLE systems (
	system_id    UUID,
	goal_id      UUID,
	name         TEXT,
	repeat       TEXT,
	date_created TIMESTAMP,

	PRIMARY KEY (system_id),
	FOREIGN KEY (goal_id) REFERENCES goals(goal_id) ON DELETE CASCADE
);`,
	},
}

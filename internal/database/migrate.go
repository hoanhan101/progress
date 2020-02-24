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
	goal_id      UUID      NOT NULL,
	name         TEXT      NOT NULL,
	context      TEXT      NOT NULL,
	date_created TIMESTAMP NOT NULL,

	PRIMARY KEY (goal_id)
);`,
	},
	{
		Version:     2,
		Description: "Create systems table",
		Script: `
CREATE TABLE systems (
	system_id    UUID      NOT NULL,
	goal_id      UUID      NOT NULL,
	name         TEXT      NOT NULL,
	repeat       TEXT      NOT NULL,
	date_created TIMESTAMP NOT NULL,

	PRIMARY KEY (system_id),
	FOREIGN KEY (goal_id) REFERENCES goals(goal_id) ON DELETE CASCADE
);`,
	},
	{
		Version:     3,
		Description: "Create progress table",
		Script: `
CREATE TABLE progress (
	progress_id     UUID      NOT NULL,
	system_id       UUID      NOT NULL,
	date_created    TIMESTAMP NOT NULL,
	summary         TEXT      NOT NULL,
	completed       BOOLEAN   NOT NULL,
	measurable_data INTEGER   NOT NULL,
	measurable_unit TEXT      NOT NULL,
	sets            INTEGER   NOT NULL,
	reps            INTEGER   NOT NULL,
	link            TEXT      NOT NULL,
	PRIMARY KEY (progress_id),
	FOREIGN KEY (system_id) REFERENCES systems(system_id) ON DELETE CASCADE
);`,
	},
}

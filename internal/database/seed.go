package database

import (
	"github.com/jmoiron/sqlx"
)

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seeds); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
const seeds = `
INSERT INTO goals (goal_id, name, date_created) VALUES
	('a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 'Find my voice and build an audience', '2020-01-01 00:00:01.000001+00'),
	('72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Lift 2x my bodyweight', '2019-01-01 00:00:02.000001+00')
	ON CONFLICT DO NOTHING;

INSERT INTO systems (system_id, goal_id, name, repeat, date_created) VALUES
	('98b6d4b8-f04b-4c79-8c2e-a0aef46854b7', 'a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 'Write every Monday and Thursday', 'Every Mon and Thu', '2020-01-01 00:00:03.000001+00'),
	('a235be9e-ab5d-44e6-a987-fa1c749264c7', '72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Squat 3 times a week', '', '2019-01-01 00:00:05.000001+00'),
	('a335be9e-ab5d-44e6-a987-fa1c749264c7', '72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Bench 3 times a week', '', '2019-01-01 00:00:05.000001+00'),
	('a435be9e-ab5d-44e6-a987-fa1c749264c7', '72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Deadlift 3 times a week', '', '2019-01-01 00:00:05.000001+00')
	ON CONFLICT DO NOTHING;
`

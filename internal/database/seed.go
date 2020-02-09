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
	('a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 'Write consistently to find my voice, build an audience, churn through my average ideas to uncover the great ones', '2020-01-01 00:00:01.000001+00'),
	('72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Lift 2x my bodyweight, look better naked', '2019-01-01 00:00:02.000001+00')
	ON CONFLICT DO NOTHING;

INSERT INTO systems (system_id, goal_id, name, repeat, date_created) VALUES
	('98b6d4b8-f04b-4c79-8c2e-a0aef46854b7', 'a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 'Write every Monday and Thursday', 'Every Mon and Thu', '2020-01-01 00:00:03.000001+00'),
	('a235be9e-ab5d-44e6-a987-fa1c749264c7', '72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Squat 3 times a week', '', '2019-01-01 00:00:05.000001+00'),
	('a335be9e-ab5d-44e6-a987-fa1c749264c7', '72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Bench 3 times a week', '', '2019-01-01 00:00:05.000001+00'),
	('a435be9e-ab5d-44e6-a987-fa1c749264c7', '72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Deadlift 3 times a week', '', '2019-01-01 00:00:05.000001+00')
	ON CONFLICT DO NOTHING;

INSERT INTO progress(progress_id, system_id, date_created, summary, completed, measurable_data, measurable_unit, sets, reps, link) VALUES
	('18b6d4b8-f14b-4c79-8c2e-a0aef46854b7', 'a235be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-02 00:00:05.000001+00', '', 'true', 120, 'lbs', 3, 5, ''),
	('28b6d4b8-f24b-4c79-8c2e-a0aef46854b7', 'a335be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-02 00:00:05.000001+00', '', 'true', 80, 'lbs', 3, 5, ''),
	('38b6d4b8-f34b-4c79-8c2e-a0aef46854b7', 'a435be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-02 00:00:05.000001+00', '', 'true', 140, 'lbs', 1, 5, ''),
	('48b6d4b8-f44b-4c79-8c2e-a0aef46854b7', 'a235be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-04 00:00:05.000001+00', '', 'true', 130, 'lbs', 3, 5, ''),
	('58b6d4b8-f54b-4c79-8c2e-a0aef46854b7', 'a335be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-04 00:00:05.000001+00', '', 'true', 90, 'lbs', 3, 5, ''),
	('68b6d4b8-f64b-4c79-8c2e-a0aef46854b7', 'a435be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-04 00:00:05.000001+00', '', 'true', 150, 'lbs', 1, 5, ''),
	('78b6d4b8-f74b-4c79-8c2e-a0aef46854b7', 'a235be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-06 00:00:05.000001+00', '', 'true', 140, 'lbs', 3, 5, ''),
	('88b6d4b8-f84b-4c79-8c2e-a0aef46854b7', 'a335be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-06 00:00:05.000001+00', '', 'true', 100, 'lbs', 3, 5, ''),
	('98b6d4b8-f94b-4c79-8c2e-a0aef46854b7', 'a435be9e-ab5d-44e6-a987-fa1c749264c7', '2019-01-06 00:00:05.000001+00', '', 'true', 160, 'lbs', 1, 5, ''),
	('11b6d4b8-f01b-4c79-8c1e-a0aef46854b7', '98b6d4b8-f04b-4c79-8c2e-a0aef46854b7', '2019-01-02 00:00:05.000001+00', 'Successful People Start Before They Feel Ready', 'true', 0, '', 1, 1, 'https://jamesclear.com/successful-people-start-before-they-feel-ready'),
	('22b6d4b8-f02b-4c79-8c2e-a0aef46854b7', '98b6d4b8-f04b-4c79-8c2e-a0aef46854b7', '2019-01-04 00:00:05.000001+00', 'How to Be Happy: A Surprising Lesson on Happiness From an African Tribe', 'true', 0, '', 1, 1, 'https://jamesclear.com/how-can-i-be-happy-if-you-are-sad'),
	('33b6d4b8-f05b-4c79-8c3e-a0aef46854b7', '98b6d4b8-f04b-4c79-8c2e-a0aef46854b7', '2019-01-06 00:00:05.000001+00', 'Feeling Fat? Use These 2 Easy Ways to Lose Weight', 'true', 0, '', 1, 1, 'https://jamesclear.com/feeling-fat'),
	('44b6d4b8-f05b-4c79-8c3e-a0aef46854b7', '98b6d4b8-f04b-4c79-8c2e-a0aef46854b7', '2019-01-08 00:00:05.000001+00', 'Believe in Yourself (And Why Nothing Will Work If You Don’t…)', 'true', 0, '', 1, 1, 'https://jamesclear.com/nothing-will-work-if-you-dont-believe-in-it')
	ON CONFLICT DO NOTHING;
`

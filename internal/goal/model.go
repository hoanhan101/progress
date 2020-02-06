package goal

import (
	"time"
)

type Goal struct {
	ID          string    `db:"goal_id" json:"goal_id"`
	Name        string    `db:"name" json:"name"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

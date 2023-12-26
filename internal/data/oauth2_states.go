package data

import "time"

type OAuth2States interface {
	New() OAuth2States

	Get() (*OAuth2State, error)
	Create(oAuth2State OAuth2State) (*OAuth2State, error)
	Delete() error

	FilterByID(id int64) OAuth2States
	FilterByState(state string) OAuth2States
}

type OAuth2State struct {
	ID        int64      `db:"id"`
	State     string     `db:"state"`
	CreatedAt *time.Time `db:"created_at"`
	ValidTill *time.Time `db:"valid_till"`
}

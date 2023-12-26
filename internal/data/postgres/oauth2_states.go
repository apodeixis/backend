package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/apodeixis/backend/internal/data"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const oAuth2StatesTable = "oauth2_states"

var ErrNoSuchState = errors.New("no such state")

type oAuth2StatesQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewOAuth2StatesQ(db *pgdb.DB) data.OAuth2States {
	return &oAuth2StatesQ{
		db:            db.Clone(),
		selectBuilder: sq.Select("*").From(oAuth2StatesTable),
		deleteBuilder: sq.Delete(oAuth2StatesTable),
	}
}

func (q *oAuth2StatesQ) New() data.OAuth2States {
	return NewOAuth2StatesQ(q.db)
}

func (q *oAuth2StatesQ) Create(oAuth2State data.OAuth2State) (*data.OAuth2State, error) {
	clauses := map[string]interface{}{
		"state":      oAuth2State.State,
		"valid_till": oAuth2State.ValidTill,
	}
	result := new(data.OAuth2State)
	stmt := sq.Insert(oAuth2StatesTable).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, &stmt)
	return result, errors.Wrap(err, "failed to create oauth2 state")
}

func (q *oAuth2StatesQ) Get() (*data.OAuth2State, error) {
	result := new(data.OAuth2State)
	err := q.db.Get(result, q.selectBuilder)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return result, errors.Wrap(err, "failed to get oauth2 state")
}

func (q *oAuth2StatesQ) Delete() error {
	result, err := q.db.ExecWithResult(q.deleteBuilder)
	if err != nil {
		return errors.Wrap(err, "failed to delete oauth2 state")
	}
	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return ErrNoSuchState
	}

	return nil
}

func (q *oAuth2StatesQ) FilterByID(id int64) data.OAuth2States {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"id": id})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"id": id})

	return q
}
func (q *oAuth2StatesQ) FilterByState(state string) data.OAuth2States {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"state": state})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"state": state})

	return q
}

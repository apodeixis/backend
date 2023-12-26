package postgres

import (
	"database/sql"

	"github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"
	"github.com/apodeixis/backend/internal/data"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	refreshTokensTokenConstraint = "refresh_tokens_token_uindex"
	refreshTokensTable           = "refresh_tokens"
)

var (
	ErrNoSuchRefreshToken        = errors.New("no such refresh token")
	ErrRefreshTokenAlreadyExists = errors.New("refresh token already exists")
)

type refreshTokensQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
	updateBuilder sq.UpdateBuilder
}

func NewRefreshTokensQ(db *pgdb.DB) data.RefreshTokens {
	return &refreshTokensQ{
		db:            db.Clone(),
		selectBuilder: sq.Select("*").From(refreshTokensTable),
		deleteBuilder: sq.Delete(refreshTokensTable),
		updateBuilder: sq.Update(refreshTokensTable),
	}
}

func (q *refreshTokensQ) New() data.RefreshTokens {
	return NewRefreshTokensQ(q.db)
}

func (q *refreshTokensQ) Create(refreshToken data.RefreshToken) (*data.RefreshToken, error) {
	clauses := map[string]interface{}{
		"user_id":    refreshToken.UserID,
		"token":      refreshToken.Token,
		"valid_till": refreshToken.ValidTill,
	}
	result := new(data.RefreshToken)
	stmt := sq.Insert(refreshTokensTable).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, stmt)

	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Constraint == refreshTokensTokenConstraint {
		return nil, ErrRefreshTokenAlreadyExists
	}

	return result, errors.Wrap(err, "failed to execute insert query")
}

func (q *refreshTokensQ) Update(refreshToken data.RefreshToken) (*data.RefreshToken, error) {
	clauses := map[string]interface{}{
		"user_id":    refreshToken.UserID,
		"token":      refreshToken.Token,
		"valid_till": refreshToken.ValidTill,
	}
	result := new(data.RefreshToken)
	stmt := q.updateBuilder.SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, &stmt)

	return result, errors.Wrap(err, "failed to update refresh token")
}

func (q *refreshTokensQ) Get() (*data.RefreshToken, error) {
	result := new(data.RefreshToken)
	err := q.db.Get(result, q.selectBuilder)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return result, errors.Wrap(err, "failed to get refresh token")
}

func (q *refreshTokensQ) Delete() error {
	result, err := q.db.ExecWithResult(q.deleteBuilder)
	if err != nil {
		return errors.Wrap(err, "failed to execute delete query")
	}
	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return ErrNoSuchRefreshToken
	}

	return nil
}

func (q *refreshTokensQ) FilterByID(id int64) data.RefreshTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"id": id})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"id": id})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"id": id})

	return q
}

func (q *refreshTokensQ) FilterByUserID(userID int64) data.RefreshTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"user_id": userID})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"user_id": userID})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"user_id": userID})

	return q
}

func (q *refreshTokensQ) FilterByToken(token string) data.RefreshTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"token": token})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"token": token})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"token": token})

	return q
}

func (q *refreshTokensQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}

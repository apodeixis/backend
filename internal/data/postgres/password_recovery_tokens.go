package postgres

import (
	"database/sql"

	"github.com/apodeixis/backend/internal/data"

	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const passwordRecoveryTokensTable = "password_recovery_tokens"

var ErrNoSuchPasswordRecoveryToken = errors.New("no such password recovery token")

type passwordRecoveryTokensQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
	updateBuilder sq.UpdateBuilder
}

func NewPasswordRecoveryTokensQ(db *pgdb.DB) data.PasswordRecoveryTokens {
	return &passwordRecoveryTokensQ{
		db:            db.Clone(),
		selectBuilder: sq.Select("*").From(passwordRecoveryTokensTable),
		deleteBuilder: sq.Delete(passwordRecoveryTokensTable),
		updateBuilder: sq.Update(passwordRecoveryTokensTable),
	}
}

func (q *passwordRecoveryTokensQ) New() data.PasswordRecoveryTokens {
	return NewPasswordRecoveryTokensQ(q.db)
}

func (q *passwordRecoveryTokensQ) Create(token data.PasswordRecoveryToken) (*data.PasswordRecoveryToken, error) {
	clauses := map[string]interface{}{
		"user_id":          token.UserID,
		"token":            token.Token,
		"token_expires_at": token.TokenExpiresAt,
	}
	result := new(data.PasswordRecoveryToken)
	stmt := sq.Insert(passwordRecoveryTokensTable).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, stmt)

	return result, errors.Wrap(err, "failed to create password recovery token")
}

func (q *passwordRecoveryTokensQ) Update(request data.PasswordRecoveryToken) (*data.PasswordRecoveryToken, error) {
	clauses := map[string]interface{}{
		"token":            request.Token,
		"token_expires_at": request.TokenExpiresAt,
	}
	result := new(data.PasswordRecoveryToken)
	stmt := q.updateBuilder.SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, &stmt)

	return result, errors.Wrap(err, "failed to update password recovery token")
}

func (q *passwordRecoveryTokensQ) Get() (*data.PasswordRecoveryToken, error) {
	result := new(data.PasswordRecoveryToken)
	err := q.db.Get(result, q.selectBuilder)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return result, errors.Wrap(err, "failed to get password recovery token")
}

func (q *passwordRecoveryTokensQ) Delete() error {
	result, err := q.db.ExecWithResult(q.deleteBuilder)
	if err != nil {
		return errors.Wrap(err, "failed to delete password recovery token")
	}
	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return ErrNoSuchPasswordRecoveryToken
	}

	return nil
}

func (q *passwordRecoveryTokensQ) FilterByID(requestID int64) data.PasswordRecoveryTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"id": requestID})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"id": requestID})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"id": requestID})

	return q
}

func (q *passwordRecoveryTokensQ) FilterByUserID(userID int64) data.PasswordRecoveryTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"user_id": userID})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"user_id": userID})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"user_id": userID})

	return q
}

func (q *passwordRecoveryTokensQ) FilterByToken(token string) data.PasswordRecoveryTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"token": token})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"token": token})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"token": token})

	return q
}

func (q *passwordRecoveryTokensQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}

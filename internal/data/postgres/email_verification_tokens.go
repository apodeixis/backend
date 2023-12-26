package postgres

import (
	"database/sql"

	"github.com/apodeixis/backend/internal/data"

	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const emailVerificationTokensTable = "email_verification_tokens"

var ErrNoSuchEmailVerificationToken = errors.New("no such email verification token")

type emailVerificationTokensQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewEmailVerificationTokensQ(db *pgdb.DB) data.EmailVerificationTokens {
	return &emailVerificationTokensQ{
		db:            db.Clone(),
		selectBuilder: sq.Select("*").From(emailVerificationTokensTable),
		deleteBuilder: sq.Delete(emailVerificationTokensTable),
	}
}

func (q *emailVerificationTokensQ) New() data.EmailVerificationTokens {
	return NewEmailVerificationTokensQ(q.db)
}

func (q *emailVerificationTokensQ) Create(token data.EmailVerificationToken) (*data.EmailVerificationToken, error) {
	clauses := map[string]interface{}{
		"user_id":          token.UserID,
		"token":            token.Token,
		"token_expires_at": token.TokenExpiresAt,
	}
	result := new(data.EmailVerificationToken)
	stmt := sq.Insert(emailVerificationTokensTable).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, stmt)

	return result, errors.Wrap(err, "failed to create email verification token")
}

func (q *emailVerificationTokensQ) Get() (*data.EmailVerificationToken, error) {
	result := new(data.EmailVerificationToken)
	err := q.db.Get(result, q.selectBuilder)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return result, errors.Wrap(err, "failed to get email verification token")
}

func (q *emailVerificationTokensQ) Delete() error {
	result, err := q.db.ExecWithResult(q.deleteBuilder)
	if err != nil {
		return errors.Wrap(err, "failed to delete email verification token")
	}
	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return ErrNoSuchEmailVerificationToken
	}

	return nil
}

func (q *emailVerificationTokensQ) FilterByID(requestID int64) data.EmailVerificationTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"id": requestID})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"id": requestID})

	return q
}

func (q *emailVerificationTokensQ) FilterByUserID(userID int64) data.EmailVerificationTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"user_id": userID})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"user_id": userID})

	return q
}

func (q *emailVerificationTokensQ) FilterByToken(token string) data.EmailVerificationTokens {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"token": token})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"token": token})

	return q
}

func (q *emailVerificationTokensQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}

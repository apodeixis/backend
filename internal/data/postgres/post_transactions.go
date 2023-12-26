package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/apodeixis/backend/internal/data"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	postTransactionsTable = "post_transactions"
)

var ErrNoSuchPostTransaction = errors.New("no such post transaction")

type postTransactionsQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewPostTransactionsQ(db *pgdb.DB) data.PostTransactions {
	return &postTransactionsQ{
		db:            db.Clone(),
		selectBuilder: sq.Select("*").From(postTransactionsTable),
		deleteBuilder: sq.Delete(postTransactionsTable),
	}
}

func (q *postTransactionsQ) New() data.PostTransactions {
	return NewPostTransactionsQ(q.db)
}

func (q *postTransactionsQ) Create(tx data.PostTransaction) (*data.PostTransaction, error) {
	clauses := map[string]interface{}{
		"tx":      tx.Tx,
		"post_id": tx.PostID,
	}
	result := new(data.PostTransaction)
	query := sq.Insert(postTransactionsTable).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, query)
	return result, errors.Wrap(err, "failed to create post transaction")
}

func (q *postTransactionsQ) Get() (*data.PostTransaction, error) {
	result := new(data.PostTransaction)
	err := q.db.Get(result, q.selectBuilder)
	if errors.Cause(err) == sql.ErrNoRows {
		return nil, nil
	}
	return result, errors.Wrap(err, "failed to get post transaction")
}

func (q *postTransactionsQ) Delete() error {
	result, err := q.db.ExecWithResult(q.deleteBuilder)
	if err != nil {
		return errors.Wrap(err, "failed to execute delete query")
	}
	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return ErrNoSuchPostTransaction
	}

	return nil
}

func (q *postTransactionsQ) FilterByPostID(postID int64) data.PostTransactions {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"post_id": postID})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"post_id": postID})

	return q
}

func (q *postTransactionsQ) Limit(limit uint64) data.PostTransactions {
	q.selectBuilder = q.selectBuilder.Limit(limit)
	q.deleteBuilder = q.deleteBuilder.Limit(limit)

	return q
}

func (q *postTransactionsQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}

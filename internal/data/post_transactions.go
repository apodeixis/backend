package data

import (
	"github.com/apodeixis/backend/internal/types"
)

type PostTransactions interface {
	New() PostTransactions

	Create(tx PostTransaction) (*PostTransaction, error)
	Get() (*PostTransaction, error)
	Delete() error

	FilterByPostID(postID int64) PostTransactions
	Limit(limit uint64) PostTransactions

	Transaction(func() error) error
}

type PostTransaction struct {
	ID     int64                 `db:"id"`
	PostID int64                 `db:"post_id"`
	Tx     *types.RLPTransaction `db:"tx"`
}

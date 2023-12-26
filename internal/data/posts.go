package data

import (
	"time"

	"github.com/apodeixis/backend/internal/types"
)

type Posts interface {
	New() Posts

	Create(post Post) (*Post, error)
	Update(post Post) (*Post, error)
	Select() ([]Post, error)
	SelectHeaders() ([]PostHeader, error)
	Get() (*Post, error)
	Count() (int64, error)

	FilterByID(id int64) Posts
	FilterByUserID(userID int64) Posts
	FilterByStatus(status types.PostStatus) Posts

	WhereIDGreaterThan(id int64) Posts
	WhereIDLessThan(id int64) Posts

	JoinStarredPostsOnPostID() Posts
	FilterByStarredPostsUserID(userID int64) Posts

	JoinUsersOnID() Posts
	FilterByUsersAuthorID(authorID int64) Posts

	OrderByID(sorting types.Sorting) Posts
	Limit(limit uint64) Posts
	Offset(offset uint64) Posts

	Transaction(func() error) error
}

type Post struct {
	ID          int64            `db:"id"`
	UserID      int64            `db:"user_id"`
	Title       string           `db:"title"`
	Body        string           `db:"body"`
	Status      types.PostStatus `db:"status"`
	TxHash      *string          `db:"tx_hash"`
	TxTimestamp *time.Time       `db:"tx_timestamp"`
}

type PostHeader struct {
	ID          int64            `db:"id"`
	UserID      int64            `db:"user_id"`
	Title       string           `db:"title"`
	Status      types.PostStatus `db:"status"`
	TxHash      *string          `db:"tx_hash"`
	TxTimestamp *time.Time       `db:"tx_timestamp"`
}

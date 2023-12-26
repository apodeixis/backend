package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/apodeixis/backend/internal/types"

	sq "github.com/Masterminds/squirrel"
	"github.com/apodeixis/backend/internal/data"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	postsTable = "posts"
)

var headerFields = []string{"posts.id", "posts.user_id", "title", "status", "tx_hash", "tx_timestamp"}

type postsQ struct {
	db                   *pgdb.DB
	selectBuilder        sq.SelectBuilder
	headersSelectBuilder sq.SelectBuilder
	countBuilder         sq.SelectBuilder
	updateBuilder        sq.UpdateBuilder
}

func NewPostsQ(db *pgdb.DB) data.Posts {
	return &postsQ{
		db:                   db.Clone(),
		selectBuilder:        sq.Select("DISTINCT posts.*").From(postsTable),
		headersSelectBuilder: sq.Select("DISTINCT " + strings.Join(headerFields, ",")).From(postsTable),
		countBuilder:         sq.Select("Count(id)").From(postsTable),
		updateBuilder:        sq.Update(postsTable),
	}
}

func (q *postsQ) New() data.Posts {
	return NewPostsQ(q.db)
}

func (q *postsQ) Create(post data.Post) (*data.Post, error) {
	clauses := map[string]interface{}{
		"user_id": post.UserID,
		"title":   post.Title,
		"body":    post.Body,
		"status":  types.NewPostStatus,
	}
	result := new(data.Post)
	query := sq.Insert(postsTable).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, query)
	return result, errors.Wrap(err, "failed to create post")
}

func (q *postsQ) Update(post data.Post) (*data.Post, error) {
	clauses := map[string]interface{}{
		"status":       post.Status,
		"tx_hash":      post.TxHash,
		"tx_timestamp": post.TxTimestamp,
	}
	result := new(data.Post)
	query := q.updateBuilder.SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, query)
	return result, errors.Wrap(err, "failed to update post")
}

func (q *postsQ) Select() ([]data.Post, error) {
	posts := make([]data.Post, 0)
	if err := q.db.Select(&posts, q.selectBuilder); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "failed to select posts")
		}
	}
	return posts, nil
}

func (q *postsQ) SelectHeaders() ([]data.PostHeader, error) {
	postsHeaders := make([]data.PostHeader, 0)
	if err := q.db.Select(&postsHeaders, q.headersSelectBuilder); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "failed to select posts headers")
		}
	}
	return postsHeaders, nil
}

func (q *postsQ) Count() (int64, error) {
	result := make([]int64, 0)
	if err := q.db.Select(&result, q.countBuilder); err != nil {
		if err != sql.ErrNoRows {
			return -1, errors.Wrap(err, "failed to count posts")
		}
	}
	return result[0], nil
}

func (q *postsQ) Get() (*data.Post, error) {
	result := new(data.Post)
	err := q.db.Get(result, q.selectBuilder)
	if errors.Cause(err) == sql.ErrNoRows {
		return nil, nil
	}
	return result, errors.Wrap(err, "failed to get post")
}

func (q *postsQ) FilterByID(id int64) data.Posts {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"id": id})
	q.headersSelectBuilder = q.headersSelectBuilder.Where(sq.Eq{"id": id})
	q.countBuilder = q.countBuilder.Where(sq.Eq{"id": id})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"id": id})

	return q
}

func (q *postsQ) FilterByUserID(userID int64) data.Posts {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"user_id": userID})
	q.headersSelectBuilder = q.headersSelectBuilder.Where(sq.Eq{"user_id": userID})
	q.countBuilder = q.countBuilder.Where(sq.Eq{"user_id": userID})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"user_id": userID})

	return q
}

func (q *postsQ) FilterByAuthorID(authorID int64) data.Posts {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"author_id": authorID})
	q.headersSelectBuilder = q.headersSelectBuilder.Where(sq.Eq{"author_id": authorID})
	q.countBuilder = q.countBuilder.Where(sq.Eq{"author_id": authorID})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"author_id": authorID})

	return q
}

func (q *postsQ) FilterByStatus(status types.PostStatus) data.Posts {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"status": status})
	q.headersSelectBuilder = q.headersSelectBuilder.Where(sq.Eq{"status": status})
	q.countBuilder = q.countBuilder.Where(sq.Eq{"status": status})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"status": status})

	return q
}

func (q *postsQ) OrderByID(sorting types.Sorting) data.Posts {
	clause := fmt.Sprintf("%s.id %s", postsTable, sorting)

	q.selectBuilder = q.selectBuilder.OrderBy(clause)
	q.headersSelectBuilder = q.headersSelectBuilder.OrderBy(clause)
	q.countBuilder = q.countBuilder.OrderBy(clause)
	q.updateBuilder = q.updateBuilder.OrderBy(clause)

	return q
}

func (q *postsQ) WhereIDGreaterThan(id int64) data.Posts {
	clause := sq.Gt{fmt.Sprintf("%s.id", postsTable): id}
	q.selectBuilder = q.selectBuilder.Where(clause)
	q.headersSelectBuilder = q.headersSelectBuilder.Where(clause)
	q.countBuilder = q.countBuilder.Where(clause)
	q.updateBuilder = q.updateBuilder.Where(clause)

	return q
}

func (q *postsQ) WhereIDLessThan(id int64) data.Posts {
	clause := sq.Lt{fmt.Sprintf("%s.id", postsTable): id}
	q.selectBuilder = q.selectBuilder.Where(clause)
	q.headersSelectBuilder = q.headersSelectBuilder.Where(clause)
	q.countBuilder = q.countBuilder.Where(clause)
	q.updateBuilder = q.updateBuilder.Where(clause)

	return q
}

func (q *postsQ) JoinStarredPostsOnPostID() data.Posts {
	join := fmt.Sprintf("%s ON %s.post_id = %s.id ",
		starredPostsTable, starredPostsTable, postsTable)
	q.selectBuilder = q.selectBuilder.Join(join)
	q.headersSelectBuilder = q.headersSelectBuilder.Join(join)
	q.countBuilder = q.countBuilder.Join(join)
	q.updateBuilder = q.updateBuilder.Where(join)

	return q
}

func (q *postsQ) FilterByStarredPostsUserID(userID int64) data.Posts {
	pred := sq.Eq{fmt.Sprintf("%s.user_id", starredPostsTable): userID}
	q.selectBuilder = q.selectBuilder.Where(pred)
	q.headersSelectBuilder = q.headersSelectBuilder.Where(pred)
	q.countBuilder = q.countBuilder.Where(pred)
	q.updateBuilder = q.updateBuilder.Where(pred)

	return q
}

func (q *postsQ) JoinUsersOnID() data.Posts {
	join := fmt.Sprintf("%s ON %s.user_id = %s.id ",
		usersTable, postsTable, usersTable)
	q.selectBuilder = q.selectBuilder.Join(join)
	q.headersSelectBuilder = q.headersSelectBuilder.Join(join)
	q.countBuilder = q.countBuilder.Join(join)
	q.updateBuilder = q.updateBuilder.Where(join)

	return q
}

func (q *postsQ) FilterByUsersAuthorID(authorID int64) data.Posts {
	pred := sq.Eq{fmt.Sprintf("%s.author_id", usersTable): authorID}
	q.selectBuilder = q.selectBuilder.Where(pred)
	q.headersSelectBuilder = q.headersSelectBuilder.Where(pred)
	q.countBuilder = q.countBuilder.Where(pred)
	q.updateBuilder = q.updateBuilder.Where(pred)

	return q
}

func (q *postsQ) Limit(limit uint64) data.Posts {
	q.selectBuilder = q.selectBuilder.Limit(limit)
	q.headersSelectBuilder = q.headersSelectBuilder.Limit(limit)
	q.countBuilder = q.countBuilder.Limit(limit)
	q.updateBuilder = q.updateBuilder.Limit(limit)

	return q
}

func (q *postsQ) Offset(offset uint64) data.Posts {
	q.selectBuilder = q.selectBuilder.Offset(offset)
	q.headersSelectBuilder = q.headersSelectBuilder.Offset(offset)
	q.countBuilder = q.countBuilder.Offset(offset)
	q.updateBuilder = q.updateBuilder.Offset(offset)

	return q
}

func (q *postsQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}

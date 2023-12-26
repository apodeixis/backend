package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/apodeixis/backend/internal/data"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	starredPostsTable                    = "starred_posts"
	starredPostsUserIDFkeyConstraint     = "starred_posts_user_id_fkey"
	starredPostsPostIDFkeyConstraint     = "starred_posts_post_id_fkey"
	starredPostsUserPostUniqueConstraint = "starred_posts_user_post_uindex"
)

var (
	ErrNoSuchStarredPost        = errors.New("no such starred post")
	ErrNoSuchPost               = errors.New("no such post")
	ErrNoSuchUser               = errors.New("no such user")
	ErrStarredPostAlreadyExists = errors.New("starred post already exists")
)

type starredPostsQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewStarredPostsQ(db *pgdb.DB) data.StarredPosts {
	return &starredPostsQ{
		db:            db.Clone(),
		selectBuilder: sq.Select("*").From(starredPostsTable),
		deleteBuilder: sq.Delete(starredPostsTable),
	}
}

func (q *starredPostsQ) New() data.StarredPosts {
	return NewStarredPostsQ(q.db)
}

func (q *starredPostsQ) Create(starredPost data.StarredPost) (*data.StarredPost, error) {
	clauses := map[string]interface{}{
		"user_id": starredPost.UserID,
		"post_id": starredPost.PostID,
	}
	result := new(data.StarredPost)
	query := sq.Insert(starredPostsTable).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, query)
	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Constraint == starredPostsPostIDFkeyConstraint {
		return nil, ErrNoSuchPost
	}
	if errors.As(err, &pgErr) && pgErr.Constraint == starredPostsUserIDFkeyConstraint {
		return nil, ErrNoSuchUser
	}
	if errors.As(err, &pgErr) && pgErr.Constraint == starredPostsUserPostUniqueConstraint {
		return nil, ErrStarredPostAlreadyExists
	}
	return result, errors.Wrap(err, "failed to create starred post")
}

func (q *starredPostsQ) Delete() error {
	result, err := q.db.ExecWithResult(q.deleteBuilder)
	if err != nil {
		return errors.Wrap(err, "failed to delete starred post")
	}
	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return ErrNoSuchStarredPost
	}
	return nil
}

func (q *starredPostsQ) Get() (*data.StarredPost, error) {
	result := new(data.StarredPost)
	err := q.db.Get(result, q.selectBuilder)
	if errors.Cause(err) == sql.ErrNoRows {
		return nil, nil
	}
	return result, errors.Wrap(err, "failed to get starred post")
}

func (q *starredPostsQ) FilterByUserID(userID int64) data.StarredPosts {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"user_id": userID})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"user_id": userID})

	return q
}

func (q *starredPostsQ) FilterByPostID(postID int64) data.StarredPosts {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"post_id": postID})
	q.deleteBuilder = q.deleteBuilder.Where(sq.Eq{"post_id": postID})

	return q
}

func (q *starredPostsQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}

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
	usersTable              = "users"
	usersAuthorIDConstraint = "users_author_id_uindex"
	usersEmailConstraint    = "users_email_uindex"
)

var (
	ErrUserWithSuchAuthorIDAlreadyExists = errors.New("user with such author id already exists")
	ErrUserWithSuchEmailAlreadyExists    = errors.New("user with such email already exists")
)

type usersQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	updateBuilder sq.UpdateBuilder
}

func NewUsersQ(db *pgdb.DB) data.Users {
	return &usersQ{
		db:            db.Clone(),
		selectBuilder: sq.Select("*").From(usersTable),
		updateBuilder: sq.Update(usersTable),
	}
}

func (q *usersQ) New() data.Users {
	return NewUsersQ(q.db)
}

func (q *usersQ) CreateUserFromSignUp(user data.User) (*data.User, error) {
	clauses := map[string]interface{}{
		"author_id":   user.AuthorID,
		"email":       user.Email,
		"password":    user.Password,
		"name":        user.Name,
		"oauth2_user": false,
	}

	return q.createUserFromClauses(clauses)
}

func (q *usersQ) CreateUserFromOAuth2(user data.User) (*data.User, error) {
	clauses := map[string]interface{}{
		"author_id":       user.AuthorID,
		"email":           user.Email,
		"password":        user.Password,
		"name":            user.Name,
		"oauth2_user":     true,
		"oauth2_provider": user.OAuth2Provider,
	}

	return q.createUserFromClauses(clauses)
}

func (q *usersQ) createUserFromClauses(clauses map[string]interface{}) (*data.User, error) {
	result := new(data.User)
	stmt := sq.Insert(usersTable).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, &stmt)

	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Constraint == usersEmailConstraint {
		return nil, ErrUserWithSuchEmailAlreadyExists
	}
	if errors.As(err, &pgErr) && pgErr.Constraint == usersAuthorIDConstraint {
		return nil, ErrUserWithSuchAuthorIDAlreadyExists
	}

	return result, errors.Wrap(err, "failed to create user")
}

func (q *usersQ) Update(user data.User) (*data.User, error) {
	clauses := map[string]interface{}{
		"name":           user.Name,
		"password":       user.Password,
		"email_verified": user.EmailVerified,
	}
	result := new(data.User)
	query := q.updateBuilder.SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, query)
	return result, errors.Wrap(err, "failed to update user")
}

func (q *usersQ) Get() (*data.User, error) {
	result := new(data.User)
	err := q.db.Get(result, q.selectBuilder)
	if errors.Cause(err) == sql.ErrNoRows {
		return nil, nil
	}
	return result, errors.Wrap(err, "failed to get user")
}

func (q *usersQ) FilterByID(id int64) data.Users {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"id": id})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"id": id})

	return q
}

func (q *usersQ) FilterByEmail(email string) data.Users {
	q.selectBuilder = q.selectBuilder.Where(sq.Eq{"email": email})
	q.updateBuilder = q.updateBuilder.Where(sq.Eq{"email": email})

	return q
}

func (q *usersQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}

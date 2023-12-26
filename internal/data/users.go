package data

import (
	"time"

	"github.com/apodeixis/backend/internal/types"
)

type Users interface {
	New() Users

	CreateUserFromSignUp(user User) (*User, error)
	CreateUserFromOAuth2(user User) (*User, error)

	Update(user User) (*User, error)
	Get() (*User, error)

	FilterByID(id int64) Users
	FilterByEmail(email string) Users

	Transaction(func() error) error
}

type User struct {
	ID             int64                 `db:"id"`
	AuthorID       int64                 `db:"author_id"`
	Email          string                `db:"email"`
	Password       *string               `db:"password"`
	Name           string                `db:"name"`
	EmailVerified  bool                  `db:"email_verified"`
	OAuth2User     bool                  `db:"oauth2_user"`
	OAuth2Provider *types.OAuth2Provider `db:"oauth2_provider"`
	CreatedAt      *time.Time            `db:"created_at"`
	UpdatedAt      *time.Time            `db:"updated_at"`
}

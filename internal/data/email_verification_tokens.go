package data

import "time"

type EmailVerificationTokens interface {
	New() EmailVerificationTokens

	Create(token EmailVerificationToken) (*EmailVerificationToken, error)
	Get() (*EmailVerificationToken, error)
	Delete() error

	FilterByID(id int64) EmailVerificationTokens
	FilterByUserID(id int64) EmailVerificationTokens
	FilterByToken(token string) EmailVerificationTokens

	Transaction(func() error) error
}

type EmailVerificationToken struct {
	ID             int64      `db:"id"`
	UserID         int64      `db:"user_id"`
	Token          string     `db:"token"`
	CreatedAt      *time.Time `db:"created_at"`
	TokenExpiresAt *time.Time `db:"token_expires_at"`
}

package data

import "time"

type PasswordRecoveryTokens interface {
	New() PasswordRecoveryTokens

	Create(token PasswordRecoveryToken) (*PasswordRecoveryToken, error)
	Update(token PasswordRecoveryToken) (*PasswordRecoveryToken, error)
	Get() (*PasswordRecoveryToken, error)
	Delete() error

	FilterByID(id int64) PasswordRecoveryTokens
	FilterByUserID(id int64) PasswordRecoveryTokens
	FilterByToken(token string) PasswordRecoveryTokens

	Transaction(func() error) error
}

type PasswordRecoveryToken struct {
	ID             int64      `db:"id"`
	UserID         int64      `db:"user_id"`
	Token          string     `db:"token"`
	CreatedAt      *time.Time `db:"created_at"`
	TokenExpiresAt *time.Time `db:"token_expires_at"`
}

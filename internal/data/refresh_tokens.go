package data

import "time"

type RefreshTokens interface {
	New() RefreshTokens

	Create(refreshToken RefreshToken) (*RefreshToken, error)
	Update(refreshToken RefreshToken) (*RefreshToken, error)
	Get() (*RefreshToken, error)
	Delete() error

	FilterByID(tokenID int64) RefreshTokens
	FilterByUserID(userID int64) RefreshTokens
	FilterByToken(token string) RefreshTokens

	Transaction(func() error) error
}

type RefreshToken struct {
	ID        int64      `db:"id"`
	UserID    int64      `db:"user_id"`
	Token     string     `db:"token"`
	CreatedAt *time.Time `db:"created_at"`
	ValidTill time.Time  `db:"valid_till"`
}

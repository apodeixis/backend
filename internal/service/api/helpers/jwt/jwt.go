package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

const TokenPattern = "(^[A-Za-z0-9-_]*\\.[A-Za-z0-9-_]*\\.[A-Za-z0-9-_]*$)"

type Claims struct {
	OwnerId   int64 `json:"owner_id"`
	ExpiresAt int64 `json:"expires_at"`
}

func (c *Claims) Valid() error {
	if c.ExpiresAt < time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func CreateToken(claims *Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	return signedToken, errors.Wrap(err, "failed to sign token")
}

func ExtractClaims(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}
	return claims, nil
}

package helpers

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
)

func GenerateToken(secret string) (string, error) {
	// Generate a 32-byte random token
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	tokenString := base64.URLEncoding.EncodeToString(token)

	// Append the current Unix time to the token to create a unique ID
	tokenString += fmt.Sprintf("%d", time.Now().Unix())

	// Generate a digital signature for the token using HMAC-SHA256
	hmacHash := hmac.New(sha256.New, []byte(secret))
	hmacHash.Write([]byte(tokenString))
	signature := hmacHash.Sum(nil)

	// Append the signature to the end of the token string
	tokenString += base64.URLEncoding.EncodeToString(signature)

	return tokenString, nil
}

func VerifyToken(token, secret string) error {
	// Extract the signature from the end of the token string
	signatureString := token[len(token)-44:]
	signature, err := base64.URLEncoding.DecodeString(signatureString)
	if err != nil {
		return err
	}

	// Verify the signature using HMAC-SHA256
	tokenString := token[:len(token)-44]
	hmacHash := hmac.New(sha256.New, []byte(secret))
	hmacHash.Write([]byte(tokenString))
	expectedSignature := hmacHash.Sum(nil)

	if !hmac.Equal(expectedSignature, signature) {
		return errors.New("invalid signature")
	}
	return nil
}

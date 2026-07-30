package authservice

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
}

func (c Claims) Valid() error {
	return nil
}

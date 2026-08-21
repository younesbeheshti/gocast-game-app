package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/younesbeheshti/gocast_game/entity"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint        `json:"user_id"`
	Role   entity.Role `json:"role"`
}

func (c Claims) Valid() error {
	return nil
}

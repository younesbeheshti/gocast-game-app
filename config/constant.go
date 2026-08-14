package config

import "time"

const (
	JwtSecret                  = "secret"
	AccessTokenSubject         = "access_token"
	RefreshTokenSubject        = "refresh_token"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
	AuthMiddlewareContextKey   = "claims"
)

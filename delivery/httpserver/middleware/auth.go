package middleware

import (
	"github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	cfg "github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/service/authservice"
)

func Auth(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ContextKey:    cfg.AuthMiddlewareContextKey,
		SigningKey:    []byte(config.SignKey),
		SigningMethod: "HS256",
		ParseTokenFunc: func(c *echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claims, nil
		}})
}

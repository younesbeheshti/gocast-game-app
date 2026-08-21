package claim

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/service/authservice"
)

func GetClaimsFromEchoContext(c *echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}

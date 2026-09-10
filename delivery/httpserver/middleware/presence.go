package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/claim"
	"github.com/younesbeheshti/gocast_game/pkg/timestamp"
	"github.com/younesbeheshti/gocast_game/service/presenceservice"
	"net/http"
)

func UpsertPresence(service presenceservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			claims := claim.GetClaimsFromEchoContext(c)

			_, err := service.Upsert(c.Request().Context(), &param.UpsertPresenceRequest{
				UserID:    claims.UserID,
				Timestamp: timestamp.Now(),
			})

			if err != nil {

				// TODO : log unexpected error
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "something went wrong",
				})
			}

			return next(c)

		}

	}
}

package middleware

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/claim"
	"github.com/younesbeheshti/gocast_game/service/authorizationservice"
	"net/http"
)

func AccessCheck(service authorizationservice.Service, permissions ...entity.PermissionTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			claims := claim.GetClaimsFromEchoContext(c)
			isAllowed, err := service.CheckAccess(claims.UserID, claims.Role, permissions...)
			if err != nil {

				// TODO : log unexpected error

				fmt.Println("err : ", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "something went wrong",
				})
			}

			if !isAllowed {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": "user not allowed",
				})
			}

			return next(c)

		}

	}
}

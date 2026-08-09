package userhandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/pkg/httpmsg"
	"net/http"
)

func (h Handler) userProfileHandler(c *echo.Context) error {

	authToken := c.Request().Header.Get("Authorization")
	claims, err := h.authSvc.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := h.userSvc.GetProfile(dto.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)

}

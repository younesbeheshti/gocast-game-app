package userhandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/httpmsg"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"net/http"
)

func getClaims(c *echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}

func (h Handler) userProfileHandler(c *echo.Context) error {

	claims := getClaims(c)
	resp, err := h.userSvc.GetProfile(param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)

}

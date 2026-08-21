package userhandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/claim"
	"github.com/younesbeheshti/gocast_game/pkg/httpmsg"
	"net/http"
)

func (h Handler) userProfileHandler(c *echo.Context) error {

	claims := claim.GetClaimsFromEchoContext(c)
	resp, err := h.userSvc.GetProfile(param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)

}

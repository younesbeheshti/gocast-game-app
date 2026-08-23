package matchinghandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/claim"
	"github.com/younesbeheshti/gocast_game/pkg/httpmsg"
	"net/http"
)

func (h Handler) addToWaitingList(c *echo.Context) error {

	var req param.AddToWaitingListRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	claims := claim.GetClaimsFromEchoContext(c)
	req.UserID = claims.UserID

	if fieldErrors, err := h.matchingValidator.ValidateAddToWaitingListRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, map[string]interface{}{
			"message": msg,
			"error":   fieldErrors,
		})
	}

	resp, err := h.matchingSvc.AddToWaitingList(&req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}

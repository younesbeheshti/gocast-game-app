package backofficeuserhandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/pkg/httpmsg"
	"net/http"
)

func (h Handler) listUserHandler(c *echo.Context) error {
	list, err := h.backofficeUserSvc.DoSomething()
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
	})
}

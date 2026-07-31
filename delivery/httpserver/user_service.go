package httpserver

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"net/http"
)

func (s Server) userRegister(c *echo.Context) error {

	var req userservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

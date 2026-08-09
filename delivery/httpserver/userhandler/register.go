package userhandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/pkg/httpmsg"
	"net/http"
)

func (h Handler) userRegisterHandler(c *echo.Context) error {

	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if fieldErrors, err := h.userValidator.ValidateRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, map[string]interface{}{
			"message": msg,
			"error":   fieldErrors,
		})
	}

	resp, err := h.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

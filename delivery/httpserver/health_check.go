package httpserver

import (
	"github.com/labstack/echo/v5"
	"net/http"
)

func (Server) healthCheck(c *echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

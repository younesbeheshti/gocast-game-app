package matchinghandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/middleware"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {

	usersGroup := e.Group("/matching")

	usersGroup.POST("/add-to-waiting-list", h.addToWaitingList, middleware.Auth(h.authSvc, h.authConfig), middleware.UpsertPresence(h.presenceSvc))

}

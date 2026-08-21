package backofficeuserhandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/middleware"
	"github.com/younesbeheshti/gocast_game/entity"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {

	usersGroup := e.Group("/backoffice/users")

	usersGroup.GET("/", h.listUserHandler, middleware.Auth(h.authSvc, h.authConfig), middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))
}

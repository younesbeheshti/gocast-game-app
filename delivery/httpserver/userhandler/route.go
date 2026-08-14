package userhandler

import (
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/middleware"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {

	usersGroup := e.Group("/users")
	usersGroup.POST("/register", h.userRegisterHandler)
	usersGroup.POST("/login", h.userLoginHandler)
	usersGroup.GET("/userprofile", h.userProfileHandler, middleware.Auth(h.authSvc, h.authConfig))
}

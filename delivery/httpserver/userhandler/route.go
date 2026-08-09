package userhandler

import "github.com/labstack/echo/v5"

func (h Handler) SetUserRoutes(e *echo.Echo) {

	usersGroup := e.Group("/users")
	usersGroup.POST("/register", h.userRegisterHandler)
	usersGroup.POST("/login", h.userLoginHandler)
	usersGroup.GET("/userprofile", h.userProfileHandler)
}

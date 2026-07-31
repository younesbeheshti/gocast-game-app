package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"log/slog"
)

type Server struct {
	config  config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service) *Server {
	return &Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/health_check", s.healthCheck)

	usersGroup := e.Group("/users")
	usersGroup.POST("/register", s.userRegisterHandler)
	usersGroup.POST("/login", s.userLoginHandler)
	usersGroup.GET("/userprofile", s.userProfileHandler)

	if err := e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}

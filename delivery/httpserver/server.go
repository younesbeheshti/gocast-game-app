package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/backofficeuserhandler"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/matchinghandler"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/userhandler"
	"github.com/younesbeheshti/gocast_game/service/authorizationservice"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/backofficeuserservice"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"github.com/younesbeheshti/gocast_game/validator/matchingvalidator"
	"github.com/younesbeheshti/gocast_game/validator/uservalidator"
	"log/slog"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, validator uservalidator.Validator, backofficeUserSvc backofficeuserservice.Service, authorizationSvc authorizationservice.Service, matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator) *Server {
	return &Server{
		config:                config,
		userHandler:           userhandler.New(authSvc, userSvc, validator, config.Auth),
		backofficeUserHandler: backofficeuserhandler.New(authSvc, config.Auth, backofficeUserSvc, authorizationSvc),
		matchingHandler:       matchinghandler.New(authSvc, config.Auth, matchingSvc, matchingValidator),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/health_check", s.healthCheck)

	s.userHandler.SetUserRoutes(e)
	s.backofficeUserHandler.SetUserRoutes(e)
	s.matchingHandler.SetUserRoutes(e)

	if err := e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}

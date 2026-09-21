package httpserver

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/backofficeuserhandler"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/matchinghandler"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver/userhandler"
	"github.com/younesbeheshti/gocast_game/logger"
	"github.com/younesbeheshti/gocast_game/service/authorizationservice"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/backofficeuserservice"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
	"github.com/younesbeheshti/gocast_game/service/presenceservice"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"github.com/younesbeheshti/gocast_game/validator/matchingvalidator"
	"github.com/younesbeheshti/gocast_game/validator/uservalidator"
	"go.uber.org/zap"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, validator uservalidator.Validator, backofficeUserSvc backofficeuserservice.Service, authorizationSvc authorizationservice.Service, matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator, presenceSvc presenceservice.Service) *Server {
	return &Server{
		config:                config,
		userHandler:           userhandler.New(authSvc, userSvc, validator, config.Auth, presenceSvc),
		backofficeUserHandler: backofficeuserhandler.New(authSvc, config.Auth, backofficeUserSvc, authorizationSvc),
		matchingHandler:       matchinghandler.New(authSvc, config.Auth, matchingSvc, matchingValidator, presenceSvc),
	}
}

func (s Server) Serve(ctx context.Context) error {
	e := echo.New()

	sc := echo.StartConfig{
		Address:         fmt.Sprintf(":%d", s.config.HttpServer.Port),
		GracefulTimeout: s.config.Application.GracefulShutdownTimeout,
		OnShutdownError: func(err error) {
			e.Logger.Error("graceful shutdown failed", "error", err)
		},
	}

	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogContentLength: true,
		LogRequestID:     true,
		LogHost:          true,
		LogMethod:        true,
		LogLatency:       true,
		LogRemoteIP:      true,
		LogResponseSize:  true,
		LogProtocol:      true,
		HandleError:      true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""
			if v.Error != nil {
				errMsg = v.Error.Error()
			}
			logger.Logger.Named("http-server").Info(
				"request",
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content-length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error", errMsg),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	e.GET("/health_check", s.healthCheck)

	s.userHandler.SetUserRoutes(e)
	s.backofficeUserHandler.SetUserRoutes(e)
	s.matchingHandler.SetUserRoutes(e)

	return sc.Start(ctx, e)
}

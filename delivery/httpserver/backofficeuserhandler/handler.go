package backofficeuserhandler

import (
	"github.com/younesbeheshti/gocast_game/service/authorizationservice"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/backofficeuserservice"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	authorizationSvc  authorizationservice.Service
	backofficeUserSvc backofficeuserservice.Service
}

func New(authSvc authservice.Service, authConfig authservice.Config, backOfficeUserSvc backofficeuserservice.Service, authorizationSvc authorizationservice.Service) Handler {
	return Handler{
		authSvc:           authSvc,
		authConfig:        authConfig,
		backofficeUserSvc: backOfficeUserSvc,
		authorizationSvc:  authorizationSvc,
	}
}

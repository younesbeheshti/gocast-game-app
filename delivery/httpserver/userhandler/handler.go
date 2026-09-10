package userhandler

import (
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/presenceservice"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"github.com/younesbeheshti/gocast_game/validator/uservalidator"
)

type Handler struct {
	authConfig    authservice.Config
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
	presenceSvc   presenceservice.Service
}

func New(authSvc authservice.Service, userSvc userservice.Service, validator uservalidator.Validator, authConfig authservice.Config, presenceSvc presenceservice.Service) Handler {
	return Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: validator,
		authConfig:    authConfig,
		presenceSvc:   presenceSvc,
	}
}

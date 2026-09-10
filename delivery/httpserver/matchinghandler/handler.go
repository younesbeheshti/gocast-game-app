package matchinghandler

import (
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
	"github.com/younesbeheshti/gocast_game/service/presenceservice"
	"github.com/younesbeheshti/gocast_game/validator/matchingvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
	presenceSvc       presenceservice.Service
}

func New(authSvc authservice.Service, authConfig authservice.Config, matchingSvc matchingservice.Service, validator matchingvalidator.Validator, presenceSve presenceservice.Service) Handler {
	return Handler{
		authSvc:           authSvc,
		authConfig:        authConfig,
		matchingSvc:       matchingSvc,
		matchingValidator: validator,
		presenceSvc:       presenceSve,
	}
}

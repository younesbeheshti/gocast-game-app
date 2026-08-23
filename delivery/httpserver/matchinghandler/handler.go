package matchinghandler

import (
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
	"github.com/younesbeheshti/gocast_game/validator/matchingvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
}

func New(authSvc authservice.Service, authConfig authservice.Config, matchingSvc matchingservice.Service, validator matchingvalidator.Validator) Handler {
	return Handler{
		authSvc:           authSvc,
		authConfig:        authConfig,
		matchingSvc:       matchingSvc,
		matchingValidator: validator,
	}
}

package userhandler

import (
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"github.com/younesbeheshti/gocast_game/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authSvc authservice.Service, userSvc userservice.Service, validator uservalidator.Validator) Handler {
	return Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: validator,
	}
}

package userservice

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
)

func (s *Service) Login(req param.LoginRequest) (*param.LoginResponse, error) {
	const op = "userservic.Login"
	// TODO - it would be better for user to have two separate methods for existence and getUserBYPhoneNumber

	//check the existence of phone number from repository
	//get the user by phone number
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage("unexpected").WithKind(richerror.KindUnexpected)
	}

	//compare user.pass with req.pass
	if user.HashedPassword != getMD5Hash(req.Password) {
		return nil, fmt.Errorf("username or passwword is invalid")
	}

	// generate jwt
	accessToken, err := s.auth.CreateAccessToken(*user)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}

	refreshToken, err := s.auth.CreateAccessToken(*user)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}

	return &param.LoginResponse{User: param.UserInfo{ID: user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
	},
		Tokens: param.Tokens{AccessToken: accessToken,
			RefreshToken: refreshToken},
	}, nil
}

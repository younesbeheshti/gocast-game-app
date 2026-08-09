package userservice

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
)

func (s *Service) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	const op = "userservic.Login"
	// TODO - it would be better for user to have two separate methods for existence and getUserBYPhoneNumber

	//check the existence of phone number from repository
	//get the user by phone number
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage("unexpected").WithKind(richerror.KindUnexpected)
	}

	if !exist {
		return nil, fmt.Errorf("username or password is invalid")
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

	return &dto.LoginResponse{User: dto.UserInfo{ID: user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
	},
		Tokens: dto.Tokens{AccessToken: accessToken,
			RefreshToken: refreshToken},
	}, nil
}

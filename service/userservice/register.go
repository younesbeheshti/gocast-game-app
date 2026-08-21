package userservice

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/param"
)

func (s *Service) Register(req param.RegisterRequest) (*param.RegisterResponse, error) {

	//TODO: verifying phone number with verification code

	//create new user in storage
	user := entity.User{
		ID:             0,
		PhoneNumber:    req.PhoneNumber,
		Name:           req.Name,
		HashedPassword: getMD5Hash(req.Password),
		Role:           entity.UserRole,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}

	//return created user
	return &param.RegisterResponse{
		User: param.UserInfo{
			ID:          createdUser.ID,
			PhoneNumber: createdUser.PhoneNumber,
			Name:        createdUser.Name,
		}}, nil
}

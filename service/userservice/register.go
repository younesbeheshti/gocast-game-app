package userservice

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/entity"
)

func (s *Service) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {

	//TODO: verifying phone number with verification code

	//create new user in storage
	user := entity.User{
		ID:             0,
		PhoneNumber:    req.PhoneNumber,
		Name:           req.Name,
		HashedPassword: getMD5Hash(req.Password),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}

	//return created user
	return &dto.RegisterResponse{
		User: dto.UserInfo{
			ID:          createdUser.ID,
			PhoneNumber: createdUser.PhoneNumber,
			Name:        createdUser.Name,
		}}, nil
}

package userservice

import (
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
)

func (s *Service) GetProfile(req dto.ProfileRequest) (*dto.ProfileResponse, error) {
	const op = "userservice.GetProfile"
	//getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {

		//TODO: can use rich error
		return nil, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"req": req})
	}

	return &dto.ProfileResponse{user.Name}, nil
}

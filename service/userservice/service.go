package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (*entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	entity.User
}

func (s *Service) Register(req RegisterRequest) (*RegisterResponse, error) {

	//TODO: verifying phone number with verification code

	//validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return nil, fmt.Errorf("invalid phone number")
	}

	//check uniqueness of phone number

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return nil, fmt.Errorf("unexpected error %w", err)
		} else {
			return nil, fmt.Errorf("phone number is not unique")
		}
	}

	//validate name
	if len(req.Name) < 3 {
		return nil, fmt.Errorf("name is too short")
	}

	// TODO - check the password with regex pattern
	// validate password
	if len(req.Password) < 8 {
		return nil, fmt.Errorf("password is too short, at least 8 is required")

	}

	//create new user in storage
	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    getMD5Hash(req.Password),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}
	//return created user
	return &RegisterResponse{*createdUser}, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])

}

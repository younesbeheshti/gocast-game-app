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
	GetUserByPhoneNumber(phoneNumber string) (*entity.User, bool, error)
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
	return &RegisterResponse{*createdUser}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct{}

func (s *Service) Login(req LoginRequest) (*LoginResponse, error) {

	// TODO - it would be better for user to have two separate methods for existence and getUserBYPhoneNumber

	//check the existence of phone number from repository
	//get the user by phone number
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}

	if !exist {
		return nil, fmt.Errorf("username or password is invalid")
	}

	//compare user.pass with req.pass
	if user.HashedPassword != getMD5Hash(req.Password) {
		return nil, fmt.Errorf("username or passwword is invalid")
	}

	return &LoginResponse{}, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])

}

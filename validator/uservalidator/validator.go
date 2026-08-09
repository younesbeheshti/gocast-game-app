package uservalidator

import "github.com/younesbeheshti/gocast_game/entity"

const (
	PhoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNubmer string) (*entity.User, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}

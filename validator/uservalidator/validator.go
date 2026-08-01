package uservalidator

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}

func (v Validator) ValidateRegisterRequest(req dto.RegisterRequest) (error, map[string]string) {
	const op = "uservalidator.ValidateRegisterRequest"

	if err := validation.ValidateStruct(&req,

		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),

		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9]{8,}$`))),

		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile("^09[0-9]{9}$")),
			validation.By(v.checkPhoneNumberUnique)),
	); err != nil {

		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for k, v := range errV {
				if v != nil {
					fieldErrors[k] = v.Error()
				}
			}
		}
		return richerror.New(op).WithMessage("invalid input").WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err), fieldErrors
	}

	return nil, nil
}

func (v Validator) checkPhoneNumberUnique(value interface{}) error {
	phoneNumber := value.(string)

	if isUnique, err := v.repo.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		} else {
			return fmt.Errorf("phone number is not unique")
		}
	}

	return nil

}

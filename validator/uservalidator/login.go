package uservalidator

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"log"
	"regexp"
)

func (v Validator) ValidateLoginRequest(req dto.LoginRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateLoginRequest"

	if err := validation.ValidateStruct(&req,

		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(PhoneNumberRegex)),
			validation.By(v.doesPhoneNumberExist)),

		validation.Field(&req.Password, validation.Required),
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
		return fieldErrors, richerror.New(op).WithMessage("invalid input").WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	}

	return nil, nil
}

func (v Validator) doesPhoneNumberExist(value interface{}) error {
	phoneNumber := value.(string)

	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		log.Println("log error", err)
		return fmt.Errorf("record not found")
	}

	return nil

}

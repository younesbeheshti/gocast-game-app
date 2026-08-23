package matchingvalidator

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
)

func (v Validator) ValidateAddToWaitingListRequest(req param.AddToWaitingListRequest) (map[string]string, error) {
	const op = "matching.validator.addToWaitingList"

	if err := validation.ValidateStruct(&req,

		validation.Field(&req.Category, validation.Required,
			validation.By(v.isCategoryValid)),
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

func (v Validator) isCategoryValid(value interface{}) error {
	category := value.(entity.Category)

	if !category.IsValid() {
		return fmt.Errorf("invalid category")
	}

	return nil
}

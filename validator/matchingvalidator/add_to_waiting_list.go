package matchingvalidator

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/param"
	"github.com/mobin-alz/gameapp/pkg/errmsg"
	"github.com/mobin-alz/gameapp/pkg/richerror"
)

func (v Validator) ValidateAddToWaitingList(req param.AddToWaitingListRequest) (map[string]string, error) {
	const op = "matching.ValidatorLoginRequest"
	if vErr := validation.ValidateStruct(&req,

		validation.Field(&req.Category, validation.Required,
			validation.By(v.isCategoryValid)),
	); vErr != nil {
		var errValidation validation.Errors
		fieldErrors := make(map[string]string)
		ok := errors.As(vErr, &errValidation)
		if ok {
			for key, value := range errValidation {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithError(vErr).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req})
	}

	return nil, nil

}

func (v Validator) isCategoryValid(value interface{}) error {
	category := value.(entity.Category)
	if !category.IsValid() {
		return fmt.Errorf(errmsg.ErrorMsgCategoryIsNotValid)
	}

	return nil

}

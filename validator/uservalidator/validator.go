package uservalidator

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	_ "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mobin-alz/gameapp/dto"
	"github.com/mobin-alz/gameapp/pkg/errmsg"
	"github.com/mobin-alz/gameapp/pkg/richerror"
	"regexp"
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

func (v Validator) ValidatorRegisterRequest(req dto.RegisterRequest) (error, map[string]string) {
	// TODO - we should verify phone number by verification code
	const op = "uservalidator.ValidatorRegisterRequest"
	// validate phone number
	if vErr := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9_-]{10,}$"))),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^09[0-9]{9}$")).Error("phone number is not valid"),
			validation.By(v.checkPhoneNumberUniqueness)),
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
		return richerror.New(op).
			WithError(vErr).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}), fieldErrors
	}

	return nil, nil

}

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	const op = "uservalidator.ValidatorcheckPhoneNumberUniqueness"
	phoneNumber := value.(string)
	if isUnique, err := v.repo.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf(errmsg.ErrorMsgPhoneNumberIsNotUnique)
		}
	}
	return nil
}

package uservalidator

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mobin-alz/gameapp/param"
	"github.com/mobin-alz/gameapp/pkg/errmsg"
	"github.com/mobin-alz/gameapp/pkg/richerror"
	"regexp"
)

func (v Validator) ValidatorLoginRequest(req param.LoginRequest) (map[string]string, error) {
	// TODO - we should verify phone number by verification code
	const op = "uservalidator.ValidatorLoginRequest"
	// validate phone number
	if vErr := validation.ValidateStruct(&req,

		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(PhoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
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
		return fieldErrors, richerror.New(op).
			WithError(vErr).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req})
	}

	return nil, nil

}

func (v Validator) doesPhoneNumberExist(value interface{}) error {
	const op = "uservalidator.doesPhoneNumberExist"
	phoneNumber := value.(string)
	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)

	if err != nil {
		return fmt.Errorf(errmsg.ErrorMsgPhoneNumberIsNotFound)
	}

	return nil
}

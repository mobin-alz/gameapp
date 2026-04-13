package uservalidator

import (
	_ "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mobin-alz/gameapp/entity"
)

const (
	PhoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
}
type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}

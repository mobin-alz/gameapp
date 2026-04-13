package userhandler

import (
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/userservice"
	"github.com/mobin-alz/gameapp/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}

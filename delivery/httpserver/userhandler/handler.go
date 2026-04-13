package userhandler

import (
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/userservice"
	"github.com/mobin-alz/gameapp/validator/uservalidator"
)

type Handler struct {
	authConfig    authservice.Config
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authConfig authservice.Config, authSvc authservice.Service,
	userSvc userservice.Service,
	userValidator uservalidator.Validator,
) Handler {
	return Handler{
		authConfig:    authConfig,
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}

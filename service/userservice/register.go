package userservice

import (
	"fmt"
	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/param"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	// validate password
	if len(req.Password) < 8 {
		return param.RegisterResponse{}, fmt.Errorf("password length should be greater than 8")
	}

	// hash password

	// create new user in storage (db , ...)
	user := entity.User{
		ID:          1,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    GetMD5Hash(req.Password),
	}
	createdUser, err := s.repo.RegisterUser(user)
	if err != nil {
		return param.RegisterResponse{}, err
	}

	return param.RegisterResponse{User: param.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	},
	}, nil
}

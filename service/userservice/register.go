package userservice

import (
	"fmt"
	"github.com/mobin-alz/gameapp/dto"
	"github.com/mobin-alz/gameapp/entity"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	// validate password
	if len(req.Password) < 8 {
		return dto.RegisterResponse{}, fmt.Errorf("password length should be greater than 8")
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
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{User: dto.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	},
	}, nil
}

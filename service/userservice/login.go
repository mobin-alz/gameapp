package userservice

import (
	"fmt"
	"github.com/mobin-alz/gameapp/dto"
	"github.com/mobin-alz/gameapp/pkg/richerror"
)

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {

	const op = "userservice.Login"
	//TODO-it would be better to user two separate method for existence check and getUserByPhoneNumber
	// check the existence of phone number from repository
	// get the user by phone_number
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).
			WithError(err).
			WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}
	// compare user.Password with the req.Password
	if user.Password != GetMD5Hash(req.Password) {
		return dto.LoginResponse{}, fmt.Errorf("invalid credentials")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("error on create token : %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("error on create refresh token : %w", err)
	}
	return dto.LoginResponse{
			User: dto.UserInfo{
				ID:          user.ID,
				PhoneNumber: user.PhoneNumber,
				Name:        user.Name,
			},
			Tokens: dto.Tokens{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}},
		nil
}

package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/mobin-alz/gameapp/dto"
	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/pkg/richerror"
)

type Repository interface {
	RegisterUser(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{repo: repo, auth: authGenerator}
}

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

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type LoginResponse struct {
	User   dto.UserInfo `json:"user"`
	Tokens Tokens       `json:"tokens"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {

	const op = "userservice.Login"
	//TODO-it would be better to user two separate method for existence check and getUserByPhoneNumber
	// check the existence of phone number from repository
	// get the user by phone_number
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.New(op).
			WithError(err).
			WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("invalid credentials")
	}

	// compare user.Password with the req.Password
	if user.Password != GetMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("invalid credentials")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("error on create token : %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("error on create refresh token : %w", err)
	}
	return LoginResponse{
			User: dto.UserInfo{
				ID:          user.ID,
				PhoneNumber: user.PhoneNumber,
				Name:        user.Name,
			},
			Tokens: Tokens{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}},
		nil
}

type ProfileRequest struct {
	UserID uint
}
type ProfileResponse struct {
	Name string `json:"name"`
}

// all request inputs for interactor/service should be sanitized.

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	const op = "userservice.Profile"
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I have not expected the repository call return "record not found" error,
		//because I assume the interactor input is sanitized.
		return ProfileResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"req": req})
	}
	// return User
	return ProfileResponse{Name: user.Name}, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

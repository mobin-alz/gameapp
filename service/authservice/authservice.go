package authservice

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mobin-alz/gameapp/entity"
	"strings"
	"time"
)

type Service struct {
	SignKey               string
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	accessSubject         string
	refreshSubject        string
}

func New(signKey, accessSubject, refreshSubject string,
	accessExpirationTime, refreshExpirationTime time.Duration) Service {
	return Service{
		SignKey:               signKey,
		AccessExpirationTime:  accessExpirationTime,
		RefreshExpirationTime: refreshExpirationTime,
		accessSubject:         accessSubject,
		refreshSubject:        refreshSubject,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.accessSubject, s.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(user entity.User) (accessToken string, err error) {
	return s.createToken(user.ID, s.refreshSubject, s.RefreshExpirationTime)
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	tokenString := strings.Replace(bearerToken, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.SignKey), nil

	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}

func (s Service) createToken(userID uint, subject string, expireDuration time.Duration) (string, error) {
	//create a signer for rsa 256
	// TODO - relace with 256 RS256

	// set our claims
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.SignKey))
	if err != nil {
		return "", err
	}

	// returns two param
	return tokenString, nil
}

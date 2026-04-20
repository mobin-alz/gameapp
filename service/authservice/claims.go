package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mobin-alz/gameapp/entity"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint        `json:"user_id"`
	Role   entity.Role `json:"role"`
}

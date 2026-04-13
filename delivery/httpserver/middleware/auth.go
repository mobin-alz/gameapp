package middleware

import (
	mv "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/pkg/constant"
	"github.com/mobin-alz/gameapp/service/authservice"
)

func Auth(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	return mv.WithConfig(mv.Config{
		ContextKey:    constant.AuthMiddlewareContextKey,
		SigningKey:    []byte(config.SignKey),
		SigningMethod: "HS256",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claims, nil
		},
	})
}

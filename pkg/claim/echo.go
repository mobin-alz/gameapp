package claim

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/config"
	"github.com/mobin-alz/gameapp/service/authservice"
)

func GetClaimsFromEchoContext(c echo.Context) *authservice.Claims {
	claims := c.Get(config.AuthMiddlewareContextKey)
	return claims.(*authservice.Claims)
}

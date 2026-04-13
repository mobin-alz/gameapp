package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/param"
	"github.com/mobin-alz/gameapp/pkg/constant"
	"github.com/mobin-alz/gameapp/pkg/httpmsg"
	"github.com/mobin-alz/gameapp/service/authservice"
	"net/http"
)

func getClaims(c echo.Context) *authservice.Claims {
	claims := c.Get(constant.AuthMiddlewareContextKey)
	return claims.(*authservice.Claims)
}

func (h Handler) userProfile(c echo.Context) error {
	claims := getClaims(c)

	resp, err := h.userSvc.Profile(param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, echo.Map{
			"error": msg,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

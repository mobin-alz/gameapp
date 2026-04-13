package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/dto"
	"github.com/mobin-alz/gameapp/pkg/httpmsg"
	"net/http"
)

func (h Handler) userProfile(c echo.Context) error {

	// validate jwt token and retrieve userID from token payload
	authHeader := c.Request().Header.Get("Authorization")

	claims, err := h.authSvc.ParseToken(authHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := h.userSvc.Profile(dto.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, echo.Map{
			"error": msg,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

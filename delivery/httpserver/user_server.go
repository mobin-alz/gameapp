package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/service/userservice"
	"net/http"
)

func (s Server) userRegister(c echo.Context) error {

	var req userservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, registerErr := s.userSvc.Register(req)
	if registerErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, registerErr.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (s Server) userLogin(c echo.Context) error {

	var lReq userservice.LoginRequest
	err := c.Bind(&lReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := s.userSvc.Login(lReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": "invalid credentials",
		})
	}

	return c.JSON(http.StatusOK, resp)

}

func (s Server) userProfile(c echo.Context) error {

	// validate jwt token and retrieve userID from token payload
	authHeader := c.Request().Header.Get("Authorization")

	claims, err := s.authSvc.ParseToken(authHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := s.userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{})
	}

	return c.JSON(http.StatusOK, resp)
}

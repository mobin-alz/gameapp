package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/dto"
	"github.com/mobin-alz/gameapp/pkg/httpmsg"
	"net/http"
)

func (s Server) userRegister(c echo.Context) error {

	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if fieldErrors, err := s.userValidator.ValidatorRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"error":   fieldErrors,
		})
	}

	resp, registerErr := s.userSvc.Register(req)
	if registerErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, registerErr.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (s Server) userLogin(c echo.Context) error {

	var lReq dto.LoginRequest
	err := c.Bind(&lReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := s.userSvc.Login(lReq)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, echo.Map{
			"error": msg,
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

	resp, err := s.userSvc.Profile(dto.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, echo.Map{
			"error": msg,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

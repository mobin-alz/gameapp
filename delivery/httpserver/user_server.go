package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/service/userservice"
	"net/http"
)

func (s Server) userRegister(c echo.Context) error {

	var uReq userservice.RegisterRequest
	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, registerErr := s.userSvc.Register(uReq)
	if registerErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, registerErr.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/param"
	"github.com/mobin-alz/gameapp/pkg/httpmsg"
	"net/http"
)

func (h Handler) userRegister(c echo.Context) error {

	var req param.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if fieldErrors, err := h.userValidator.ValidatorRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"error":   fieldErrors,
		})
	}

	resp, registerErr := h.userSvc.Register(req)
	if registerErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, registerErr.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

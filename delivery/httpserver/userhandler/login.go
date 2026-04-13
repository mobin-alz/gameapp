package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/dto"
	"github.com/mobin-alz/gameapp/pkg/httpmsg"
	"net/http"
)

func (h Handler) userLogin(c echo.Context) error {

	var req dto.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if fieldErrors, err := h.userValidator.ValidatorLoginRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"error":   fieldErrors,
		})
	}

	resp, err := h.userSvc.Login(req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, echo.Map{
			"error": msg,
		})
	}

	return c.JSON(http.StatusOK, resp)

}

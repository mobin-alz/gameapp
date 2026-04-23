package matchinghandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/param"
	"github.com/mobin-alz/gameapp/pkg/claim"
	"github.com/mobin-alz/gameapp/pkg/httpmsg"
	"net/http"
)

func (h Handler) addToWaitingList(c echo.Context) error {
	claims := claim.GetClaimsFromEchoContext(c)

	var req param.AddToWaitingListRequest
	req.UserID = claims.UserID
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if fieldErrors, err := h.matchingValidator.ValidateAddToWaitingList(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"error":   fieldErrors,
		})
	}

	resp, err := h.matchingSvc.AddToWaitingList(req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)
}

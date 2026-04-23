package matchinghandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/delivery/httpserver/middleware"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/matching")
	{
		userGroup.POST("/add-to-waiting-list", h.addToWaitingList, middleware.Auth(h.authSvc, h.authConfig))
	}

}

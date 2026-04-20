package backofficeuserhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/delivery/httpserver/middleware"
	"github.com/mobin-alz/gameapp/entity"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/backoffice/users")

	userGroup.GET("/", h.listUsers, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))
}

package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mobin-alz/gameapp/delivery/httpserver/middleware"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	{
		userGroup.POST("/register", h.userRegister)
		userGroup.POST("/login", h.userLogin)
		userGroup.GET("/profile", h.userProfile, middleware.Auth(h.authSvc, h.authConfig))

	}

}

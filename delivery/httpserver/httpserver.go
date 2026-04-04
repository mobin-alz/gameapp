package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mobin-alz/gameapp/config"
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/userservice"
)

type Server struct {
	config  config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service) Server {
	return Server{config: config, authSvc: authSvc, userSvc: userSvc}
}

func (s Server) Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/healthcheck", s.healthCheck)

	userGroup := e.Group("/users")
	{
		userGroup.POST("/register", s.userRegister)
		userGroup.POST("/login", s.userLogin)
		userGroup.GET("/profile", s.userProfile)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}

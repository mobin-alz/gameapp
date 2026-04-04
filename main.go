package main

import (
	"fmt"
	"github.com/mobin-alz/gameapp/config"
	"github.com/mobin-alz/gameapp/delivery/httpserver"
	"github.com/mobin-alz/gameapp/repository/mysql"
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/userservice"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {

	// Echo instance (engine)
	fmt.Println("Starting echo server")
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "gameappPassword",
			Port:     3306,
			Host:     "localhost",
			DBName:   "gameapp_db",
		},
	}

	authSvc, userSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)
	server.Serve()

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	repo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, repo)

	return authSvc, userSvc
}

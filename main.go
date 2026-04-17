package main

import (
	"fmt"
	"github.com/mobin-alz/gameapp/config"
	"github.com/mobin-alz/gameapp/delivery/httpserver"
	"github.com/mobin-alz/gameapp/repository/mysql"
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/userservice"
	"github.com/mobin-alz/gameapp/validator/uservalidator"
	"os"
	"strconv"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func getHTTPServerPort(fallback int) int {
	os.Getenv("HTTP_PORT")
	if port, err := strconv.Atoi(os.Getenv("GAMEAPP_PORT")); err == nil {
		return port
	}
	return fallback
}
func main() {

	// order loading values for configuration :
	// 1 - load default values(third priority)
	// 2 - read file and merge (overwrite)(second priority)
	// 3- get env values and merge (overwrite) (first priority)

	cfg2 := config.Load("config.yml")
	fmt.Printf("%+v\n", cfg2)

	// Echo instance (engine)
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: getHTTPServerPort(8080)},
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

	//TODO - add command for migrations
	//mgr := migrator.New(cfg.Mysql)
	//mgr.Up()

	authSvc, userSvc, userValidator := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator)

	fmt.Println("Start echo server")
	server.Serve()

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)

	repo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, repo)
	uV := uservalidator.New(repo)
	return authSvc, userSvc, uV
}

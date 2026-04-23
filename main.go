package main

import (
	"fmt"
	"github.com/mobin-alz/gameapp/adapter/redis"
	"github.com/mobin-alz/gameapp/config"
	"github.com/mobin-alz/gameapp/delivery/httpserver"
	"github.com/mobin-alz/gameapp/repository/migrator"
	"github.com/mobin-alz/gameapp/repository/mysql"
	"github.com/mobin-alz/gameapp/repository/mysql/mysqlaccesscontrol"
	"github.com/mobin-alz/gameapp/repository/mysql/mysqluser"
	"github.com/mobin-alz/gameapp/repository/redis/redismatching"
	"github.com/mobin-alz/gameapp/service/authorizationservice"
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/backofficeuserservice"
	"github.com/mobin-alz/gameapp/service/matchingservice"
	"github.com/mobin-alz/gameapp/service/userservice"
	"github.com/mobin-alz/gameapp/validator/matchingvalidator"
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

	cfg := config.Load("config.yml")
	fmt.Printf("%+v\n", cfg)

	//TODO - add command for migrations
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV)

	fmt.Println("Start echo server")
	server.Serve()

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator,
	backofficeuserservice.Service, authorizationservice.Service, matchingservice.Service, matchingvalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)

	MySqlRepo := mysql.New(cfg.Mysql)
	userMySql := mysqluser.New(MySqlRepo)
	userSvc := userservice.New(authSvc, userMySql)

	backofficeUserSvc := backofficeuserservice.New()

	aclMySql := mysqlaccesscontrol.New(MySqlRepo)
	authorizationSvc := authorizationservice.New(aclMySql)

	uV := uservalidator.New(userMySql)
	matchingV := matchingvalidator.New()

	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := redismatching.New(redisAdapter)
	matchingSvc := matchingservice.New(cfg.MatchingService, matchingRepo)

	return authSvc, userSvc, uV, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV
}

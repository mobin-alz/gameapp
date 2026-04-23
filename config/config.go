package config

import (
	"github.com/mobin-alz/gameapp/adapter/redis"
	"github.com/mobin-alz/gameapp/repository/mysql"
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/matchingservice"
)

type HTTPServer struct {
	Port int `koanf:"port"`
}
type Config struct {
	HTTPServer      HTTPServer             `koanf:"http_server"`
	Auth            authservice.Config     `koanf:"auth"`
	Mysql           mysql.Config           `koanf:"mysql"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
	Redis           redis.Config           `koanf:"redis"`
}

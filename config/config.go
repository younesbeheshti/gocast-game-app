package config

import (
	"github.com/younesbeheshti/gocast_game/adapter/redis"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
)

type Config struct {
	HttpServer      HttpServer             `koanf:"http_server"`
	Auth            authservice.Config     `koanf:"auth"`
	Psql            postgres.Config        `koanf:"postgres"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
	Redis           redis.Config           `koanf:"redis"`
}

type HttpServer struct {
	Port int `koanf:"port"`
}

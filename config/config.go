package config

import (
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	"github.com/younesbeheshti/gocast_game/service/authservice"
)

type Config struct {
	HttpServer HttpServer         `koanf:"http_server""`
	Auth       authservice.Config `koanf:"auth"`
	Psql       postgres.Config    `koanf:"postgres"`
}

type HttpServer struct {
	Port int `koanf:"port"`
}

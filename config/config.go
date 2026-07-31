package config

import (
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	"github.com/younesbeheshti/gocast_game/service/authservice"
)

type Config struct {
	HttpServer HttpServer
	Auth       authservice.Config
	Psql       postgres.Config
}

type HttpServer struct {
	Port int
}

package config

import (
	"github.com/younesbeheshti/gocast_game/adapter/redis"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	"github.com/younesbeheshti/gocast_game/scheduler"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
	"github.com/younesbeheshti/gocast_game/service/presenceservice"
	"time"
)

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}

type Config struct {
	Application     Application            `koanf:"application"`
	HttpServer      HttpServer             `koanf:"http_server"`
	Auth            authservice.Config     `koanf:"auth"`
	Psql            postgres.Config        `koanf:"postgres"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
	Redis           redis.Config           `koanf:"redis"`
	PresenceService presenceservice.Config `koanf:"presence_service"`
	Scheduler       scheduler.Config       `koanf:"scheduler"`
}

type HttpServer struct {
	Port int `koanf:"port"`
}

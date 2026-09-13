package main

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/adapter/redis"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/delivery/grpcserver/presenceserver"
	"github.com/younesbeheshti/gocast_game/repository/redis/redispresence"
	"github.com/younesbeheshti/gocast_game/service/presenceservice"
	"log"
)

func main() {

	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	redisAdapter := redis.New(cfg.Redis)
	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(cfg.PresenceService, presenceRepo)

	server := presenceserver.New(presenceSvc)
	server.Start()
}

package main

import (
	"context"
	"fmt"
	"github.com/younesbeheshti/gocast_game/adapter/redis"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/protobufencoder"
	"log"
)

func main() {

	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	redisAdapter := redis.New(cfg.Redis)

	topic := entity.MatchingUsersMatchedEvent
	mu := entity.MatchedPlayers{
		Category: entity.FootballCategory,
		UserIDs:  []uint{1, 2},
	}

	payloadStr := protobufencoder.EncodeEvent(entity.MatchingUsersMatchedEvent, mu)

	if err := redisAdapter.Client().Publish(context.Background(), string(topic), payloadStr).Err(); err != nil {
		panic(fmt.Sprintf("publish err: %v", err))
	}

}

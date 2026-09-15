package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/younesbeheshti/gocast_game/adapter/redis"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/contract/golang/matching"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/slice"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {

	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	redisAdapter := redis.New(cfg.Redis)

	topic := "matching.users_matched"
	mu := entity.MatchedPlayers{
		Category: entity.FootballCategory,
		UserIDs:  []uint{1, 2},
	}
	pbMu := matching.MatchedUsers{
		Category: string(mu.Category),
		UserIds:  slice.MapFromUintToUint64(mu.UserIDs),
	}

	payload, err := proto.Marshal(&pbMu)
	if err != nil {
		panic(err)
	}

	payloadStr := base64.StdEncoding.EncodeToString(payload)

	if err := redisAdapter.Client().Publish(context.Background(), topic, payloadStr).Err(); err != nil {
		panic(fmt.Sprintf("publish err: %v", err))
	}

}

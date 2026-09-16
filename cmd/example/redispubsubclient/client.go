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

	subscriber := redisAdapter.Client().Subscribe(context.Background(), string(topic))

	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		switch entity.Event(msg.Channel) {
		case topic:
			processUsersMatchedEvent(msg.Channel, msg.Payload)
		default:
			fmt.Println("Unsupported Message Channel:", msg.Channel)
		}

	}

}

func processUsersMatchedEvent(topic string, data string) {

	//payload, err := base64.StdEncoding.DecodeString(data)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//pbMu := &matching.MatchedUsers{}
	//if err := proto.Unmarshal(payload, pbMu); err != nil {
	//	log.Fatal(err)
	//}
	//
	//mu := &entity.MatchedPlayers{
	//	Category: entity.Category(pbMu.Category),
	//	UserIDs:  slice.MapFromUint64ToUint(pbMu.UserIds),
	//}

	payload := protobufencoder.DecodeEvent(entity.Event(topic), data)
	mu, ok := payload.(entity.MatchedPlayers)
	if !ok {
		panic("Failed to decode message")
	}

	fmt.Println("received message from", topic, "channel")
	fmt.Println("Matched Players:", mu)

}

package redis

import (
	"context"
	"github.com/younesbeheshti/gocast_game/entity"
	"log"
	"time"
)

func (a RedisAdapter) Publish(event entity.Event, payload string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := a.client.Publish(ctx, string(event), payload).Err(); err != nil {
		log.Println("redis publish error:", err)
		// TODO:log
		// TODO: update metrics
		return
	}

	// TODO: update metrics

}

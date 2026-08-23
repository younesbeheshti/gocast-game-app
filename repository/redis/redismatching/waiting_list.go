package redismatching

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"time"
)

// TODO - add to config in usecase layer ...
const WaitingListPrefix = "waitinglist"

func (d *DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = "redis.matching.AddToWaitingList"

	_, err := d.adapter.Client().ZAdd(context.Background(), fmt.Sprintf("%s:%s", WaitingListPrefix, category), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: fmt.Sprintf("%d", userID),
	}).Result()

	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

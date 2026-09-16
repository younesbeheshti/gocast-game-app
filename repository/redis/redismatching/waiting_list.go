package redismatching

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"github.com/younesbeheshti/gocast_game/pkg/timestamp"
	"log"
	"strconv"
	"time"
)

// TODO - add to config in usecase layer ...
const WaitingListPrefix = "waitinglist"

func (d *DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = "redis.matching.AddToWaitingList"

	_, err := d.adapter.Client().ZAdd(context.Background(), fmt.Sprintf("%s:%s", WaitingListPrefix, category), redis.Z{
		Score:  float64(timestamp.Now()),
		Member: fmt.Sprintf("%d", userID),
	}).Result()

	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

func (d *DB) GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error) {
	const op = "redis.matching.GetWaitingListByCategory"

	min := fmt.Sprintf("%d", timestamp.Add(-2000000*time.Hour))
	max := fmt.Sprintf("%d", timestamp.Now())

	list, err := d.adapter.Client().ZRangeByScoreWithScores(ctx, getCategory(category), &redis.ZRangeBy{
		Min: min,
		Max: max,
	}).Result()

	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	result := make([]entity.WaitingMember, 0)
	for _, item := range list {
		userID, _ := strconv.Atoi(item.Member.(string))

		result = append(result, entity.WaitingMember{
			UserID:    uint(userID),
			Timestamp: int64(item.Score),
			Category:  category,
		})
	}

	return result, nil
}

func (d *DB) RemoveUsersFromWaitingList(category entity.Category, userIDs []uint) {
	const op = "redis.matching.RemoveUsersFromWaitingList"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	members := make([]any, 0)
	for _, u := range userIDs {
		members = append(members, strconv.Itoa(int(u)))
	}

	numberOfRemovedMember, err := d.adapter.Client().ZRem(ctx, getCategory(category), members...).Result()
	if err != nil {
		log.Println("Error removing users:", err)
		// TODO: log metrics
	}

	log.Printf("Removed %d users from waited list", numberOfRemovedMember)
	// TODO: log metrics

}

func getCategory(category entity.Category) string {
	return fmt.Sprintf("%s:%s", WaitingListPrefix, category)
}

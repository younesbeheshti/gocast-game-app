package redispresence

import (
	"context"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"github.com/younesbeheshti/gocast_game/pkg/timestamp"
	"time"
)

func (d *DB) Upsert(ctx context.Context, key string, timestamp int64, expTime time.Duration) error {
	const op = richerror.Op("redispresence.Upsert")

	_, err := d.adapter.Client().Set(ctx, key, timestamp, expTime).Result()
	if err != nil {
		return richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}

	return nil
}

func (d *DB) GetPresence(ctx context.Context, prefixKey string, userIDs []uint) (map[uint]int64, error) {
	const op = richerror.Op("redispresence.GetPresence")

	m := make(map[uint]int64)

	for _, userID := range userIDs {
		m[userID] = timestamp.Add(time.Millisecond * -100)
	}

	return m, nil
}

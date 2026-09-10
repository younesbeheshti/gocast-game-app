package redispresence

import (
	"context"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
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

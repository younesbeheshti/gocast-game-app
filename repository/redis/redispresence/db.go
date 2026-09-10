package redispresence

import "github.com/younesbeheshti/gocast_game/adapter/redis"

type DB struct {
	adapter *redis.RedisAdapter
}

func New(adapter *redis.RedisAdapter) *DB {
	return &DB{adapter: adapter}
}

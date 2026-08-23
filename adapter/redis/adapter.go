package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}

type RedisAdapter struct {
	client *redis.Client
}

func New(cfg Config) *RedisAdapter {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &RedisAdapter{client: client}
}

func (r RedisAdapter) Client() *redis.Client {
	return r.client
}

package presenceservice

import (
	"context"
	"fmt"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"time"
)

type Config struct {
	ExpirationTime time.Duration `koanf:"expiration_time"`
	Prefix         string        `koanf:"prefix"`
}

type Repository interface {
	Upsert(ctx context.Context, key string, timestamp int64, expTime time.Duration) error
}

type Service struct {
	config Config
	repo   Repository
}

func New(config Config, repo Repository) Service {
	return Service{
		config: config,
		repo:   repo,
	}
}

func (s Service) Upsert(ctx context.Context, req *param.UpsertPresenceRequest) (*param.UpsertPresenceResponse, error) {
	const op = "service.upsert"

	if err := s.repo.Upsert(ctx, fmt.Sprintf("%s:%d", s.config.Prefix, req.UserID), req.Timestamp, s.config.ExpirationTime); err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return nil, nil
}

func (s Service) GetPresence(ctx context.Context, req *param.GetPresenceRequest) (*param.GetPresenceResponse, error) {

	const op = "service.getPresence"
	fmt.Println(op, "req:", req)

	// TODO - implement me
	return &param.GetPresenceResponse{Items: []param.GetPresenceItem{
		{UserID: 1, Timestamp: 12312312},
		{UserID: 2, Timestamp: 12332132},
	}}, nil
}

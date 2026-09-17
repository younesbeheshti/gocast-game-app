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
	GetPresence(ctx context.Context, key string, userIDs []uint) (map[uint]int64, error)
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
	list, err := s.repo.GetPresence(ctx, s.config.Prefix, req.UserIDs)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	resp := param.GetPresenceResponse{}
	for k, v := range list {
		resp.Items = append(resp.Items, param.GetPresenceItem{
			UserID:    k,
			Timestamp: v,
		})
	}
	return &resp, nil
}

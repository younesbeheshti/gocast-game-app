package matchingservice

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"time"
)

type Repository interface {
	AddToWaitingList(userID uint, category entity.Category) error
}

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}

type Service struct {
	repo   Repository
	config Config
}

func New(cfg Config, repo Repository) Service {
	return Service{config: cfg, repo: repo}
}

func (s Service) AddToWaitingList(req *param.AddToWaitingListRequest) (*param.AddToWaitingListResponse, error) {
	const op = "matchingservice.AddToWaitingList"

	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return &param.AddToWaitingListResponse{s.config.WaitingTimeout}, nil

}

func (s Service) MatchWaitedUsers(req *param.MatchWaitedUsersRequest) (*param.MatchWaitedUsersResponse, error) {
	fmt.Println("Match Waited Users")
	return nil, nil
}

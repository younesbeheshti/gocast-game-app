package matchingservice

import (
	"context"
	"fmt"
	funk "github.com/thoas/go-funk"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"github.com/younesbeheshti/gocast_game/pkg/timestamp"
	"log"
	"sync"
	"time"
)

type Repository interface {
	AddToWaitingList(userID uint, category entity.Category) error
	GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error)
}

type PresenceClient interface {
	GetPresence(ctx context.Context, req *param.GetPresenceRequest) (*param.GetPresenceResponse, error)
}

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}

type Service struct {
	repo           Repository
	config         Config
	presenceClient PresenceClient
}

func New(cfg Config, repo Repository, presenceClient PresenceClient) Service {
	return Service{config: cfg, repo: repo, presenceClient: presenceClient}
}

func (s Service) AddToWaitingList(req *param.AddToWaitingListRequest) (*param.AddToWaitingListResponse, error) {
	const op = "matchingservice.AddToWaitingList"

	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return &param.AddToWaitingListResponse{s.config.WaitingTimeout}, nil

}

func (s Service) MatchWaitedUsers(ctx context.Context, req *param.MatchWaitedUsersRequest) (*param.MatchWaitedUsersResponse, error) {
	const op = "matchingservice.MatchWaitedUsers"

	fmt.Println("Match Waited Users")

	var wg sync.WaitGroup
	for _, category := range entity.CategoryList() {
		wg.Add(1)
		go s.match(ctx, category, &wg)
	}

	wg.Wait()
	return nil, nil
}

func (s Service) match(ctx context.Context, category entity.Category, wg *sync.WaitGroup) {

	const op = "matchingservice.Match"
	defer wg.Done()

	list, err := s.repo.GetWaitingListByCategory(ctx, category)
	if err != nil {
		log.Println("Error getting waiting list:", op, err)
		return
	}

	userIDs := make([]uint, 0)
	for _, u := range list {
		userIDs = append(userIDs, u.UserID)
	}

	if len(userIDs) < 2 {
		return
	}

	presenceList, err := s.presenceClient.GetPresence(ctx, &param.GetPresenceRequest{UserIDs: userIDs})
	if err != nil {
		// TODO: log error
		// TODO: update metrics
		return
	}

	presenceUserIDs := make([]uint, len(list))
	for _, u := range presenceList.Items {
		presenceUserIDs = append(presenceUserIDs, u.UserID)
	}

	// TODO: merge presenceList with list based on userID
	// also consider the presence timestamp of each user
	// and remove users from waiting list if the user's timestamp is older than time.Now(-20 seconds)
	//if t < timestamp.Add(-30*time.Second) {
	// remove from list
	//}

	finalList := make([]entity.WaitingMember, 0)
	for i, l := range list {
		if funk.ContainsUInt(presenceUserIDs, l.UserID) && l.Timestamp < timestamp.Add(-20*time.Second) {
			finalList = append(finalList, l)
		} else {
			// remove from list
			list = append(list[:i], list[i+1:]...)
		}

	}

	for i := 0; i < len(list)-1; i = +2 {

		mu := entity.MatchedPlayers{
			Category: category,
			UserIDs:  []uint{finalList[i].UserID, finalList[i+1].UserID},
		}

		fmt.Println("Matched Players:", mu)

		//publish a new event for mu
		//remove mu users from waiting list
	}
}

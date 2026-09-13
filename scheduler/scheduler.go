package scheduler

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
	"log"
	"sync"
	"time"
)

type Config struct {
	MatchWaitedUsersIntervalInSeconds time.Duration `koanf:"match_waited_users_interval_in_seconds"`
}

type Scheduler struct {
	sch      gocron.Scheduler
	matchSvc matchingservice.Service
	config   Config
}

func New(matchSvc matchingservice.Service, config Config) Scheduler {

	sch, _ := gocron.NewScheduler()

	return Scheduler{
		sch:      sch,
		matchSvc: matchSvc,
		config:   config,
	}
}
func (s Scheduler) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	j, err := s.sch.NewJob(
		gocron.DurationJob(s.config.MatchWaitedUsersIntervalInSeconds),
		gocron.NewTask(s.MatchWaitedUsers),
	)
	if err != nil {
		log.Println("Error creating job:", err)
		return
	}

	s.sch.Start()

	fmt.Println("Scheduler started:", j.ID())

	<-ctx.Done()

	fmt.Println("Stopping scheduler...")

	if err := s.sch.StopJobs(); err != nil {
		log.Println("Error stopping scheduler jobs:", err)
	}

	if err := s.sch.Shutdown(); err != nil {
		log.Println("Error shutting down scheduler:", err)
	}

	fmt.Println("Scheduler stopped")
}

func (s Scheduler) MatchWaitedUsers() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	_, err := s.matchSvc.MatchWaitedUsers(ctx, &param.MatchWaitedUsersRequest{})
	if err != nil {
		// TODO: log err
		// TODO: update metrics
		log.Println("Error getting match waited users:", err)
	}
}

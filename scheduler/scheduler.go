package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
	"log"
	"time"
)

type Scheduler struct {
	sch      gocron.Scheduler
	matchSvc matchingservice.Service
}

func New(matchSvc matchingservice.Service) Scheduler {

	sch, _ := gocron.NewScheduler()

	return Scheduler{
		sch:      sch,
		matchSvc: matchSvc,
	}
}

func (s Scheduler) Start(done chan bool) {
	j, err := s.sch.NewJob(
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(
			s.MatchWaitedUsers,
		),
	)
	if err != nil {
		log.Println("Error creating job", err)
	}
	fmt.Println("Starting job", j.ID())
	s.sch.Start()
	defer s.sch.Shutdown()
	<-done

	fmt.Println("Scheduler stopped")
	s.sch.StopJobs()
}

func (s Scheduler) MatchWaitedUsers() {
	s.matchSvc.MatchWaitedUsers(&param.MatchWaitedUsersRequest{})
}

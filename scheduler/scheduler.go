package scheduler

import (
	"fmt"
	"time"
)

type Scheduler struct{}

func New() Scheduler {
	return Scheduler{}
}

func (s Scheduler) Start(done chan bool) {
	fmt.Println("Scheduler started")

	for {
		select {
		case d := <-done:
			fmt.Println("Scheduler done", d)
			return
		default:
			t := time.Now()
			time.Sleep(3 * time.Second)
			fmt.Println("Scheduler time", t)

		}
	}
}

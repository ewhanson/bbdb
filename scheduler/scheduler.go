package scheduler

import (
	"github.com/go-co-op/gocron"
	"time"
)

// Scheduler is a wrapper around gocron.Scheduler with some convenience functions for working with bbdb
type Scheduler struct {
	scheduler *gocron.Scheduler
}

// New returns a reference to a new Scheduler instance
func New() *Scheduler {
	return &Scheduler{
		scheduler: gocron.NewScheduler(time.Local),
	}
}

// GetScheduler returns a pointer to the underlying gocron.Scheduler
func (s Scheduler) GetScheduler() *gocron.Scheduler {
	return s.scheduler
}

// Start runs the gocron.Scheduler start methods
func (s Scheduler) Start() {
	s.scheduler.StartAsync()
}

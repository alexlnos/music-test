package usecases

import (
	"os"
	"time"
)

type Scheduler struct {
	siteChecker *SiteChecker
	ticker      *time.Ticker
	quit        chan os.Signal
}

func NewScheduler(checker *SiteChecker, ticker *time.Ticker, quit chan os.Signal) *Scheduler {
	return &Scheduler{
		siteChecker: checker,
		ticker:      ticker,
		quit:        quit,
	}
}

func (s *Scheduler) Start() {
	go func() {
		for {
			select {
			case <-s.ticker.C:
				s.siteChecker.CheckSites()
			case <-s.quit:
				return
			}
		}
	}()
}

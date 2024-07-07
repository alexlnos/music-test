package usecases

import (
	"time"
)

type Scheduler struct {
	siteChecker *SiteChecker
	ticker      *time.Ticker
	urls        []string
	quit        chan struct{} // Добавляем канал для остановки
}

func NewScheduler(checker *SiteChecker, ticker *time.Ticker, urls []string) *Scheduler {
	return &Scheduler{
		siteChecker: checker,
		ticker:      ticker,
		urls:        urls,
		quit:        make(chan struct{}), // Инициализируем канал
	}
}

func (s *Scheduler) Start() {
	func() {
		for {
			select {
			case <-s.ticker.C: // Слушаем тикер
				for _, url := range s.urls {
					s.siteChecker.CheckSite(url) // Запускаем проверку в отдельной горутине
				}
			case <-s.quit: // Слушаем сигнал остановки
				return
			}
		}
	}()
}

func (s *Scheduler) Stop() {
	close(s.quit) // Закрываем канал для остановки работы Scheduler
}

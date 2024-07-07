package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"test-music/internal/adapters"
	"test-music/internal/config"
	"test-music/internal/usecases"
	"time"
)

func distributeSites(sites []string, numWorkers int) [][]string {
	var divided [][]string
	chunkSize := (len(sites) + numWorkers - 1) / numWorkers // Округление вверх

	for i := 0; i < len(sites); i += chunkSize {
		end := i + chunkSize
		if end > len(sites) {
			end = len(sites)
		}
		divided = append(divided, sites[i:end])
	}
	return divided
}

func main() {
	cfg := config.LoadConfig()

	client := adapters.NewHTTPClient()
	checker := usecases.NewSiteChecker(client)

	fmt.Println("Enter the URLs, each on a new line, then press Ctrl-D:")
	scanner := bufio.NewScanner(os.Stdin)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	urlGroups := distributeSites(urls, cfg.NumWorkers)
	var wg sync.WaitGroup

	// Запускаем Scheduler для каждой группы URL
	for _, group := range urlGroups {
		wg.Add(1)
		go func(group []string) {
			ticker := time.NewTicker(cfg.CheckInterval)
			defer ticker.Stop()
			scheduler := usecases.NewScheduler(checker, ticker, group)
			scheduler.Start()
			wg.Done()
		}(group)
	}

	println("11111")
	wg.Wait() // Ждем завершения всех Scheduler
}

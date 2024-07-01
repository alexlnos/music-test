package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test-music/internal/adapters"
	"test-music/internal/config"
	"test-music/internal/usecases"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	interval := cfg.CheckInterval
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the URLs, each on a new line, then press Ctrl-D:")
	urls := []string{}
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	println("Start check sites...")

	client := adapters.NewHTTPClient()
	siteChecker := usecases.NewSiteChecker(client, urls)
	scheduler := usecases.NewScheduler(siteChecker, ticker, quit)

	scheduler.Start()

	<-quit
	fmt.Println("Shutting down the application...")
}

package usecases

import (
	"fmt"
	"strings"
	"test-music/internal/adapters"
	"test-music/internal/models"
	"time"
)

type SiteChecker struct {
	httpClient    adapters.HTTPClient
	urls          []string
	lastFailTime  map[string]time.Time
	lastCheckFail map[string]bool
}

func NewSiteChecker(client adapters.HTTPClient, urls []string) *SiteChecker {
	return &SiteChecker{
		httpClient:    client,
		urls:          urls,
		lastFailTime:  make(map[string]time.Time),
		lastCheckFail: make(map[string]bool),
	}
}

func (sc *SiteChecker) CheckSites() {
	for _, url := range sc.urls {
		result := sc.httpClient.Fetch(url)
		site := models.Site{URL: url}
		if strings.Contains(result, "advmusic.com") {
			if sc.lastCheckFail[url] {
				duration := time.Since(sc.lastFailTime[url])
				fmt.Printf("%s - recovered after %v\n", site.URL, duration)
				sc.lastCheckFail[url] = false
			} else {
				fmt.Printf("%s ok\n", site.URL)
			}
		} else {
			if !sc.lastCheckFail[url] {
				sc.lastFailTime[url] = time.Now()
				sc.lastCheckFail[url] = true
			}
			fmt.Printf("%s - fail\n", site.URL)
		}
	}
}

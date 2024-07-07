package usecases

import (
	"fmt"
	"strings"
	"test-music/internal/adapters"
)

type SiteChecker struct {
	httpClient adapters.HTTPClient
}

func NewSiteChecker(client adapters.HTTPClient) *SiteChecker {
	return &SiteChecker{httpClient: client}
}

func (sc *SiteChecker) CheckSite(url string) {
	result := sc.httpClient.Fetch(url)
	if strings.Contains(result, "advmusic.com") {
		fmt.Printf("%s ok\n", url)
	} else {
		fmt.Printf("%s - fail\n", url)
	}
}

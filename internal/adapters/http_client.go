package adapters

import (
	"io/ioutil"
	"net/http"
)

type HTTPClient interface {
	Fetch(url string) string
}

type httpClient struct{}

func NewHTTPClient() HTTPClient {
	return &httpClient{}
}

func (c *httpClient) Fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

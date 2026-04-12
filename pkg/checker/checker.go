package checker

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	URL        string        `json:"url"`
	StatusCode int           `json:"status_code"`
	Latency    time.Duration `json:"latency"`
	Error      string        `json:"error,omitempty"`
}

type Checker struct {
	Client  *http.Client
	Timeout time.Duration
}

func New(timeout time.Duration) *Checker {
	return &Checker{
		Client:  &http.Client{Timeout: timeout},
		Timeout: timeout,
	}
}

func (c *Checker) Check(url string) Result {
	start := time.Now()
	resp, err := c.Client.Get(url)
	latency := time.Since(start)

	if err != nil {
		return Result{URL: url, Error: fmt.Sprintf("%v", err), Latency: latency}
	}
	defer resp.Body.Close()

	return Result{URL: url, StatusCode: resp.StatusCode, Latency: latency}
}
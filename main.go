package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func checkURL(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		results <- fmt.Sprintf("[FAIL] %s - %v", url, err)
		return
	}
	defer resp.Body.Close()
	results <- fmt.Sprintf("[%d] %s", resp.StatusCode, url)
}

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Println("Usage: netutils <url1> <url2> ...")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, &wg, results)
	}

	go func() { wg.Wait(); close(results) }()

	for r := range results {
		fmt.Println(r)
	}
}
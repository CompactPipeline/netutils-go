package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestCheckURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	var wg sync.WaitGroup
	results := make(chan string, 1)
	wg.Add(1)

	go checkURL(server.URL, &wg, results)
	wg.Wait()
	close(results)

	r := <-results
	if r == "" {
		t.Fatal("Expected result, got empty string")
	}
}

func TestCheckURL_InvalidURL(t *testing.T) {
	var wg sync.WaitGroup
	results := make(chan string, 1)
	wg.Add(1)

	go checkURL("http://invalid.localhost.test", &wg, results)
	wg.Wait()
	close(results)

	r := <-results
	if r == "" {
		t.Fatal("Expected error result")
	}
}
package cmd

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/user/netutils-go/pkg/checker"
	"github.com/user/netutils-go/pkg/output"
)

var (
	timeout    int
	jsonOutput bool
	workers    int
)

var rootCmd = &cobra.Command{
	Use:   "netutils [urls...]",
	Short: "Concurrent network utility tool",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runCheck,
}

func init() {
	rootCmd.Flags().IntVarP(&timeout, "timeout", "t", 10, "Request timeout in seconds")
	rootCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output as JSON")
	rootCmd.Flags().IntVarP(&workers, "workers", "w", 5, "Max concurrent workers")
}

func runCheck(cmd *cobra.Command, args []string) error {
	c := checker.New(time.Duration(timeout) * time.Second)
	var mu sync.Mutex
	var results []checker.Result
	var wg sync.WaitGroup

	sem := make(chan struct{}, workers)
	for _, url := range args {
		wg.Add(1)
		sem <- struct{}{}
		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()
			r := c.Check(u)
			mu.Lock()
			results = append(results, r)
			mu.Unlock()
		}(url)
	}
	wg.Wait()

	var f output.Formatter
	if jsonOutput {
		f = output.NewJSON(true)
	} else {
		f = output.NewText()
	}
	return f.Format(results)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
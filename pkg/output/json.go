package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/user/netutils-go/pkg/checker"
)

type Formatter interface {
	Format(results []checker.Result) error
}

type JSONFormatter struct {
	Writer io.Writer
	Pretty bool
}

func NewJSON(pretty bool) *JSONFormatter {
	return &JSONFormatter{Writer: os.Stdout, Pretty: pretty}
}

func (f *JSONFormatter) Format(results []checker.Result) error {
	enc := json.NewEncoder(f.Writer)
	if f.Pretty {
		enc.SetIndent("", "  ")
	}
	return enc.Encode(results)
}

type TextFormatter struct {
	Writer io.Writer
}

func NewText() *TextFormatter {
	return &TextFormatter{Writer: os.Stdout}
}

func (f *TextFormatter) Format(results []checker.Result) error {
	for _, r := range results {
		if r.Error != "" {
			fmt.Fprintf(f.Writer, "[FAIL] %s - %s\n", r.URL, r.Error)
		} else {
			fmt.Fprintf(f.Writer, "[%d] %s (%s)\n", r.StatusCode, r.URL, r.Latency)
		}
	}
	return nil
}
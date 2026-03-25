package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Timeout  int      `json:"timeout"`
	Workers  int      `json:"workers"`
	Format   string   `json:"format"`
	Headers  map[string]string `json:"headers"`
	URLs     []string `json:"urls"`
}

func Default() *Config {
	return &Config{
		Timeout: 10,
		Workers: 5,
		Format:  "text",
		Headers: map[string]string{},
	}
}

func Load(path string) (*Config, error) {
	if path == "" {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, ".netutils.json")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Default(), nil
	}

	cfg := Default()
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
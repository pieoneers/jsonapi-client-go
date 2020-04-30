package client

import (
	"time"
)

type Config struct {
	BaseURL string
	Timeout time.Duration
}

func NewConfig(c Config) Config {
	defaults := Config{
		BaseURL: "http://localhost",
		Timeout: time.Second * 10,
	}

	if len(c.BaseURL) == 0 {
		c.BaseURL = defaults.BaseURL
	}

	if c.Timeout == 0 {
		c.Timeout = defaults.Timeout
	}

	return c
}

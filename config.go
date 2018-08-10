package client

import (
  "time"
)

type Config struct {
  BaseURL string
  Timeout time.Duration
}

func NewConfig(c Config) Config {
  config := Config{
    BaseURL: "http://localhost",
    Timeout: time.Second * 10,
  }

  if len(c.BaseURL) > 0 {
    config.BaseURL = c.BaseURL
  }

  if c.Timeout > 0 {
    config.Timeout = c.Timeout
  }

  return config
}

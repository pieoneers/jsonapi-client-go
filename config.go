// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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

// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"time"
)

//Config represents client configuration data, like base URL and client timeout
//Base URL used as prefix for all requests
type Config struct {
	BaseURL string
	Timeout time.Duration
}

// NewConfig check the provided config values, if some value is empty it will fill it by default value.
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

// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client_test

import (
	"time"

	. "github.com/pieoneers/jsonapi-client-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Describe("NewConfig", func() {

		When("configuration is default", func() {
			var config Config

			baseURL := "http://localhost"
			timeout := time.Second * 10

			BeforeEach(func() {
				config = NewConfig(Config{})
			})

			It("should use default base URL", func() {
				Ω(config.BaseURL).Should(Equal(baseURL))
			})

			It("should use default timeout", func() {
				Ω(config.Timeout).Should(Equal(timeout))
			})
		})

		When("configuration is specified", func() {
			var config Config

			baseURL := "https://api.pieoneers.com"
			timeout := time.Second * 1

			BeforeEach(func() {
				config = NewConfig(Config{
					BaseURL: baseURL,
					Timeout: timeout,
				})
			})

			It("should use default base URL", func() {
				Ω(config.BaseURL).Should(Equal(baseURL))
			})

			It("should use default timeout", func() {
				Ω(config.Timeout).Should(Equal(timeout))
			})
		})
	})
})

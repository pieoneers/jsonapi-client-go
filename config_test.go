package client_test

import (
  "time"

  . "github.com/pieoneers/jsonapi-client-go.git"

  . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

  Describe("NewConfig", func() {

    When("configuration is default", func() {
      var config Config

      expectedBaseURL := "http://localhost"
      expectedTimeout := time.Second * 10

      BeforeEach(func() {
        config = NewConfig(Config{})
      })

      It("should use default base URL", func() {
        Ω(config.BaseURL).Should(Equal(expectedBaseURL))
      })

      It("should use default timeout", func() {
        Ω(config.Timeout).Should(Equal(expectedTimeout))
      })
    })

    When("configuration is specified", func() {
      var config Config

      expectedBaseURL := "https://api.pieoneers.com"
      expectedTimeout := time.Second * 1

      BeforeEach(func() {
        config = NewConfig(Config{
          BaseURL: expectedBaseURL,
          Timeout: expectedTimeout,
        })
      })

      It("should use default base URL", func() {
        Ω(config.BaseURL).Should(Equal(expectedBaseURL))
      })

      It("should use default timeout", func() {
        Ω(config.Timeout).Should(Equal(expectedTimeout))
      })
    })
  })
})

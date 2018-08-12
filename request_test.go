package client_test

import (
  "io/ioutil"

  . "github.com/pieoneers/jsonapi-client-go"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {

  Describe("NewRequest", func() {

    Describe("Method", func() {
      for _, method := range []string{ "GET", "POST" } {

        When(method, func() {
          var request *Request

          BeforeEach(func() {
            request, _ = NewRequest(method, "/foo", nil)
          })

          It("should be equal GET", func() {
            Ω(request.Method).Should(Equal(method))
          })
        })
      }
    })

    Describe("URL", func() {
      var request *Request

      url := "/bar"

      BeforeEach(func() {
        request, _ = NewRequest("GET", url, nil)
      })

      It("should be correct", func() {
        Ω(request.URL).Should(Equal(url))
      })
    })

    Describe("Header", func() {
      var request *Request

      When("there is no payload", func() {
        BeforeEach(func() {
          request, _ = NewRequest("GET", "/books", nil)
        })

        It("should have correct Accept header", func() {
          actual := request.Header.Get("Accept")
          Ω(actual).Should(Equal("application/vnd.api+json"))
        })

        It("should have no Content-Type header", func() {
          actual := request.Header.Get("Content-Type")
          Ω(actual).Should(Equal(""))
        })
      })

      When("there is payload", func() {
        BeforeEach(func() {
          request, _ = NewRequest("POST", "/books", Book{
            Title: "An Introduction to Programming in Go",
            Year: "2012",
          })
        })

        It("should have correct Accept header", func() {
          actual := request.Header.Get("Accept")
          Ω(actual).Should(Equal("application/vnd.api+json"))
        })

        It("should have correct Content-Type header", func() {
          actual := request.Header.Get("Content-Type")
          Ω(actual).Should(Equal("application/vnd.api+json"))
        })
      })
    })

    Describe("Body", func() {
      var request *Request

      When("there is no payload", func() {
        BeforeEach(func() {
          request, _ = NewRequest("GET", "/books", nil)
        })

        It("should have empty body", func() {
          actual := request.Body
          Ω(actual).Should(BeNil())
        })
      })

      When("there is payload", func() {
        payload := Book{
          Title: "An Introduction to Programming in Go",
          Year: "2012",
        }

        BeforeEach(func() {
          request, _ = NewRequest("POST", "/books", payload)
        })

        It("should have book payload in body", func() {
          actual, _ := ioutil.ReadAll(request.Body)
          expected, _ := Template("book-payload", payload)

          Ω(actual).Should(MatchJSON(expected))
        })
      })
    })
  })
})

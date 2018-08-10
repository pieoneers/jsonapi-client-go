package client_test

import (
  "io/ioutil"
  . "github.com/pieoneers/jsonapi-client-go.git"
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
            request, _ = NewRequest(method, "/foobar", nil)
          })

          It("should be equal GET", func() {
            Ω(request.Method).Should(Equal(method))
          })
        })
      }
    })

    Describe("URL", func() {
      var request *Request

      expectedURL := "/foobar"

      BeforeEach(func() {
        request, _ = NewRequest("GET", expectedURL, nil)
      })

      It("should be correct", func() {
        Ω(request.URL).Should(Equal(expectedURL))
      })
    })

    Describe("Header", func() {
      var request *Request

      When("there is no payload", func() {
        BeforeEach(func() {
          request, _ = NewRequest("GET", "/users", nil)
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
          request, _ = NewRequest("POST", "/users", User{
            Email:     "andrew@pieoneers.com",
            Password:  "password",
            FirstName: "Andrew",
            LastName:  "Manshin",
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
          request, _ = NewRequest("GET", "/users", nil)
        })

        It("should have empty body", func() {
          actual := request.Body
          Ω(actual).Should(BeNil())
        })
      })

      When("there is payload", func() {
        BeforeEach(func() {
          request, _ = NewRequest("POST", "/users", User{
            Email:     "andrew@pieoneers.com",
            Password:  "password",
            FirstName: "Andrew",
            LastName:  "Manshin",
          })
        })

        It("should have jsonapi representation of user in body", func() {
          buf, _ := ioutil.ReadAll(request.Body)
          actual := string(buf)
          Ω(actual).Should(MatchJSON(`
            {
              "data": {
                "type": "users",
                "attributes": {
                  "email": "andrew@pieoneers.com",
                  "password": "password",
                  "first_name": "Andrew",
                  "last_name": "Manshin"
                }
              }
            }
          `))
        })
      })
    })
  })
})

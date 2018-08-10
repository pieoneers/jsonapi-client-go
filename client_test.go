package client_test

import (
  "net/http"
  "io/ioutil"

  . "github.com/pieoneers/jsonapi-client-go.git"

  . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

  Describe("Get", func() {
    var request *Request

    expectedURL := "/foo"

    BeforeEach(func() {
      request, _ = client.Get(expectedURL)
    })

    It("should return request with correct method", func() {
      Ω(request.Method).Should(Equal("GET"))
    })

    It("should return request with correct url", func() {
      Ω(request.URL).Should(Equal(expectedURL))
    })
  })

  Describe("Post", func() {
    var request *Request

    expectedURL := "/bar"

    BeforeEach(func() {
      request, _ = client.Post(expectedURL, User{
        Email:     "andrew@pieoneers.com",
        Password:  "password",
        FirstName: "Andrew",
        LastName:  "Manshin",
      })
    })

    It("should return request with correct method", func() {
      Ω(request.Method).Should(Equal("POST"))
    })

    It("should return request with correct url", func() {
      Ω(request.URL).Should(Equal(expectedURL))
    })

    It("should return request with correct body", func() {
      buf, _ := ioutil.ReadAll(request.Body)
      body := string(buf)

      Ω(body).Should(MatchJSON(`
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

  Describe("Do", func() {
    var (
      res *Response
      err error
    )

    var target User

    BeforeEach(func() {
      req, _ := client.Post("/users", User{
        Email:     "andrew@pieoneers.com",
        Password:  "password",
        FirstName: "Andrew",
        LastName:  "Manshin",
      })

      res, err = client.Do(req, &target)
    })

    It("should respond with 200", func() {
      Ω(res.StatusCode).Should(Equal(http.StatusOK))
    })

    It("should unmarshal response body to target", func() {
      Ω(target).Should(Equal(User{
        ID:        "1",
        Email:     "andrew@pieoneers.com",
        Password:  "password",
        FirstName: "Andrew",
        LastName:  "Manshin",
      }))
    })

    It("should not have error occured", func() {
      Ω(err).ShouldNot(HaveOccurred())
    })
  })
})

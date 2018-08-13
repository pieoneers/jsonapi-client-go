package client_test

import (
  "net/http"
  "io/ioutil"
  "github.com/pieoneers/jsonapi-go"

  . "github.com/pieoneers/jsonapi-client-go"

  . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

  Describe("Get", func() {
    var request *Request

    url := "/foo"

    BeforeEach(func() {
      request, _ = client.Get(url)
    })

    It("should return request with correct method", func() {
      Ω(request.Method).Should(Equal("GET"))
    })

    It("should return request with correct url", func() {
      Ω(request.URL).Should(Equal(url))
    })
  })

  Describe("Post", func() {
    var request *Request

    url := "/bar"

    book := Book{
      Title: "An Introduction to Programming in Go",
      Year: "2012",
    }

    BeforeEach(func() {
      request, _ = client.Post(url, book)
    })

    It("should return request with correct method", func() {
      Ω(request.Method).Should(Equal("POST"))
    })

    It("should return request with correct url", func() {
      Ω(request.URL).Should(Equal(url))
    })

    It("should return request with correct body", func() {
      actual, _ := ioutil.ReadAll(request.Body)
      expected, _ := Template("book-payload", book)

      Ω(actual).Should(MatchJSON(expected))
    })
  })

  Describe("Do", func() {
    var (
      res *Response
      err error
    )

    Describe("GET /books", func() {

      var target []*Book

      BeforeEach(func() {
        target = []*Book{}

        req, _ := client.Get("/books")

        res, err = client.Do(req, &target)
      })

      It("should have 200 status", func() {
        Ω(res.StatusCode).Should(Equal(http.StatusOK))
      })

      It("should unmarshal response body to target", func() {
        Ω(target).Should(Equal([]*Book{
          {
            ID: "1",
            Title: "An Introduction to Programming in Go",
            Year: "2012",
          },
          {
            ID: "2",
            Title: "Introducing Go",
            Year: "2016",
          },
        }))
      })
    })

    Describe("POST /books", func() {
      var target Book

      BeforeEach(func() {
        target = Book{}
      })

      When("successful", func() {

        BeforeEach(func() {
          req, _ := client.Post("/books/successful", Book{
            Title: "An Introduction to Programming in Go",
            Year: "2012",
          })

          res, err = client.Do(req, &target)
        })

        It("should respond with 201 status", func() {
          Ω(res.StatusCode).Should(Equal(http.StatusCreated))
        })

        It("should unmarshal response body to target", func() {
          Ω(target).Should(Equal(Book{
            ID:    "1",
            Title: "An Introduction to Programming in Go",
            Year:  "2012",
          }))
        })
      })

      When("with errors", func() {

        BeforeEach(func() {
          req, _ := client.Post("/books/unsuccessful", Book{
            Title: "",
            Year: "2012",
          })

          res, err = client.Do(req, &target)
        })

        It("should respond with 403 status", func() {
          Ω(res.StatusCode).Should(Equal(http.StatusForbidden))
        })

        It("should leave target empty", func() {
          Ω(target).Should(Equal(Book{}))
        })

        It("should store errors in response", func() {
          Ω(res.Document.Errors).Should(Equal([]*jsonapi.ErrorObject{
            {
              Title: "is required",
              Source: jsonapi.ErrorObjectSource{
                Pointer: "/data/attributes/title",
              },
            },
          }))
        })
      })

      When("there is no target", func() {

        BeforeEach(func() {
          req, _ := client.Post("/books/successful", Book{
            Title: "An Introduction to Programming in Go",
            Year: "2012",
          })

          res, err = client.Do(req, nil)
        })

        It("should not have document assigned", func() {
          Ω(res.Document).Should(BeNil())
        })

        It("should not have error occurred", func() {
          Ω(err).ShouldNot(HaveOccurred())
        })
      })
    })
  })
})

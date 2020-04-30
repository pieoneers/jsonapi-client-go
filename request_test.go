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
			for _, method := range []string{"GET", "POST"} {

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
			var req *Request

			url := "/bar?include=foo"

			BeforeEach(func() {
				req, _ = NewRequest("GET", url, nil)
			})

			It("should set correct path", func() {
				Ω(req.URL.Path).Should(Equal("/bar"))
			})

			It("should set correct query", func() {
				Ω(req.URL.RawQuery).Should(Equal("include=foo"))
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
						Year:  "2012",
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
					Year:  "2012",
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

	Describe("RequestURI", func() {
		var (
			req *Request
			url string
		)

		BeforeEach(func() {
			req, _ = NewRequest("GET", "/foo", nil)
		})

		JustBeforeEach(func() {
			url = req.RequestURI()
		})

		It("should return correct url", func() {
			Ω(url).Should(Equal("/foo"))
		})

		When("query is passed in url", func() {

			BeforeEach(func() {
				req, _ = NewRequest("GET", "/foo?include=bar,baz", nil)
			})

			It("should return correct url", func() {
				Ω(url).Should(Equal("/foo?include=bar%2Cbaz"))
			})
		})

		When("query is set afterwards", func() {

			BeforeEach(func() {
				query := req.Query
				query.Set("include", "bar,baz")
				query.Set("filter[id]", "quux")
			})

			It("should return correct url", func() {
				Ω(url).Should(Equal("/foo?filter%5Bid%5D=quux&include=bar%2Cbaz"))
			})
		})
	})
})

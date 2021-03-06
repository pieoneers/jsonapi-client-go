// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client_test

import (
	"github.com/pieoneers/jsonapi-go"
	"io/ioutil"
	"net/http"

	. "github.com/pieoneers/jsonapi-client-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

	Describe("Get", func() {
		var request *Request

		path := "/foo"

		BeforeEach(func() {
			request, _ = client.Get(path)
		})

		It("should return request with correct method", func() {
			Ω(request.Method).Should(Equal("GET"))
		})

		It("should return request with correct path", func() {
			Ω(request.URL.Path).Should(Equal(path))
		})
	})

	Describe("Head", func() {
		var request *Request

		path := "/foo"

		BeforeEach(func() {
			request, _ = client.Head(path)
		})

		It("should return request with correct method", func() {
			Ω(request.Method).Should(Equal("HEAD"))
		})

		It("should return request with correct path", func() {
			Ω(request.URL.Path).Should(Equal(path))
		})
	})

	Describe("Post", func() {
		var request *Request

		path := "/bar"

		book := Book{
			Title: "An Introduction to Programming in Go",
			Year:  "2012",
		}

		BeforeEach(func() {
			request, _ = client.Post(path, book)
		})

		It("should return request with correct method", func() {
			Ω(request.Method).Should(Equal("POST"))
		})

		It("should return request with correct path", func() {
			Ω(request.URL.Path).Should(Equal(path))
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
			var target Books

			BeforeEach(func() {
				target = Books{}

				req, _ := client.Get("/books")

				res, err = client.Do(req, &target)
			})

			It("should have 200 status", func() {
				Ω(res.StatusCode).Should(Equal(http.StatusOK))
			})

			It("should unmarshal resource objects collection into target", func() {
				Ω(target).Should(Equal(Books{
					{
						ID:    "1",
						Title: "An Introduction to Programming in Go",
						Year:  "2012",
					},
					{
						ID:    "2",
						Title: "Introducing Go",
						Year:  "2016",
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
						Year:  "2012",
					})

					res, err = client.Do(req, &target)
				})

				It("should respond with 201 status", func() {
					Ω(res.StatusCode).Should(Equal(http.StatusCreated))
				})

				It("should unmarshal resource object into target", func() {
					Ω(target).Should(Equal(Book{
						ID:    "1",
						Title: "An Introduction to Programming in Go",
						Year:  "2012",
					}))
				})
			})

			When("with validation errors", func() {

				BeforeEach(func() {
					req, _ := client.Post("/books/unsuccessful", Book{
						Title: "",
						Year:  "2012",
					})

					res, err = client.Do(req, &target)
				})

				It("should respond with 400 status", func() {
					Ω(res.StatusCode).Should(Equal(http.StatusBadRequest))
				})

				It("should unmarshal errors into target", func() {
					Ω(target).Should(Equal(Book{
						ValidationErrors: []*jsonapi.ErrorObject{
							{
								Title: "is required",
								Source: jsonapi.ErrorObjectSource{
									Pointer: "/data/attributes/title",
								},
							},
						},
					}))
				})
			})

			When("there is no target", func() {

				BeforeEach(func() {
					req, _ := client.Post("/books/successful", Book{
						Title: "An Introduction to Programming in Go",
						Year:  "2012",
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

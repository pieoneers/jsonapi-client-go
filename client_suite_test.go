// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client_test

import (
	"github.com/gin-gonic/gin"
	"github.com/pieoneers/jsonapi-go"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/pieoneers/jsonapi-client-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJSONAPIClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JSONAPI Client Suite")
}

var (
	ts     *httptest.Server
	client *Client
)

var _ = BeforeSuite(func() {

	InitTemplates()

	router := gin.Default()

	router.GET("/books", func(c *gin.Context) {
		body, _ := Template("books", Books{
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
		})

		c.Data(http.StatusOK, jsonapi.ContentType, body.Bytes())
	})

	router.POST("/books/successful", func(c *gin.Context) {
		body, _ := Template("book", Book{
			ID:    "1",
			Title: "An Introduction to Programming in Go",
			Year:  "2012",
		})

		c.Data(http.StatusCreated, jsonapi.ContentType, body.Bytes())
	})

	router.POST("/books/unsuccessful", func(c *gin.Context) {
		body, _ := Template("errors", []*jsonapi.ErrorObject{
			{
				Title: "is required",
				Source: jsonapi.ErrorObjectSource{
					Pointer: "/data/attributes/title",
				},
			},
		})

		c.Data(http.StatusBadRequest, jsonapi.ContentType, body.Bytes())
	})

	ts = httptest.NewServer(router)

	client = NewClient(Config{
		BaseURL: ts.URL,
	})
})

var _ = AfterSuite(func() {
	ts.Close()
})

package client_test

import (
	"testing"
  "net/http"
  "net/http/httptest"
  "github.com/pieoneers/jsonapi-go"

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

  mux := http.NewServeMux()

  mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
    body, _ := Template("books", []Book{
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
    })

    w.WriteHeader(http.StatusOK)
    w.Write(body.Bytes())
  })

  mux.HandleFunc("/books/successful", func(w http.ResponseWriter, r *http.Request) {
    body, _ := Template("book", Book{
      ID: "1",
      Title: "An Introduction to Programming in Go",
      Year: "2012",
    })

    w.WriteHeader(http.StatusCreated)
    w.Write(body.Bytes())
  })

  mux.HandleFunc("/books/unsuccessful", func(w http.ResponseWriter, r *http.Request) {
    body, _ := Template("errors", []*jsonapi.ErrorObject{
      {
        Title: "is required",
        Source: jsonapi.ErrorObjectSource{
          Pointer: "/data/attributes/title",
        },
      },
    })

    w.WriteHeader(http.StatusForbidden)
    w.Write(body.Bytes())
  })

  ts = httptest.NewServer(mux)

  client = NewClient(Config{
    BaseURL: ts.URL,
  })
})

var _ = AfterSuite(func() {
  ts.Close()
})

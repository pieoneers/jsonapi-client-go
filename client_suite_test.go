package client_test

import (
	"testing"
  "fmt"
  "net/http"
  "net/http/httptest"

  . "github.com/pieoneers/jsonapi-client-go.git"

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
  ts = httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, `
      {
        "data": {
          "type": "users",
          "id": "1",
          "attributes": {
            "email": "andrew@pieoneers.com",
            "password": "password",
            "first_name": "Andrew",
            "last_name": "Manshin"
          }
        }
      }
    `)
  }))

  ts.Start()

  client = NewClient(Config{
    BaseURL: ts.URL,
  })
})

var _ = AfterSuite(func() {
  ts.Close()
})

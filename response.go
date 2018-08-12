package client

import (
  "net/http"
  "github.com/pieoneers/jsonapi-go"
)

type Response struct {
  http.Response
  Document *jsonapi.Document
}

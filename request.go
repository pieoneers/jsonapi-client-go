package client

import (
  "io"
  "io/ioutil"
  "bytes"
  "reflect"
  "net/http"
  "github.com/pieoneers/jsonapi-go"
)

const jsonapiContentType = "application/vnd.api+json"

type Request struct {
  Method string
  URL string
  Header http.Header
  Body io.ReadCloser
}

func NewRequest(method, url string, in interface{}) (*Request, error) {
  req := Request{
    Method: method,
    URL: url,
    Header: make(http.Header),
  }

  req.Header.Add("Accept", jsonapiContentType)

  if reflect.TypeOf(in) != nil {
    req.Header.Add("Content-Type", jsonapiContentType)

    payload, marshalErr := jsonapi.Marshal(in)
    if marshalErr != nil {
      return nil, marshalErr
    }

    req.Body = ioutil.NopCloser(bytes.NewReader(payload))
  }

  return &req, nil
}

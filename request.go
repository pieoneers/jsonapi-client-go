package client

import (
  "io"
  "io/ioutil"
  "bytes"
  "reflect"
  "net/url"
  "net/http"
  "github.com/pieoneers/jsonapi-go"
)

const jsonapiContentType = "application/vnd.api+json"

type Request struct {
  Method string
  URL *url.URL
  Query url.Values
  Header http.Header
  Body io.ReadCloser
}

func NewRequest(method, rawurl string, in interface{}) (*Request, error) {
  url, urlErr := url.ParseRequestURI(rawurl)
  if urlErr != nil {
    return nil, urlErr
  }

  req := Request{
    Method: method,
    URL: url,
    Header: make(http.Header),
  }

  req.Query = req.URL.Query()

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

func(req *Request) RequestURI() string {
  req.URL.RawQuery = req.Query.Encode()
  return req.URL.RequestURI()
}

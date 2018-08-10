package client

import (
  "reflect"
  "net/http"
  "io/ioutil"
  "github.com/pieoneers/jsonapi-go"
)

type Client struct {
  httpClient *http.Client
  Config      Config
}

func NewClient(c Config) *Client {
  config := NewConfig(c)

  client := &Client{
    httpClient: &http.Client{
      Timeout: config.Timeout,
    },
    Config: config,
  }

  return client
}

func(c *Client) Get(path string) (*Request, error) {
  req, reqErr := NewRequest("GET", path, nil)
  if reqErr != nil {
    return nil, reqErr
  }

  return req, nil
}

func(c *Client) Post(path string, in interface{}) (*Request, error) {
  req, reqErr := NewRequest("POST", path, in)
  if reqErr != nil {
    return nil, reqErr
  }

  return req, nil
}

func(c *Client) Do(req *Request, out interface{}) (*Response, error) {
  baseURL    := c.Config.BaseURL
  httpClient := c.httpClient

  httpReq, reqErr := http.NewRequest(req.Method, baseURL + req.URL, req.Body)
  if reqErr != nil {
    return nil, reqErr
  }

  httpReq.Header = req.Header

  httpRes, resErr := httpClient.Do(httpReq)
  if resErr != nil {
    return nil, resErr
  }

  res := Response{
    Response: http.Response{
      StatusCode: httpRes.StatusCode,
      Header:     httpRes.Header,
      Body:       httpRes.Body,
      Request:    httpRes.Request,
    },
  }

  if res.StatusCode >= 200 && res.StatusCode <= 299 {
    if reflect.TypeOf(out) != nil {
      payload, readErr := ioutil.ReadAll(res.Body)
      if readErr != nil {
        return nil, readErr
      }

      unmarshalErr := jsonapi.Unmarshal(payload, out)
      if unmarshalErr != nil {
        return nil, unmarshalErr
      }
    }
  }

  if res.StatusCode >= 400 && res.StatusCode <= 499 {
    payload, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
      return nil, readErr
    }

    errs := []*jsonapi.ErrorObject{}

    unmarshalErr := jsonapi.Unmarshal(payload, &errs)
    if unmarshalErr != nil {
      return nil, unmarshalErr
    }

    res.Errors = errs
  }

  return &res, nil
}

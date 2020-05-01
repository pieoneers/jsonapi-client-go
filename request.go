// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"bytes"
	"github.com/pieoneers/jsonapi-go"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

//Request represents the request data
type Request struct {
  //HTTP request method
  Method string
  //Request URL or it may be a path to the resource
	URL    *url.URL
  // Query params
	Query  url.Values
  // Request headers
	Header http.Header
  // Request body
	Body   io.ReadCloser
}

// NewRequest return new Request instance with corresponding parameters
func NewRequest(method, rawurl string, in interface{}) (*Request, error) {
	url, urlErr := url.ParseRequestURI(rawurl)
	if urlErr != nil {
		return nil, urlErr
	}

	req := Request{
		Method: method,
		URL:    url,
		Header: make(http.Header),
	}

	req.Query = req.URL.Query()

	req.Header.Add("Accept", jsonapi.ContentType)

	if reflect.TypeOf(in) != nil {
		req.Header.Add("Content-Type", jsonapi.ContentType)

		payload, marshalErr := jsonapi.Marshal(in)
		if marshalErr != nil {
			return nil, marshalErr
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(payload))
	}

	return &req, nil
}

//RequestURI returns request URI 
func (req *Request) RequestURI() string {
	req.URL.RawQuery = req.Query.Encode()
	return req.URL.RequestURI()
}

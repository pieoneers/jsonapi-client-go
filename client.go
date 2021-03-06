// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"github.com/pieoneers/jsonapi-go"
	"io/ioutil"
	"net/http"
	"reflect"
)

//Client represents JSON API client and contain configuration values in Config
type Client struct {
	httpClient *http.Client
	Config     Config
}

//NewClient method returns new Client instance
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

//Get returns GET ready Request
func (c *Client) Get(path string) (*Request, error) {
	req, reqErr := NewRequest("GET", path, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	return req, nil
}

//Head returns HEAD ready Request
func (c *Client) Head(path string) (*Request, error) {
	req, reqErr := NewRequest("HEAD", path, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	return req, nil
}

//Post returns POST ready Request
func (c *Client) Post(path string, in interface{}) (*Request, error) {
	req, reqErr := NewRequest("POST", path, in)
	if reqErr != nil {
		return nil, reqErr
	}

	return req, nil
}

//Do proceeds the provided Request
func (c *Client) Do(req *Request, out interface{}) (*Response, error) {
	baseURL := c.Config.BaseURL
	httpClient := c.httpClient

	httpReq, reqErr := http.NewRequest(req.Method, baseURL+req.RequestURI(), req.Body)
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

	payload, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	if len(payload) > 0 && reflect.TypeOf(out) != nil {
		document, unmarshalErr := jsonapi.Unmarshal(payload, out)
		if unmarshalErr != nil {
			return nil, unmarshalErr
		}

		res.Document = document
	}

	return &res, nil
}

// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"github.com/pieoneers/jsonapi-go"
	"net/http"
)

// Response extends http.Response with jsonapi.Document.
//
// jsonapi.Document contatin follow fields
// See more in jsonapi-go package documentation
//
// type Document struct {
//   Data     *documentData     `json:"data,omitempty"`
//   Errors   []*ErrorObject    `json:"errors,omitempty"`
//   Included []*ResourceObject `json:"included,omitempty"`
//   Meta     json.RawMessage   `json:"meta,omitempty"`
// }
type Response struct {
	http.Response
	Document *jsonapi.Document
}

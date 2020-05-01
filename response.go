// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"github.com/pieoneers/jsonapi-go"
	"net/http"
)

type Response struct {
	http.Response
	Document *jsonapi.Document
}

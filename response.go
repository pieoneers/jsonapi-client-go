package client

import (
	"github.com/pieoneers/jsonapi-go"
	"net/http"
)

type Response struct {
	http.Response
	Document *jsonapi.Document
}

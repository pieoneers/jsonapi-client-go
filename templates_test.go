// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client_test

import (
	"bytes"
	"fmt"
	"text/template"
)

var templates map[string]*template.Template

func Template(name string, data interface{}) (*bytes.Buffer, error) {
	var err error

	buf := &bytes.Buffer{}

	if t, ok := templates[name]; ok {
		err = t.Execute(buf, data)
	} else {
		err = fmt.Errorf("Template %v was not found", name)
	}

	return buf, err
}

func InitTemplates() {

	errorsDocumentString := `
    {
      "errors": [
        {{range $index, $element := .}}
          {{if $index}},{{end}}
          {
            "title": "{{.Title}}",
            "source": {
              "pointer": "{{.Source.Pointer}}"
            }
          }
        {{end}}
      ]
    }
  `

	booksDocumentString := `
    {
      "data": [
        {{range $index, $element := .}}
          {{if $index}},{{end}}
          {{template "book" $element}}
        {{end}}
      ]
    }
  `

	bookDocumentString := `
    {
      "data": {{template "book" .}}
    }
  `

	bookPayloadDocumentString := `
    {
      "data": {
        "type": "books",
        "attributes": {{template "book-attributes" .}}
      }
    }
  `

	bookString := `
    {{define "book"}}
      {
        "type": "books",
        "id": "{{.ID}}",
        "attributes": {{template "book-attributes" .}}
      }
    {{end}}
  `

	bookAttributesString := `
    {{define "book-attributes"}}
      {
        "title": "{{.Title}}",
        "year": "{{.Year}}"
      }
    {{end}}
  `

	errorsDocumentTemplate := template.Must(template.New("errors").Parse(errorsDocumentString))

	booksDocumentTemplate := template.Must(template.New("books-document").Parse(booksDocumentString))
	booksDocumentTemplate = template.Must(booksDocumentTemplate.Parse(bookString))
	booksDocumentTemplate = template.Must(booksDocumentTemplate.Parse(bookAttributesString))

	bookDocumentTemplate := template.Must(template.New("book-document").Parse(bookDocumentString))
	bookDocumentTemplate = template.Must(bookDocumentTemplate.Parse(bookString))
	bookDocumentTemplate = template.Must(bookDocumentTemplate.Parse(bookAttributesString))

	bookPayloadDocumentTemplate := template.Must(template.New("book-payload-document").Parse(bookPayloadDocumentString))
	bookPayloadDocumentTemplate = template.Must(bookPayloadDocumentTemplate.Parse(bookAttributesString))

	templates = map[string]*template.Template{
		"errors":       errorsDocumentTemplate,
		"books":        booksDocumentTemplate,
		"book":         bookDocumentTemplate,
		"book-payload": bookPayloadDocumentTemplate,
	}
}

package client_test

import (
  "github.com/pieoneers/jsonapi-go"
)

type Book struct {
  ID               string                 `json:"-"`
  Title            string                 `json:"title"`
  Year             string                 `json:"year"`
  ValidationErrors []*jsonapi.ErrorObject `json:"-"`
}

func(b Book) GetID() string {
  return b.ID
}

func(b Book) GetType() string {
  return "books"
}

func(b Book) SetType(string) error {
  return nil
}

func(b Book) GetData() interface{} {
  return b
}

func(b *Book) SetID(id string) error {
  b.ID = id
  return nil
}

func(b *Book) SetData(to func(target interface{}) error) error {
  return to(b)
}

func(b *Book) SetErrors(errors []*jsonapi.ErrorObject) error {
  b.ValidationErrors = errors
  return nil
}

type Books []Book

func(b *Books) SetData(to func(target interface{}) error) error {
  return to(b)
}

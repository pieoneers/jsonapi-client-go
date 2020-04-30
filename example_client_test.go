// Copyright (c) 2020 Pieoneers Software Incorporated. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client_test

import (
	"fmt"
	jsonapiClient "github.com/pieoneers/jsonapi-client-go"
	"log"
	"time"
)

type Item struct {
	ID        string    `json:"-"`
	Type      string    `json:"-"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func (d *Item) SetID(id string) error {
	d.ID = id
	return nil
}

func (d *Item) SetType(t string) error {
	d.Type = t
	return nil
}

func (b *Item) SetData(to func(target interface{}) error) error {
	return to(b)
}

type Items []Item

func (b *Items) SetData(to func(target interface{}) error) error {
	return to(b)
}

func ExampleClient() {
	var target []Item

	config := jsonapiClient.Config{
		BaseURL: "http://localhost",
		Timeout: time.Second,
	}

	client := jsonapiClient.NewClient(config)

	request, requestErr := jsonapiClient.NewRequest("GET", "/data", nil)

	if requestErr != nil {
		log.Println("requestErr: ", requestErr)
		return
	}

	request.Query.Set("filter[color]", "silver")

	_, responseErr := client.Do(request, &target)

	if responseErr != nil {
		log.Println("responseErr: ", responseErr)
		return
	}

	for _, item := range target {
		fmt.Printf("\titem#%v:\t%+v\n", item.ID, item)
	}
}

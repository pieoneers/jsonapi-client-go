# Go Jsonapi client
This JSON API client library for Go.

# Usage

``` go
package main

import(
  "log"
  "time"
  "github.com/pieoneers/jsonapi-client-go"
)

//Data structure of example data from a server
type Item struct{
  ID string `json:"-"`
  Type string `json:"-"`
  Title string `json:"title"`
  CreatedAt time.Time `json:"created_at"`
}

// jsonapi-go library require a few methods for data structures to be able unmarshal data from json api document.
//
//   SetID(string) error
//   SetType(string) error
//   SetData(func(interface{}) error) error

func(d *Item) SetID(id string) error {
  d.ID = id
  return nil
}

func(d *Item) SetType(t string) error {
  d.Type = t
  return nil
}

func(b *Item) SetData(to func(target interface{}) error) error {
  return to(b)
}

type Items []Item

// If the response will contain collections, it should be wrapped by data type and method SetData should me implemented for the collection data type
func(b *Items) SetData(to func(target interface{}) error) error {
  return to(b)
}

func main() {
  var target Items

  config := client.Config{
    BaseURL: "http://json-api-server.com", //The default value is "http://localhost"
    Timeout: time.Second,  //The default value is 10 seconds(time.Second * 10)
  }

  jsonapiClient := client.NewClient(config) //Creates a new client instance with the config

  request, requestErr := client.NewRequest("GET", "/items", nil) //Creates a new request

  if requestErr != nil {
    log.Println(requestErr)
    return
  }

  request.Query.Set("filter[color]", "red") //Query is an url.Values attribute of Request, it used to pass URL query params

  response, responseErr := jsonapiClient.Do(request, &target) //Proceed the request

  if responseErr != nil {
    log.Println(responseErr)
    return
  }
  //Proceed the data(target) here

}
```

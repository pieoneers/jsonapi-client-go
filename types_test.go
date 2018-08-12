package client_test

type Book struct {
  ID    string `json:"-"`
  Title string `json:"title"`
  Year  string `json:"year"`
}

func(b Book) GetID() string {
  return b.ID
}

func(b Book) GetType() string {
  return "books"
}

func(b *Book) SetID(id string) error {
  b.ID = id
  return nil
}

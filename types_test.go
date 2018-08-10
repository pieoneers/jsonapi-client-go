package client_test

type User struct {
  ID        string `json:"-"`
  Email     string `json:"email"`
  Password  string `json:"password"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
}

func(u User) GetID() string {
  return u.ID
}

func(u User) GetType() string {
  return "users"
}

func(u *User) SetID(id string) error {
  u.ID = id
  return nil
}

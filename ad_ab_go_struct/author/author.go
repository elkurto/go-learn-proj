package author

import (
  "fmt"
  "time"
)

type Author struct {
  Id        int
  Name      string
  CreatedAt time.Time
}

func NewAuthor(id int, name string) *Author {
  return &Author{
    Id:        id,
    Name:      name,
    CreatedAt: time.Now(),
  }
}

func (author *Author) ToString() string {
  return fmt.Sprintf("author.Author:{Id:%d,Name:%s,CreatedAt:%s",
    author.Id, author.Name, author.CreatedAt)
}

func (author *Author) Println() {
  fmt.Printf("%s\n", author.toString())
}

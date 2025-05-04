package author

import (
  "fmt"
  "time"
)

type Author struct {
  Id        int
  Name      string
  CreatedAt int64
}

func NewAuthor(id int, name string) *Author {
  return &Author{
    Id:        id,
    Name:      name,
    CreatedAt: time.Now().Unix(),
  }
}

func (author *Author) ToString() string {
  return fmt.Sprintf("author.Author:{Id:%d,Name:%s,CreatedAt:%d}",
    author.Id, author.Name, author.CreatedAt)
}

func (author *Author) Println() {
  fmt.Printf("%s\n", author.ToString())
}

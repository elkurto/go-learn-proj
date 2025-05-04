package person

import "fmt"

type Person struct {
	Id        int64
	Name      string
	BirthDate int64 // unix epoch time of birthday
}

func New(id int64, name string, birthDate int64) *Person {
	return &Person{
		Id:        id,
		Name:      name,
		BirthDate: birthDate,
	}
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Person {id:%d, Name:%s}", p.Id, p.Name)
}

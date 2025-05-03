package main

import (
  "fmt"
  "time"
)

// 1. create a book struct
type Book struct {
  id        int
  title     string
  author    string
  isbn      string
  createdAt time.Time
}

// 2. define a method on the :struct:Book
func (book Book) printDetails() {
  fmt.Printf(
    `Book: {id:%d, %s, %s }
`, book.id, book.title, book.author)
}

func main() {
  fmt.Println("\nRunning ad_ab_go_struct/main.go:")

  // 3. create a book instance (using a struct composite literal)
  var book00 = Book{
    id:        1,
    title:     "Pro Go",
    author:    "Freeman",
    isbn:      "00-1234-12345678-90123456-0001",
    createdAt: time.Now(),
  }

  // 3. print book details via function
  fmt.Println("printDetails(book00) // a proper func")
  printBookDetails(book00)

  // 4. print book details via method of :struct:Book
  fmt.Println("\nbook00.printDetails() // a struct method // note: encapsulation")
  book00.printDetails()
}

func printBookDetails(book Book) {
  fmt.Printf(`Book: {id:%d, %s, %s }
`, book.id, book.title, book.author)
}

package main

import (
	"abc.com/ad_ab_go_struct/author"
	"fmt"
	"time"
)

// Book 1. create a book struct
type Book struct {
	id        int
	title     string
	author    string
	isbn      string
	createdAt time.Time
}

func newBook(id int, title, author, isbn string) (*Book, error) {
	var book = Book{}

	book.id = id
	book.title = title
	book.author = author
	book.isbn = isbn
	book.createdAt = time.Now()

	return &book, nil
}

// 2. define a method on :struct:Book
// note: Should use :type:"*Book" to pass by ref -- and avoid pass by copy :type:Book
func (book *Book) printDetails() {
	fmt.Printf(
		`Book: {id:%d, title:%s, author:%s }
`, book.id, book.title, book.author)
}

// note: Should use :type:"*Book" to pass by ref -- and avoid pass by copy :type:Book
func (book *Book) mutateTitle(newTitle string) {
	book.title = newTitle
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
	fmt.Println("ex01: printDetails(book00) // a proper func")
	printBookDetails(&book00) // pass address of book to avoid "pass by copy"

	// 4. print book details via method of :struct:Book
	fmt.Println("\nex02: book00.printDetails() // a struct method // note: encapsulation")
	book00.printDetails()

	// 5.
	fmt.Println("\nex03: book00.mutateTitle(\"Pro Go!!!!!\") // a struct mutator-method ")
	fmt.Println("book00.printDetails() // print state after mutating :member:title")
	book00.mutateTitle("Pro Go!!!!!")
	book00.printDetails()

	// 6. use factoryFunction to create Book
	fmt.Println("\nex04: var book01, err = newBook(2, \"The Art Of WebAssembly\", \"Battagline\", \"00-123123-123123-99\")")
	fmt.Println("  use a factoryFunction Book::newBook as an ersatz constructor")
	var book01, err = newBook(2, "The Art Of WebAssembly", "Battagline", "00-123123-123123-99")
	if err != nil {
		fmt.Printf("Error: failed to create book %s\n", err)
	} else if book01 == nil {
		fmt.Printf("Error: Book instance was null\n")
	} else {
		book01.printDetails()
	}

	// 7. create, import, print an :struct:author
	fmt.Println("\nex06: import, create, print :struct:author.Author\n",
		"  var author05 = author.NewAuthor(5, \"Ian Fleming\")")
	fmt.Println("  author.Println()")
	var author05 = author.NewAuthor(5, "Ian Fleming")
	author05.Println()
}

func printBookDetails(book *Book) {
	fmt.Printf(`Book: {id:%d, %s, %s }
`, book.id, book.title, book.author)
}

/**
USAGE:
cd ad_ab_go_struct/
go run main.go

// OUTPUT:
Running ad_ab_go_struct/main.go:
ex01: printDetails(book00) // a proper func
Book: {id:1, Pro Go, Freeman }

ex02: book00.printDetails() // a struct method // note: encapsulation
Book: {id:1, title:Pro Go, author:Freeman }

ex03: book00.mutateTitle("Pro Go!!!!!") // a struct mutator-method
book00.printDetails() // print state after mutating :member:title
Book: {id:1, title:Pro Go!!!!!, author:Freeman }

ex04: var book01, err = newBook(2, "The Art Of WebAssembly", "Battagline", "00-123123-123123-99")
  use a factoryFunction Book::newBook as an ersatz constructor
Book: {id:2, title:The Art Of WebAssembly, author:Battagline }



*/

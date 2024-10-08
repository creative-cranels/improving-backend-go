package main

import (
	"errors"
	"fmt"
)

type BookNotAvailableError struct {
	BookID uint
}

func (e BookNotAvailableError) Error() string {
	return fmt.Sprintf("Book %d is not available for lending", e.BookID)
}

type Book struct {
	ID           uint
	Title        string
	Author       string
	Availability bool
}

func (b *Book) Lend() (bool, error) {
	switch b.Availability {
	case false:
		return false, BookNotAvailableError{BookID: b.ID}
	default:
		b.Availability = false
		return true, nil
	}
}

func (b *Book) Return() error {
	switch b.Availability {
	case false:
		b.Availability = true
		return nil
	default:
		return errors.New("book already returned")
	}
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) error {
	l.Books = append(l.Books, Book{
		ID:           book.ID,
		Title:        book.Title,
		Author:       book.Author,
		Availability: true,
	})
	return nil
}

func main() {
	lib := Library{}
	lib.AddBook(Book{
		ID:     1,
		Title:  "Dune",
		Author: "Frank Herbert",
	})

	lib.AddBook(Book{
		ID:     2,
		Title:  "Dune Messiah",
		Author: "Frank Herbert",
	})

	lib.Books[0].Lend()
	if err := lib.Books[1].Return(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("lib books:", lib.Books)
}

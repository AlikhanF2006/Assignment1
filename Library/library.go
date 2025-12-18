package library

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() Library {
	return Library{
		Books: make(map[string]Book),
	}
}

func (l Library) AddBook(id, title, author string) Library {
	l.Books[id] = Book{
		ID:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
	return l
}

func (l Library) BorrowBook(id string) Library {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book not found")
		return l
	}

	if book.IsBorrowed {
		fmt.Println("Book already borrowed")
		return l
	}

	book.IsBorrowed = true
	l.Books[id] = book
	fmt.Println("Book borrowed successfully")
	return l
}

func (l Library) ReturnBook(id string) Library {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book not found")
		return l
	}

	book.IsBorrowed = false
	l.Books[id] = book
	fmt.Println("Book returned successfully")
	return l
}

func (l Library) ListAvailableBooks() {
	fmt.Println("Available books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("- %s by %s\n", book.Title, book.Author)
		}
	}
}

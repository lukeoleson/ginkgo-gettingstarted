package main

import (
	"fmt"
	"lukeoleson/ginko-gettingstarted/books"
)

func main() {
	book := books.GetBook("The Gunslinger")
	fmt.Printf("Title: %v\nAuthor: %v", book.Title, book.Author)
}

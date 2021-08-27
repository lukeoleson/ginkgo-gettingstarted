package books

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

var books = map[string]Book{
	"War & Peace": {
		Title:  "War & Peace",
		Author: "Leo Tolstoy",
	},
	"The Gunslinger": {
		Title:  "The Gunslinger",
		Author: "Stephen King",
	},
}

func GetBook(title string) Book {
	return books[title]
}

func (b Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}

func NewBookFromJSON(jsonBook string) Book {

	var book Book
	err := json.Unmarshal([]byte(jsonBook), &book)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return book
}

package books

type Book struct {
	Title string
	Author string
}

var books = map[string]Book {
	"War & Peace": {
		Title: "War & Peace",
		Author: "Leo Tolstoy",
	},
	"The Gunslinger": {
		Title: "The Gunslinger",
		Author: "Stephen King",
	},
}

func GetBook(title string) Book {
	return books[title]
}
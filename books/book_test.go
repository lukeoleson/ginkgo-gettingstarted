package books_test

import (
	"lukeoleson/ginko-gettingstarted/books"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() { // Container to group a set of tests (test data, test cases, ...)
	var (
		longBook  books.Book
		shortBook books.Book
	)

	BeforeEach(func() { // Runs before each It in scope. Do not initialize data anywhere else as it could get modified by tests.
		longBook = books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		shortBook = books.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() { // You can continue to group as you see fit

		BeforeEach(func() { // Runs before each It in scope and after the BeforeEach above (outermost to innermost)
			longBook.Pages = 2784
		})

		Context("With more than 300 pages", func() { // Adds context to a specific test (or series of tests)
			It("should be a novel", func() { // An actual test - referred to as a "spec"in Ginkgo
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL")) // Assertions
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

	Describe("Parse", func() {
		Context("From JSON", func() {
			It("can be loaded from JSON", func() {
				book := books.NewBookFromJSON(`{
					"title":"Les Miserables",
					"author":"Victor Hugo",
					"pages":2783
				}`)

				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hug"))
				Expect(book.Pages).To(Equal(2783))
			})
		})

		Context("From JSON", func() {
			Specify("a parsing error", func() { // Alias for "It" to make the code more readable/grammatically correct.
				book := books.NewBookFromJSON(`{
					"name":"Les Miserables",
				}`)

				Expect(book.Title).To(Equal("Les Miserables"))
			})
		})

	})
})

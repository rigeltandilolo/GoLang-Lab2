package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func displayBook(b Book) {
	fmt.Printf("Title: %s\nAuthor: %s\nISBN: %s\n", b.Title, b.Author, b.ISBN)
}

func validateISBN(isbn string) error {
	if len(isbn) != 13 {
		return errors.New("ISBN must be 13 characters long")
	}
	return nil
}

func createBook() (Book, error) {
	var title, author, isbn string

	fmt.Print("Enter Book Title: ")
	fmt.Scanln(&title)

	fmt.Print("Enter Author: ")
	fmt.Scanln(&author)

	fmt.Print("Enter ISBN: ")
	fmt.Scanln(&isbn)

	err := validateISBN(isbn)
	if err != nil {
		return Book{}, err
	}

	return Book{Title: title, Author: author, ISBN: isbn}, nil
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	var library []Book

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"library": library})
	})

	router.POST("/add", func(c *gin.Context) {
		book, err := createBook()
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": err.Error()})
			return
		}
		library = append(library, book)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}

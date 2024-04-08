package main

import (
	"fmt"
	"os"

	"github.com/giovannymassuia/kindle-highlight-sync/internal/books"
	"github.com/giovannymassuia/kindle-highlight-sync/internal/http"
)

func main() {
	xMain := os.Getenv("X_MAIN")
	ubidMain := os.Getenv("UBID_MAIN")
	atMain := os.Getenv("AT_MAIN")

	httpClient := http.NewClient(http.ClientProps{
		Cookies: map[string]string{
			"ubid-main": ubidMain,
			"x-main":    xMain,
			"at-main":   atMain,
		},
	})

	booksResult, err := books.GetBooks(httpClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, book := range booksResult {
		// book id and title
		fmt.Printf("Book ID: %s, Title: %s\n", book.ID, book.Title)
	}

	// highlights
	highlight, err := books.GetHighlightsAndNotes(httpClient, booksResult[0].ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, h := range highlight {
		fmt.Printf("Highlight ID: %s,\nHighlight Text: %s\n", h.ID, h.Text)
	}
}

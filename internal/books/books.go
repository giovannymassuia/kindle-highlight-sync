package books

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/giovannymassuia/kindle-highlight-sync/internal/http"
)

const (
	BOOKS_URL = "https://read.amazon.com/notebook"
	NOTES_URL = "https://read.amazon.com/notebook?asin=%s&contentLimitState=&"
)

type Book struct {
	ID    string
	Title string
	Image string
}

type Highlight struct {
	ID   string
	Text string
}

func GetBooks(httpClient *http.Client) ([]Book, error) {
	booksDoc, err := httpClient.FetchHTML(BOOKS_URL)
	if err != nil {
		return nil, err
	}

	return extractBooks(booksDoc), nil
}

func extractBooks(doc *goquery.Document) []Book {
	var books []Book

	doc.Find(".kp-notebook-library-each-book").Each(func(i int, s *goquery.Selection) {
		bookID, _ := s.Attr("id")
		title := s.Find("h2").Text()
		imgUrl, _ := s.Find("img").Attr("src")

		books = append(books, Book{ID: bookID, Title: title, Image: imgUrl})
	})

	return books
}

func GetHighlightsAndNotes(httpClient *http.Client, bookID string) ([]Highlight, error) {
	notesURL := buildNotesURL(bookID)
	notesDoc, err := httpClient.FetchHTML(notesURL)
	if err != nil {
		return nil, err
	}

	return extractHighlightsAndNotes(notesDoc), nil
}

func buildNotesURL(bookID string) string {
	return fmt.Sprintf(NOTES_URL, bookID)
}

func extractHighlightsAndNotes(doc *goquery.Document) []Highlight {
	var highlights []Highlight

	doc.Find(".kp-notebook-highlight").Each(func(index int, item *goquery.Selection) {
		divID, exists := item.Attr("id")
		if exists {
			text := item.Text()
			text = strings.TrimSpace(text)
			highlights = append(highlights, Highlight{ID: divID, Text: text})
		}
	})

	return highlights
}

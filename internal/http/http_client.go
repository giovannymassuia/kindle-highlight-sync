package http

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ClientProps struct {
	Cookies map[string]string
}

type Client struct {
	Props ClientProps
}

func NewClient(props ClientProps) *Client {
	return &Client{Props: props}
}

func (c *Client) FetchHTML(url string) (*goquery.Document, error) {
	client := &http.Client{}

	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return nil, error
	}

	for key, value := range c.Props.Cookies {
		req.AddCookie(&http.Cookie{Name: key, Value: value})
	}

	resp, error := client.Do(req)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	doc, error := goquery.NewDocumentFromReader(resp.Body)
	if error != nil {
		return nil, error
	}

	return doc, nil
}

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func (c *Client) GetAllBooks() ([]Book, error) {
	resp, err := http.Get(fmt.Sprintf("%s/books", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var books []Book
	if err := json.NewDecoder(resp.Body).Decode(&books); err != nil {
		return nil, err
	}

	return books, nil
}

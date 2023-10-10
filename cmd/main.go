package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func jsonResponse(c *gin.Context, status string, message string, data interface{}) {
	response := Response{
		Message: message,
		Status:  status,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

func GetAllBooks(c *gin.Context) {

	localhostAPIEndpoint := "http://localhost:8080/books"

	// Make a request to the localhost API
	response, err := http.Get(localhostAPIEndpoint)
	if err != nil {
		jsonResponse(c, "error", "Failed to fetch data from localhost API", nil)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		jsonResponse(c, "error", "Failed to fetch data from localhost API", nil)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		jsonResponse(c, "error", "Failed to read response from localhost API", nil)
		return
	}

	// Parse the data as JSON array
	var books []Book
	if err := json.Unmarshal(body, &books); err != nil {
		jsonResponse(c, "error", "Failed to unmarshal data from localhost API", nil)
		return
	}

	// Serve the data from the localhost API through your API
	jsonResponse(c, "success", "Books retrieved successfully", books)
}

func main() {
	router := gin.Default()
	router.GET("/books", GetAllBooks)
	router.Run("localhost:8081")
}

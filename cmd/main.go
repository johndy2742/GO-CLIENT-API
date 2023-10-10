package main

import (
	"encoding/json"
	"fmt"
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

func GetAllBooks(c *gin.Context) {
	// Replace with the actual localhost API endpoint
	localhostAPIEndpoint := "http://localhost:8080/books"

	// Make a request to the localhost API
	response, err := http.Get(localhostAPIEndpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch data from localhost API",
		})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch data from localhost API",
		})
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to read response from localhost API",
		})
		return
	}

	// Parse the data as JSON array
	var books []Book
	if err := json.Unmarshal([]byte(body), &books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to unmarshal data from localhost API",
		})
		return
	}

	// Serve the data from the localhost API through your API
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Books retrieved successfully",
		"data":    books,
	})
}

func main() {
	router := gin.Default()

	// Define the route for accessing the localhost API
	router.GET("/books", GetAllBooks)

	fmt.Println("Server running at :8081")
	router.Run("localhost:8081")
}

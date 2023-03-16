package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

func main() {
	router := gin.Default()
	router.GET("/", getBooks)
	router.GET("/first", getBooksFirst)
	router.Run("localhost:8080")
	for _, value := range books {
		fmt.Println(value)
	}
}
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func getBooksFirst(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books[0])
}

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

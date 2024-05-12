package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

var items = []item{{ID: 1, Content: "Some content here", Url: "www.ccpastebin.com/1"},
	{ID: 2, Content: "Some more content here", Url: "www.ccpastebin.com/2"}}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func postItems(c *gin.Context) {
	var newItem item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func getItemById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range items {
		newId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		if a.ID == newId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.POST("/items", postItems)
	router.GET("/items/:id", getItemById)

	router.Run("localhost:8000")
}

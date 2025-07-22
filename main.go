package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"www.example.com/met/handlers"
)

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./public")

	// Home page route
	router.GET("/", func(c *gin.Context) {
		handlers.Home(c.Writer, c.Request)
	})

	router.GET("/met", func(c *gin.Context) {
		fmt.Println("hello World!")
		handlers.DepartmentList(c.Writer, c.Request)
	})

	router.GET("/collections", func(c *gin.Context) {
		fmt.Println("getting collection objects")
		handlers.CollectionObjectList(c.Writer, c.Request)

	})

	router.GET("/collection/:id", handlers.GetCollectionObjIDItem)

	router.Run(":8085")

}

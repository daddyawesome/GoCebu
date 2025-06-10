package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Gin Framework!")
	})

	r.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"title":   "About Page",
			"message": "This is built with Gin!",
		})
	})

	r.GET("/contact", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"email": "contact@example.com",
			"phone": "+123456789",
		})
	})

	r.Run(":8080")
}

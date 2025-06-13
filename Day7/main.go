package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Feedback struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message" binding:"required"`
}

var feedbacks []Feedback

func main() {
	r := gin.Default()

	r.POST("/feedback", func(c *gin.Context) {
		var fb Feedback
		if err := c.ShouldBindJSON(&fb); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		feedbacks = append(feedbacks, fb)
		c.JSON(http.StatusOK, gin.H{"status": "received", "count": len(feedbacks)})
	})

	r.GET("/feedbacks", func(c *gin.Context) {
		c.JSON(http.StatusOK, feedbacks)
	})

	r.Run(":8080")
}

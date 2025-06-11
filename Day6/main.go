package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/weather/:city", func(c *gin.Context) {
		city := c.Param("city")
		unit := c.DefaultQuery("unit", "celsius")

		response := fmt.Sprintf("Weather for %s in %s", city, unit)
		c.JSON(http.StatusOK, gin.H{"forecast": response})
	})

	r.Run(":8080")
}

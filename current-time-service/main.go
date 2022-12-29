package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/current-time", currentTime)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func currentTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": time.Now(),
	})
}

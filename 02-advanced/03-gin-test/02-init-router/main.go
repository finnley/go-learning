package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pong(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ao(c *gin.Context)  {
	c.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ao", ao)
	r.GET("/pong", pong)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8083") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
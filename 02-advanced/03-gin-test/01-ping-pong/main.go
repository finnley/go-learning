package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pong(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "pong",
	//})

	//c.JSON(http.StatusOK, map[string]string{
	//	"message": "pong",
	//})

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}

// 参考: https://github.com/gin-gonic/gin
// go get -u github.com/gin-gonic/gin
func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.GET("/ping", pong)
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8083")
}

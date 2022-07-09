package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义中间件
func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "123456")
		// 让原本该执行的逻辑继续执行
		c.Next()

		end := time.Since(t)
		fmt.Printf("耗时: %v\n", end)

		status := c.Writer.Status()
		fmt.Println("状态", status)
	}
}

func main() {
	//router := gin.New()
	////router.Use(gin.Logger())
	////router.Use(gin.Recovery())
	//router.Use(gin.Logger(), gin.Recovery())

	//authorized := router.Group("/goods")
	//authorized.Use(AuthRequired)

	router := gin.Default()
	router.Use(MyLogger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	_ = router.Run(":8083")
}

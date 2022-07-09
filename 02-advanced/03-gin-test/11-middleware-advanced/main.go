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
		//return // return 后续逻辑还能继续执行

		// 让原本该执行的逻辑继续执行
		c.Next()

		end := time.Since(t)
		fmt.Printf("耗时: %v\n", end)

		status := c.Writer.Status()
		fmt.Println("状态", status)
	}
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			} else {
				fmt.Println(k, v)
			}
			fmt.Println(k, v, token)
		}

		if token != "finnley" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			//return // return 没有起到做，如果要起作用，应该是用 abort
			// 为什么return阻止不了后续逻辑额执行?
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	// 执行后面逻辑之前判断下token
	router.Use(TokenRequired())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	_ = router.Run(":8083")
}

func AuthRequired(c *gin.Context) {

}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", welcome)
	router.POST("/form_post", formPost)
	router.POST("/post", getPost)
	router.Run(":8083")
}

// 127.0.0.1:8083/post?id=111&page=20
// body: {"name": "aaa", "message": "bbb"}
func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name") // 从post body中取出参数
	message := c.DefaultPostForm("message", "信息")

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func formPost(c *gin.Context) {
	// 支持post form 表单，不支持json
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anmonymous")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

// http://127.0.0.1:8083/welcome?firstname=aaa&lastname=bbb
func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "dog")
	lastName := c.DefaultQuery("lastname", "cat")
	//lastName := c.Query("lastname") // query不需要默认值，defaultQuery需要默认值
	c.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	//ID   string `uri:"id" binding:"required,uuid"`
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

// 路由分组
func main() {
	router := gin.Default()

	//goodsGroup := router.Group("/goods")
	//{
	//	goodsGroup.GET("", goodsList) // 商品列表
	//	//goodsGroup.GET("/:id", goodsDetail) // 商品信息
	//	//goodsGroup.GET("/:id/:action", goodsDetail) // 商品信息
	//	goodsGroup.GET("/:id/*action", goodsDetail) // 商品信息
	//	goodsGroup.POST("", createGoods)            // 添加商品
	//}

	// ShouldBindUri
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.Status(404)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": person.Name,
			"id":   person.ID,
		})
	})
	router.Run(":8083")
}

func createGoods(context *gin.Context) {

}

func goodsDetail(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "goodsList",
	})
}

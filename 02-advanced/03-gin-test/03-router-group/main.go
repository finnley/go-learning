package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由分组
func main() {
	router := gin.Default()
	//router.GET("/goods/list", goodsList)  // 商品列表
	//router.GET("/goods/1", goodsDetail)   // 商品信息
	//router.GET("/goods/add", createGoods) // 添加商品
	// 路由分组 以 goods 开头
	//goodsGroup := router.Group("/goods")
	//goodsGroup.GET("/list", goodsList)   // 商品列表
	//goodsGroup.GET("/1", goodsDetail)    // 商品信息
	//goodsGroup.POST("/add", goodsDetail) // 添加商品

	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("/list", goodsList)   // 商品列表
		goodsGroup.GET("/1", goodsDetail)    // 商品信息
		goodsGroup.POST("/add", goodsDetail) // 添加商品
	}

	router.Run(":8083")
}

func createGoods(context *gin.Context) {

}

func goodsDetail(c *gin.Context) {

}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "goodsList",
	})
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	router := gin.Default()

	// 获取当前路径
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// /private/var/folders/jc/z9lpk0752lv54fs9fprh_vj80000gn/T/GoLand 这是一个临时路径
	fmt.Println(dir)

	// 静态资源 第一个参数是只要url中以/static开头,就去第二个参数当前目录下的static去找
	router.Static("/static", "./static")

	// 这个方法会将指定目录下文件加载好，相对目录，需要在终端进入指定目录启动
	//router.LoadHTMLFiles("/Users/finnley/Coding/go-learning/02-advanced/03-gin-test/12-html/templates/index.tmpl")
	//router.LoadHTMLFiles("templates/index.tmpl", "templates/goods.html")
	//router.LoadHTMLGlob("templates/*") // 加载 templates 下所有文件
	router.LoadHTMLGlob("templates/**/*") // 加载 templates 下所有目录文件 找的是二级目录下的，一级目录没有找

	// 如果没有在模板中使用 define 定义，就可以使用默认的文件名查找
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "gin template",
		})
	})

	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"name": "Goods",
		})
	})

	router.GET("/goods/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/list.html", gin.H{
			"name": "GoodsList",
		})
	})

	router.GET("/users/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/list.html", gin.H{
			"name": "UsersList",
		})
	})

	router.Run(":8083")
}

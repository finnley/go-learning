package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-learning/02-advanced/03-gin-test/06-json-protobuf/proto"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf", returnProto)
	router.Run(":8083")
}

func returnProto(c *gin.Context) {
	course := []string{"python", "go", "php"}
	user := &proto.Teacher{
		Name:   "张三",
		Course: course,
	}
	c.ProtoBuf(http.StatusOK, user)
}

func moreJSON(c *gin.Context) {
	var msg struct {
		Name string `json:"name"`
		//Name    string `json2struct:"admin"`
		Message string
		Number  int
	}
	msg.Name = "finnley"
	msg.Message = "json测试"
	msg.Number = 20

	c.JSON(http.StatusOK, msg)
}

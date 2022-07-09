package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

/**
1. 传输协议: http
 */
func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
	// 读取返回内容
	body, _ := res.Body()
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

func main() {
	req := HttpRequest.NewRequest()
	res, _ := req.Get("http://127.0.0.1:8000/add?a=1&b=2")
	// 读取返回内容
	body, _ := res.Body()
	fmt.Println(string(body)) // {"data":3}

	// 上面写法封装如下：
	fmt.Println(Add(1, 2)) // 3
	fmt.Println(Add(4, 3)) // 7
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// 通过 http 实现 add 服务端功能
func main() {
	// 将 add 放到远程服务上，这里使用 path 来表示 call id，比如：http://127.0.0.1:8000/add?a=1&b=2
	// response: json {"data": 3}
	// 下面一个简单的方法就已经包含了远程过程调用需要解决的三个问题：
	// 1. call id: r.URL.Path，对于http请求，每个url都可以对应一个方法的
	// 2. 数据的传输协议: 传参使用 URL参数传递，返回使用 json协议
	// 3. 网络传输协议: http协议
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// 获取前端传过来的参数并解析
		_ = r.ParseForm()
		// 打印请求路径
		fmt.Println("path: ", r.URL.Path)
		// 从URL中解析参数
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		// 返回
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(jData)
	})

	// 监听
	_ = http.ListenAndServe(":8000", nil)
}

/**
客户端的一种请求方式直接使用浏览器请求:
http://127.0.0.1:8000/add?a=1&b=2
{
    "data": 3
}

另一种方式通过编写 client 客户端方式实现远程过程访问
*/

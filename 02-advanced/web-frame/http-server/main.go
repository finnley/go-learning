package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUp2(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "dederialized failed: %v", err)
		return
	}

	// 返回一个虚拟的 user id 表示注册成功
	//fmt.Fprintf(w, "%d", 123)

	resp := &commonResponse{
		Data: 123,
	}
	respJson, err := json.Marshal(resp)
	if err != nil {

	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(respJson))
}

/**
上面代码缺点，还没有写业务就已经写了好多代码，这就是原生库的缺点，这也就开始引入一个重要的点——Context
希望使用 Context 来代替一次请求
*/

// 逻辑上的内聚
func SignUp3(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	ctx := &Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败: %v\n", err)
		return
	}
}

/**
上面方法缺点是需要用户自己去创建Context，现在框架帮助用户创建Context
优化如下：
*/
// 框架在哪里帮我们创建这个Context？在server中创建
func SignUp(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadReqeustJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败: %v\n", err)
		return
	}
}

func main() {
	server := NewHttpServer("test-server")
	//server.Route("/user/signup", SignUp)
	server.Start(":8080")
}

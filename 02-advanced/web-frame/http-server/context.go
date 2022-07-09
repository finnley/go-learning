package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 逻辑上的内聚
// 为什么不声明成interface?
// 使用 Context 表示一个 http 请求
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// 从context中读取json
// 把body里面的数据直接帮我们序列化好
func (c *Context) ReadJson(req interface{}) error {
	// 帮我读取 body
	// 帮我序列化
	r := c.R
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(respJson)
	return err
}

// 还可以针对上面方法进一步封装
func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadReqeustJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

// 暴露操作Context的方法
// 因为我不希望你知道我在Route里面知道是怎么创建Context的
func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: writer,
		R: request,
	}
}

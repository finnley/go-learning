package demo_interface

import (
	"fmt"
	"testing"
)

type Handler interface {
	Handle()
}

type HandlerBaseOnTree struct {
}

// 比如私有的必然返回接口
type privateHandler struct {
}

func NewPrivateHandler() Handler {
	return nil
}

// 个人偏好，返回接口
// 原因：面向接口编程，不希望别人拿到我具体的类型的
// 就是别人用我创建的东西，不应该知道我的具体类型的，它只要拿到我的接口就行了
func NewHandlerBaseOnTreeV1() Handler {
	return nil
}

// Go推荐的，返回具体类型
func NewHandlerBaseOnTreeV2() *HandlerBaseOnTree {
	return nil
}

// demo2:
type MyService struct {
	// 使用接口可以注入 mock 的实现
	handler Handler
}

type MyHandler struct {
}

func (m MyHandler) Handle() {
	// 比如说这里调用一个下游
	// 可以是 http 可以是 rpc
}

func (ms *MyService) Serve() {
	ms.handler.Handle()
}

// mock测试
func TestMyService_Serve(t *testing.T) {
	ms := &MyService{
		handler: &mockHandler{},
	}
	fmt.Println(ms)
}

type mockHandler struct {
}

func (m mockHandler) Handle() {
	fmt.Println("hello")

}

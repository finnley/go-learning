package main

import "net/http"

// interface 的定义
// 接口：接口是一组行为的抽象
// 尽量用接口：以实现面向接口编程
//type Server interface {
//	//Route(pattern string, handleFunc http.HandlerFunc)
//	// 修改签名 由框架控制Context的创建
//	Route(pattern string, handleFunc func(ctx *Context))
//	Start(address string) error
//}

/**
上面接口中路由并不支持HTTP Method
*/
// 进一步优化
type Server interface {
	//Route(pattern string, handleFunc http.HandlerFunc)
	// 修改签名 由框架控制Context的创建
	Route(method string, pattern string, handleFunc func(ctx *Context))
	Start(address string) error
}

// sdkHttpServer 基于 http 库实现
type sdkHttpServer struct {
	Name    string
	handler *HandlerBasedOnMap
}

// Route 注册路由
//func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
//	http.HandleFunc(pattern, handleFunc)
//}
// 同样的需要修改接口实现
//func (s *sdkHttpServer) Route(pattern string, handleFunc func(ctx *Context)) {
//	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
//		ctx := NewContext(writer, request)
//		handleFunc(ctx)
//	})
//}

//func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
//	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
//		//if request.Method != method {
//		//	writer.Write([]byte("error "))
//		//}
//		ctx := NewContext(writer, request)
//		handleFunc(ctx)
//	})
//}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		// 只注册一遍，所以需要移动位置，启动的时候注册，并且handler 让sdkHttpServer持有
		//handler := &HandlerBasedOnMap{}
		//http.Handle("/", handler)
		key := s.handler.key(method, pattern)
		s.handler.handlers[key] = handleFunc
		// 这里还需要解决重复注册的问题
	})
}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

// 对定义接口通常对外暴露一个创建的方法，
// 这样就不会暴露内部实现，外部都不知道有个sdkHttpServer，因为sdkHttpServer实现了Server接口
// 当要返回一个实际类型所实现的接口时需要用指针
func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

// 暴露的另外的实现
func NewServerBasedOnXXX() Server {
	panic("implement")
}

// 或者
type Factory func() Server

var factory Factory

func RegisterFactory(f Factory) {
	factory = f
}
func NewServer() Server {
	return factory()
}

// 一般是在我使用第三方库，但又没有办法修改源码的情况下，又想仔扩展这个库的结构体方法，就会用这个
type Header map[string][]string

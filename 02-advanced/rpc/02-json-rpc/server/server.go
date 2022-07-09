package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}

func main() {
	// 1. 实例化一个 server
	listener, _ := net.Listen("tcp", ":1234")
	// 2. 注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3. 启动服务
	for {
		conn, _ := listener.Accept()
		//rpc.ServeConn(conn)
		// 替换 rpc 序列化协议为 json
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

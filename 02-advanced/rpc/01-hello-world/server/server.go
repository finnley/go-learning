package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

// 实际返回值并不是在 error 里面而是 reply，error 只是表示函数是否有错
func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改 reply 的值
	*reply = "hello, " + request
	return nil
}

func main() {
	// 1. 实例化一个 server
	listener, _ := net.Listen("tcp", ":1234")
	// 2. 注册处理逻辑 handler
	// 将Hello注册到rpc
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3. 启动服务
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)

}

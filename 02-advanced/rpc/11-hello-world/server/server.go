package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

// Hello方法必须满足Go语言的RPC规则：
// 方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。
func (p *HelloService) Hello(request string, reply *string) error {
	// 返回值通过修改 reply 的值
	*reply = "Hello:" + request
	return nil
}

func main() {
	// 将HelloService类型的对象注册为一个RPC服务
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

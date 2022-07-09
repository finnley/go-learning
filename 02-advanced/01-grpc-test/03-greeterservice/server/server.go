package main

import (
	"context"
	"go-learning/02-advanced/01-grpc-test/03-greeterservice/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

// 业务逻辑：
// 固定格式，在入参的时候加了 context ，可以用来处理协程超时问题，返回的时候加了 error
//				 SayHello(context.Context, *HelloRequest) (*HelloReply, error)
func (s *Server) SayHello(c context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 1. 实例化 gRPC Server
	g := grpc.NewServer()
	// 2. 注册 gRPC 将业务逻辑注册到server
	proto.RegisterGreeterServer(g, &Server{})
	// 3. 启动
	lis, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc: " + err.Error())
	}
}

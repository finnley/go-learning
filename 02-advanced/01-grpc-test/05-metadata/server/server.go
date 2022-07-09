package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/05-metadata/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	// 接收metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata err")
	}
	/**
	password [grpc]
	:authority [127.0.0.1:8080]
	content-type [application/grpc]
	user-agent [grpc-go/1.40.0]
	name [metadata]
	*/
	for key, val := range md {
		fmt.Println(key, val)
	}
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}

	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 1. 实例化 gRPC Server
	g := grpc.NewServer()
	// 2. 注册 gRPC
	proto.RegisterGreeterServer(g, &Server{})
	// 3. 启动 gRPC
	// 先监听 IP 和端口
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	// 再启动
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start server: " + err.Error())
	}
}

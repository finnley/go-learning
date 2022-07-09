package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/06-interceptor/proto"
	"net"

	"google.golang.org/grpc"
)

type Service struct {
}

func (s *Service) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 生成 GRPC 的 pretor option
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求")
		// handler是原本的调用逻辑
		//return handler(ctx, req)

		res, err := handler(ctx, req)
		fmt.Println("请求完成")
		return res, err
	}
	// 生成一个 UnaryInterceptor 的 operation 这个是一元grpc的拦截器，stream也有对应的拦截器
	opt := grpc.UnaryInterceptor(interceptor)

	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Service{})

	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = g.Serve(listen)
	if err != nil {
		panic("failed to start server: " + err.Error())
	}
}

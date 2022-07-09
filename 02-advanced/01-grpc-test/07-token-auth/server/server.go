package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/07-token-auth/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
}

func (s *Server) SayHello(context.Context, *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello " + "grpc",
	}, nil
}

func main() {
	// 服务端拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到一个请求")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			// gRPC 错误处理
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
			//fmt.Println("get metadata error")
		}

		var (
			appid  string
			appkey string
		)
		if va1, ok := md["appid"]; ok {
			appid = va1[0]
		}
		if va1, ok := md["appkey"]; ok {
			appkey = va1[0]
		}
		if appid != "101010" || appkey != "i am key" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}
		//if nameSlice, ok := md["appid"]; ok {
		//	fmt.Println(nameSlice)
		//	for i, e := range nameSlice {
		//		fmt.Println((i, e))
		//	}
		//}
		res, err := handler(ctx, req)
		fmt.Println("请求完成")
		return res, err
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})

	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = g.Serve(listen)
	if err != nil {
		panic("failed to start server: " + err.Error())
	}
}

package main

import (
	"context"
	"go-learning/02-advanced/01-grpc-test/09-error/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	//return &proto.HelloReply{
	//	Message: "hello " + request.Name,
	//}, nil
	// 异常处理如下：
	//return nil, status.Error(codes.InvalidArgument, "invalid username")
	return nil, status.Errorf(codes.NotFound, "未找到记录: %s", request.Name)
	//return nil, status.New(codes.InvalidArgument, "invalid username").Err()
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = g.Serve(listen)
	if err != nil {
		panic("failed to start server: " + err.Error())
	}
}

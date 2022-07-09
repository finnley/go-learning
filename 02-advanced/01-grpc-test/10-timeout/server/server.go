package main

import (
	"context"
	"go-learning/02-advanced/01-grpc-test/09-error/proto"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	time.Sleep(time.Second * 5)
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
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

package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/05-metadata/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 拨号
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 实例化 gRPC Client
	c := proto.NewGreeterClient(conn)

	// metadata
	//md := metadata.Pairs("timestamp", time.Now().Format())
	md := metadata.New(map[string]string{
		"name":     "metadata",
		"password": "grpc",
	})
	// 通过md生成新的ctx
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 调用远程服务
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "gRPC"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

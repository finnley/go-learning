package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/03-greeterservice/proto"
	"google.golang.org/grpc"
)

func main() {
	// 拨号连接
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 实例化 client
	c := proto.NewGreeterClient(conn)
	// 调用
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "gRPC",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

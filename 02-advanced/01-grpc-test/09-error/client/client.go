package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/05-metadata/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	_, err = c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "error",
	})
	if err != nil {
		//panic(err)
		// 客户端 Error 处理
		st, ok := status.FromError(err)
		if !ok {
			// Error was not a status error
			panic("解析 error 错误")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
}

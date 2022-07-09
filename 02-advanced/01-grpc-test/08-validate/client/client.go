package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/08-validate/proto"
	"google.golang.org/grpc"
)

type customCredential struct{}

func main() {
	var opts []grpc.DialOption

	//opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	//rsp, _ := c.Search(context.Background(), &empty.Empty{})
	rsp, err := c.SayHello(context.Background(), &proto.Person{
		Id:     1000,
		Email:  "43@43.com",
		Mobile: "13390703506",
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(rsp.Id)
	fmt.Println(rsp.Id)
}

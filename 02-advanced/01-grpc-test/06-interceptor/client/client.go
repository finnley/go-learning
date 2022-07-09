package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/05-metadata/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// 客户端拦截器
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时: %s\n", time.Since(start))
		return err
	}

	//opt := grpc.WithUnaryInterceptor(interceptor)
	//conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), opt)

	// grpc.WithInsecure(), opt可以进行合并，比如下面这种写法也可以
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "interpretor",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

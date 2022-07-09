package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/05-metadata/proto"
	"time"

	"google.golang.org/grpc/codes"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	"google.golang.org/grpc"
	//"github.com/grpc-ecosystem/go-grpc-middleware"
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

	retryOps := []grpc_retry.CallOption{
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(1 * time.Second),
		grpc_retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable),
	}

	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	// 调用时这个请求多长时间超时，需要重试多少次，当服务器返回什么状态码的时候重试
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOps...)))
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	//r, err := c.SayHello(context.Background(), &proto.HelloRequest{
	//	Name: "interpretor",
	//},
	//	grpc_retry.WithMax(3),
	//	grpc_retry.WithPerRetryTimeout(1*time.Second),
	//	grpc_retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable))
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "interpretor",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

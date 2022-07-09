package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/jaeger-test/otgrpc"
	"go-learning/02-advanced/jaeger-test/proto"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"google.golang.org/grpc"
)

func main() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.1.8:6831",
		},
		ServiceName: "go_shop",
	}

	// 生成tracer链路
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer) // 将tracer设为全局
	defer closer.Close()

	// 拨号连接
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
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

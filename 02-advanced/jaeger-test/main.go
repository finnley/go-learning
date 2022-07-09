package main

import (
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
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
	defer closer.Close()

	//// 1.span
	//span := tracer.StartSpan("go-grpc-web")
	//time.Sleep(time.Second)
	//defer span.Finish()

	// 2.发送多级嵌套span
	//span1 := tracer.StartSpan("funcA")
	//time.Sleep(time.Millisecond * 500)
	//defer span1.Finish() // 这边不要使用defer ，否则会出现执行时间是1.5s，如果只有1个步骤可以使用使用defer
	//span2 := tracer.StartSpan("funcB")
	//time.Sleep(time.Millisecond * 1000)
	//defer span2.Finish()
	// 上面修改为下面，不使用defer，但是下面却分开显示
	// 2.发送多级嵌套span
	//span1 := tracer.StartSpan("funcAA")
	//time.Sleep(time.Millisecond * 500)
	////defer span1.Finish()
	//span1.Finish()
	//span2 := tracer.StartSpan("funcBB")
	//time.Sleep(time.Millisecond * 1000)
	////defer span2.Finish()
	//span2.Finish()
	// 嵌套显示，显示在一起
	parentSpan := tracer.StartSpan("main")
	span1 := tracer.StartSpan("funcAAA", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(time.Millisecond * 500)
	span1.Finish()
	span2 := tracer.StartSpan("funcBBB", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(time.Millisecond * 1000)
	span2.Finish()
	parentSpan.Finish()
}

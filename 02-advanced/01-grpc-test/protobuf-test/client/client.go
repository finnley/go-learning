package main

import (
	"context"
	"fmt"
	"go-learning/02-advanced/01-grpc-test/protobuf-test/proto"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc"
	//timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	// 使用嵌套Message
	//proto.HelloReply_Result{
	//	Name: "",
	//	Url:  "",
	//}
	//proto.Pong{}

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Url: "notes.einscat.com",
		//Name: "notes",
		G: proto.Gender_MALE, // 枚举类型传值
		Map: map[string]string{
			"name":    "finnley",
			"company": "xuepincat",
		},
		AddTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(r.Message)
}

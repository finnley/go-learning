package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"

	"go-learning/02-advanced/01-grpc-test/04-steam/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//1.服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "server stream"})
	for {
		//接收
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}

	//2.客户端流模式
	//putS, _ := c.PutStream(context.Background())
	//i := 0
	//for {
	//	i++
	//	_ = putS.Send(&proto.StreamReqData{
	//		Data: fmt.Sprintf("client stream %d", i),
	//	})
	//	time.Sleep(time.Second)
	//	if i > 10 {
	//		break
	//	}
	//}

	//3.双向流模式
	//allStr, _ := c.AllStream(context.Background())
	//wg := sync.WaitGroup{}
	////两个协程  一个负责发送 一个负责接收
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	for {
	//		data, _ := allStr.Recv()
	//		fmt.Println("收到服务端消息： " + data.Data)
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	for {
	//		_ = allStr.Send(&proto.StreamReqData{Data: "客户端发送"})
	//		time.Sleep(time.Second)
	//	}
	//}()
	//wg.Wait()
}

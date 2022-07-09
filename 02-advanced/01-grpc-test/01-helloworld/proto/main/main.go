package main

import (
	"encoding/json"
	"fmt"
	helloworld "go-learning/02-advanced/01-grpc-test/01-helloworld/proto"

	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := helloworld.HelloRequest{
		Name:    "china",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	// 进行网络传输时的值 也就是序列化后的值
	rsp, _ := proto.Marshal(&req)
	fmt.Println("\n======== protobuf ========")
	fmt.Println(rsp)         // [10 5 99 104 105 110 97 16 18 26 2 103 111 26 3 103 105 110 26 9 229 190 174 230 156 141 229 138 161]
	fmt.Println(string(rsp)) // chinagogin      微服务
	fmt.Println(len(rsp))    // 29

	fmt.Println("\n======== json ========")
	jsonStruct := Hello{
		Name:    "china",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	jsonRsp, _ := json.Marshal(jsonStruct)
	fmt.Println(jsonRsp)         // [123 34 110 97 109 101 34 58 34 99 104 105 110 97 34 44 34 97 103 101 34 58 49 56 44 34 99 111 117 114 115 101 115 34 58 91 34 103 111 34 44 34 103 105 110 34 44 34 229 190 174 230 156 141 229 138 161 34 93 125]
	fmt.Println(string(jsonRsp)) // {"name":"china","age":18,"courses":["go","gin","微服务"]}
	fmt.Println(len(jsonRsp))    // 60

	// 反序列化
	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses) // china 18 [go gin 微服务]
}

/**
通过对比，protobuf 明显压缩比更高
*/

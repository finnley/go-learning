package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1.建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}
	//var reply *string // 初始化了一个指针变量，是nil,然后把nil传到下面的reply中就会出错，所以需要改成下面的方式：
	// new 的作用: 在内存中分配一块空间，并把指针赋给这个变量，这块空间里面有没有值无所谓，但至少已经有了地址，nil就不行，因为nil连地址都没有
	var reply *string = new(string)
	err = client.Call("HelloService.Hello", "golang rpc", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)

	//var reply string
	//err = client.Call("HelloService.Hello", "golang rpc", &reply)
	//if err != nil {
	//	panic("调用失败")
	//}
	//fmt.Println(reply)
}

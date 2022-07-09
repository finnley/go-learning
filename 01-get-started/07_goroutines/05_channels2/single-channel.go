package main

import "fmt"

func main() {
	channel := make(chan int, 10)
	channel <- 1
	fmt.Println(<-channel) // 1

	// 只读
	//var readonly <-chan int
	//// 只写
	//var writeonly chan<- int
	//
	//readonly = channel
	//writeonly = channel

	// 下面函数只读
}

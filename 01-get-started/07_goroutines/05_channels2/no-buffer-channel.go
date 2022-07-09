package main

import (
	"fmt"
	"time"
)

func main() {
	var channel chan int
	fmt.Printf("%T %#v\n", channel, channel) // chan int (chan int)(nil)

	channel = make(chan int)
	fmt.Printf("%T %#v\n", channel, channel) // chan int (chan int)(0xc00008c060)

	//channel <- 1 // fatal error: all goroutines are asleep - deadlock!

	//<-channel // fatal error: all goroutines are asleep - deadlock!

	//go func() {
	//	time.Sleep(time.Second * 5)
	//	// 写
	//	channel <- 1
	//}()
	//
	//fmt.Println("before")
	//
	//// 读
	//num := <-channel
	//fmt.Println(num) // 1

	go func() {
		fmt.Println("before<-")
		// 写
		channel <- 1
		fmt.Println("after <-")
	}()

	fmt.Println("before")

	time.Sleep(time.Second * 5)
	// 读
	num := <-channel
	fmt.Println(num) // 1
	time.Sleep(time.Second * 5)
}

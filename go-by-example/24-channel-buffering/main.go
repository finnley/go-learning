package main

import "fmt"

func main() {
	messages := make(chan string, 2) // make 了一个字符串通道，最多允许缓存 2 个值

	// 由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
	messages <- "buffered"
	messages <- "channel"
	//messages <- "over" // panic

	// 然后可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/**
buffered
channel

默认情况下，通道是 无缓冲 的，这意味着只有对应的接收（<- chan） 通道准备好接收时，才允许进行发送（chan <-）。
有缓冲通道 允许在没有对应接收者的情况下，缓存一定数量的值。

由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
*/

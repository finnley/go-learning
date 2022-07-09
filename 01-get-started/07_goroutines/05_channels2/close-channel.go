package main

import (
	"fmt"
	"time"
)

// 管道也可以关闭，当管道关闭以后，在读取的时候会直接返回一个数据，不会阻塞
func main() {
	channel := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		// 关闭管道，不再写入，读取的时候不会阻塞，读取到的是管道类型的零值
		//close(channel)
		channel <- struct{}{}
	}()

	fmt.Println("before")
	val, ok := <-channel
	fmt.Println(val, ok) // {} 当管道关闭，读取管道返回数据，不再进行阻塞
	fmt.Println("over")
	//time.Sleep(time.Second * 5)
}

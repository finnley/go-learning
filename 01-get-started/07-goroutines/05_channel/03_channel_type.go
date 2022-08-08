package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func main() {
	////01_channel.go 中的 channel 都是双向管道
	//var msg chan int
	//msg = make(chan int)
	//wg2.Add(1)
	////go consumer9(msg)
	//go func(queue chan int) {
	//	defer wg2.Done()
	//	for {
	//		data, ok := <-queue
	//		if !ok {
	//			break
	//		}
	//		fmt.Println(data, ok)
	//		time.Sleep(time.Second * 2)
	//		queue <- 2
	//	}
	//}(msg)
	//msg <- 1
	//fmt.Println("等待返回值")
	//fmt.Println(<-msg) // 2
	//close(msg)
	//wg2.Wait()

	//// 为了安全，还提供了单向 channel
	//var msg chan<- int // 只能放值，不能取值
	////var msg <-chan int // 只能取值，不能放值
	//msg = make(chan int, 10)
	//msg <- 1
	//msg <- 2
	////data := <-msg // 编译都不通过

	var msg chan int // 只能放值，不能取值
	msg = make(chan int, 10)
	wg2.Add(1)
	//go consumer10(msg) // 普通channel可以直接转为单向channel，反过来不行
	go func(queue <-chan int) {
		defer wg2.Done()
		for {
			data, ok := <-queue
			if !ok {
				break
			}
			fmt.Println(data, ok)
			time.Sleep(time.Second * 2)
		}
	}(msg)
	msg <- 1
	close(msg)
	wg2.Wait()
}

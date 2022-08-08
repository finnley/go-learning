package main

import (
	"fmt"
	"time"
)

// 有缓冲和无缓冲区别
func main1() {
	ch := make(chan int)
	ch <- 1
	value := <-ch
	fmt.Println(value)
}

/**
上面代码会报错：
如果定义了一个channel，没有指定缓冲区，先往里面写数据，再从里面读数据
fatal error: all goroutines are asleep - deadlock!

可修改如下：
*/
func main2() {
	ch := make(chan int)

	go func() {
		value := <-ch
		fmt.Println(value)
	}()
	ch <- 1

	time.Sleep(time.Second * 2)
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- 2
	}()

	value := <-ch
	fmt.Println(value)

	time.Sleep(time.Second * 2)
}

/**
总结：
有缓冲和无缓冲区别：
无缓冲要求发送数据和接收数据必须是同时准备好的，才能完成发送和接收，属于同步操作，会发生阻塞等待；
有缓冲的要求接收一个值，不要求同一时间来完成接收，属于异步操作
*/

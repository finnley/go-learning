package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 5)
		ch2 <- 2
	}()

	// 当两个goroutine都没有的时候会被阻塞住
	// select case 只要从任意一个读取成功就结束，如果两个都没有就阻塞
	// default 如果多个case都没有读取到数据，就执行default，否则一直读取不到会造成阻塞
	select {
	case e := <-ch1:
		fmt.Println("ch1: ", e)
	case e := <-ch2:
		fmt.Println("ch2: ", e)
	default:
		fmt.Println("default")
	}
}

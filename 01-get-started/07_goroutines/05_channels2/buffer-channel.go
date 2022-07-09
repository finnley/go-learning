package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var channel chan int
	// 带缓冲区的管道，容量是 3
	channel = make(chan int, 3)

	fmt.Printf("%T %#v %d %d\n", channel, channel, len(channel), cap(channel)) // chan int (chan int)(0xc000018080) 0 3

	// 带缓冲的当放入的数据超过缓冲长度之后，如果没有人读，再继续写入数据就会出现死锁
	//channel <- 1
	//channel <- 2
	//channel <- 3
	////channel <- 4 // fatal error: all goroutines are asleep - deadlock!

	//// 写入时慢，读取时快
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 10; i++ {
	//		//time.Sleep(time.Second * 2)
	//		fmt.Println("<-", i)
	//		channel <- i
	//	}
	//}()
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 10; i++ {
	//		time.Sleep(time.Second * 2)
	//		fmt.Println(<-channel)
	//	}
	//}()
	//wg.Wait()

	// 写入时慢，读取时快

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("<-", i)
			channel <- i
		}
		close(channel)
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 使用for range遍历的时候往管道写数据时要在写入后close，否则会发生死锁
		for e := range channel {
			time.Sleep(time.Second * 2)
			fmt.Println(e)
		}
	}()
	wg.Wait()
}

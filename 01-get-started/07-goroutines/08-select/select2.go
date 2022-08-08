package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 指定CPU核数
	//fmt.Println(runtime.GOMAXPROCS(0))  // 16
	fmt.Println(runtime.GOMAXPROCS(1)) // 16
	//fmt.Println(runtime.GOMAXPROCS(-1)) // 1

	intch := make(chan int, 1)
	strch := make(chan string, 1)
	intch <- 1
	strch <- "hello"

	select {
	case value := <-intch:
		fmt.Println(value)
	case value := <-strch:
		fmt.Println(value)
	case <-time.After(time.Second * 5): // 超时处理
		fmt.Println("超时")
	default:
		// 阻塞操作
	}
}

/**
select总结：
1.如果未准备好，会阻塞等待
2.如果多个case都满足条件，会随机选择一个执行
3.case里面必须是I/O操作
4.可以处理超时、退出的判断
5.如果select是以异步处理，需要在 for 循环中使用
6.select监控channel的数据流向
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// 用信道实现主程等待协程2s,如果超过2s,主程直接结束(不用sleep)?
	start := time.Now()
	wait := make(chan int, 1)
	go func() {
		fmt.Println("做点东西")
		time.Sleep(1 * time.Second)
		wait <- 2
	}()
	fmt.Println("这里是主程序")
	select {
	case nums := <-wait:
		fmt.Println(nums)
	case <-time.After(2 * time.Second):
		fmt.Println("2秒后")
	}
	fmt.Println(time.Since(start))
}

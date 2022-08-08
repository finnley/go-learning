package main

import (
	"fmt"
	"time"
)

func main() {
	// 空结构体不占用内存空间 起到通知作用
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("协程执行完成...")
		//ch <- struct{}{}
		close(ch)

	}()

	fmt.Println("主函数执行")
	<-ch
	fmt.Println("主函数执行完成")
}

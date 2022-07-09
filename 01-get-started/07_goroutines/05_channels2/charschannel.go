package main

import (
	"fmt"
)

// 使用channel代替waitGroup
func main() {
	channel := make(chan int)
	fmt.Println("start")
	for i := 0; i < 3; i++ {
		go func(prefix int) {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d:%c ", prefix, ch)
			}
			fmt.Printf("\n")
			channel <- prefix // 协程结束写入管道 总共写入了3次数据，如果管道读不到数据，写入会阻塞
		}(i)
	}

	// 如果能从管道读取到3次数据，就可以证明三个协程都结束了
	//<-channel
	//<-channel
	//<-channel
	// 可以上面这么写，也可以下面这么写：
	fmt.Println("before")
	for i := 0; i < 3; i++ {
		fmt.Println("over: ", <-channel)
	}
	fmt.Println("over")
}

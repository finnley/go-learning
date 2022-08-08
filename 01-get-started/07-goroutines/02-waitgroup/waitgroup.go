package main

import (
	"fmt"
	"sync"
	"time"
)

// bad
func main_bad() {
	// 下面做法不好的原因，不知道子协程什么时候结束，原来子协程可能要运行10s的，但是主协程睡了2s就结束退出了，子协程也就跟着挂了
	for i := 0; i <= 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Second * 2)
}

// 如何解决主的 goroutine 在子协程结束后自动结束
// 也就是主协程可以退出，但是要等子协程任务做完
var wg sync.WaitGroup

/**
sync.WaitGroup ，相当于定义了一个变量，是一个计数信号量
Add(): 启动之前有一个goroutine就加一个，协程的计数器，相当于添加信号
defer Done: 执行完之后要反馈执行完了，计数器减一个，相当于处理完信号
Wait: 主协程如果知道计数器数量为0，就可以结束退出了，主协程需要阻塞，相当于等待计数器归零，当所有信息都处理完成后结束
*/

// good
func main() {
	// Add 的数量必须和Done的数量相等
	wg.Add(5) // 如果4？结果如何？
	for i := 0; i <= 5; i++ {
		// 启动一个协程 add + 1  ,这里知道有5个，也可以在循环外面直接写5
		//wg.Add(1)
		//go func(n int) {
		//	fmt.Println(n)
		//	// 每运行 一个 -1
		//	defer wg.Done()
		//}(i)
		go f(i)
	}
	// 主协程只要知道里面没有了，为0了就表示子协程结束了，所以主协程需要阻塞 wait
	wg.Wait()
}

func f(n int) {
	defer wg.Done() // 如果没有写done ,就会死锁
	fmt.Println(n)
}

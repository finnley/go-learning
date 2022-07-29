// 速率限制
package main

import (
	"fmt"
	"time"
)

func main() {
	// 首先，我们将看一个基本的速率限制。 假设我们想限制对收到请求的处理，我们可以通过一个渠道处理这些请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。
	limiter := time.Tick(200 * time.Microsecond)

	// 通过在每次请求前阻塞 limiter 通道的一个接收， 可以将频率限制为，每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。 我们可以通过缓冲通道来完成此任务。 burstyLimiter 通道允许最多 3 个爆发（bursts）事件。
	burstyLimiter := make(chan time.Time, 3)

	// 填充通道，表示允许的爆发（bursts）。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Microsecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/**
request 1 2021-09-09 10:57:13.5199755 +0800 CST m=+0.003904401
request 2 2021-09-09 10:57:13.5336388 +0800 CST m=+0.017567701
request 3 2021-09-09 10:57:13.5346145 +0800 CST m=+0.018543401
request 4 2021-09-09 10:57:13.5355915 +0800 CST m=+0.019520401
request 5 2021-09-09 10:57:13.5365666 +0800 CST m=+0.020495501
request 1 2021-09-09 10:57:13.5365666 +0800 CST m=+0.020495501
request 2 2021-09-09 10:57:13.5365666 +0800 CST m=+0.020495501
request 3 2021-09-09 10:57:13.5365666 +0800 CST m=+0.020495501
request 4 2021-09-09 10:57:13.5375428 +0800 CST m=+0.021471701
request 5 2021-09-09 10:57:13.5385187 +0800 CST m=+0.022447601

速率限制 是控制服务资源利用和质量的重要机制。 基于协程、通道和打点器，Go 优雅的支持速率限制。
*/

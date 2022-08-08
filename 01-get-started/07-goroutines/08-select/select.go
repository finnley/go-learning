package main

func main() {
	///**
	//作用 channel 之上，用于解决多个channel中选择问题，可以理解为多路复用，和socket中的多路复用有区别
	//select 会随机公平的选择一个 case 语句执行
	//select 应用场景：1. timeout 超时机制
	//*/
	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)
	//ch1 <- 1
	//ch2 <- 2
	//
	//// 想从多个 channel 里面任意（随机）选择一个 channel，就是用 select
	//select {
	//case data := <-ch1:
	//	fmt.Println(data)
	//case data := <-ch2:
	//	fmt.Println(data)
	//}

	//fmt.Println("select 场景的使用:")
	///**
	//select 应用场景：1. timeout 超时机制
	//*/
	//timeout := false
	//go func() {
	//	// 该 goroutine 如果执行时间超多 1s 就报告给主的 goroutine
	//	// 主协程如何知道子协程超时了呢？可以设置一个全局变量 timeout,但是这种方式并不好，因为需要一个全局变量和一个for循环不停的去检测
	//	time.Sleep(time.Second * 2)
	//	timeout = true
	//}()
	//for {
	//	if timeout {
	//		fmt.Println("结束")
	//		break
	//	}
	//	time.Sleep(time.Millisecond * 10) // 防止 for 过快
	//}

	////可以使用select 的方式来优化
	//timeout := make(chan bool, 2)
	//go func() {
	//	time.Sleep(time.Second * 5)
	//	fmt.Println(111)
	//	timeout <- true
	//}()
	//timeout2 := make(chan bool, 2)
	//go func() {
	//	time.Sleep(time.Second * 1)
	//	fmt.Println(222)
	//	timeout2 <- true
	//}()
	////fmt.Println(<-timeout) // 这个地方取值的时候是会阻塞的
	//select {
	//case <-timeout:
	//	// 没有取出来就表示阻塞了
	//	fmt.Println("超时了")
	//case <-timeout2:
	//	fmt.Println("超时了2")
	//	default: // 使用default防止select被阻塞住
	//		fmt.Println("继续下一次") // 如果执行这句就表示前面的所有channel都阻塞了
	//}

	//fmt.Println("场景2: ")
	/**
	2. 判断某个 channel 是否阻塞
	比如当有多个channel需要监控的时候，就可以使用select来从里面取数据
	或者多个goroutine其中出现了问题比如timeout就可以使用select来达到想要的效果
	*/

}

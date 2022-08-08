package main

import (
	"fmt"
	"sync"
	"time"
)

/**
如果一个goroutine想向另一个goroutine发送消息，怎么做？使用消息队列，channel 提供了一种通信机制，定向，python java 中使用的是消息队列

golang 并发编程，也会涉及到两个 goroutine 之间的通信，这就是go语言中的channel
*/
var wg sync.WaitGroup

func main() {
	/*
		// 1.定义一个 channel 值为 nil，引用类型需要初始化，如果引用类型不初始化，值为nil
		var msg chan int // 该代码只进行了定义，没有进行初始化
		// 2.初始化
		msg = make(chan int) // 第一种初始化方式：无缓冲
		// 3.放数据
		msg <- 1 // 将 1 放到 channel 中
		// 4.取数据
		data := <-msg
		fmt.Println(data)
	*/
	//// 到这里上面的代码运行会失败 fatal error: all goroutines are asleep - deadlock!
	//// 上面错误原因是因为 channel 初始化的方式有两种，一种是上面的无缓冲的，可理解为消息空间是没有的，一种是下面有缓冲的
	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//data := <-msg
	//fmt.Println(data) // 1

	////如果改成下面的，先往 channel 放了一个1，然后又放了个2，此时会报错吗
	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//msg <- 2
	//data := <-msg
	//fmt.Println(data)
	////上面代码是会出错的 fatal error: all goroutines are asleep - deadlock!
	////所以如果缓冲空间满了，再往里面放数据就会死锁
	//// 上面的channel看起来好像就是一个有空间的数组，这个空间是定死的，已经满了再往里面放数据就会抛异常，那这个和固定长度数组有什么意义呢？
	//// 通常初始化channel后如果只管往里面放值是会出问题的，一般情况下必须保持有往里面放值的地方，就必须要有从里面取值的地方，也就是说需要用另外的goroutine来消费这个channel
	////channel 一般往里面放数据就要有从里面取值的地方，也就是需要另一个 goroutine 来消费这个channel，所以修改成下面方式:

	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//wg.Add(1)
	////go consumer(msg) // 1
	//go func(queue chan int) {
	//	defer wg.Done()
	//	data := <-queue
	//	fmt.Println(data)
	//}(msg)
	////这里防止主goroutine退出，也需要进行阻塞
	//wg.Wait()

	////上面的情况如果先往 channel 放了一个1，然后又放了个2，此时会报错吗？
	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1 //这一步是可以运行成功的，但是这一步后并没有去启动goroutine消费，所以消费者并没有来消费
	//msg <- 2 //因为上一步消费者并没有来消费，再往里面放值
	//wg.Add(1)
	////go consumer(msg) // 1
	//go func(queue chan int) {
	//	defer wg.Done()
	//	data := <-queue
	//	fmt.Println(data)
	//}(msg)
	//wg.Wait()
	////上面的还是会报错 fatal error: all goroutines are asleep - deadlock!
	////在放进去是成功的，但是并没有启动goroutine,所以消费者还没有来得及消费，所以再往里面放值的时候是还要出错的，所以可以在启动之后再放 ，如下：

	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//wg.Add(1)
	////go consumer(msg) // 1
	//go func(queue chan int) {
	//	defer wg.Done()
	//	data := <-queue
	//	fmt.Println(data)
	//}(msg)
	//msg <- 2 // 在启动goroutine之后再来往里面放值
	//wg.Wait()

	////如果现在只放了一个1，消费的时候却不断从里面取，结果如何？代码如下：
	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//wg.Add(1)
	////go consumer2(msg)
	//go func(queue chan int) {
	//	defer wg.Done()
	//	data := <-queue
	//	data2 := <-queue
	//	data3 := <-queue
	//	fmt.Println(data, data2, data3)
	//}(msg)
	//wg.Wait()
	////上面这段也会死锁 fatal error: all goroutines are asleep - deadlock!
	////在满的情况下发一个数据就要尽量把里面消费完
	////取数据还有下面这种 for range 取法：

	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//wg.Add(1)
	////go consumer3(msg) // 1 2
	//go func(queue chan int) {
	//	defer wg.Done()
	//	for data := range queue {
	//		fmt.Println(data)
	//	}
	//}(msg)
	//msg <- 2
	//wg.Wait()
	////但是还是会死锁 fatal error: all goroutines are asleep - deadlock!
	////上面运行后取出了1 和 2，是因为当2放进去，for 循环立马取出来
	////取完第二次再取第三次的时候就超过了限制，channel里面已经没有了，再强行从空的里面取数据，显然有问题
	////上面修改如下：

	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//wg.Add(1)
	////go consumer4(msg)
	//go func(queue chan int) {
	//	defer wg.Done()
	//	for data := range queue {
	//		fmt.Println(data)
	//	}
	//}(msg)
	//msg <- 2
	//// 第二个发送完关闭 channel
	//// 1.已经关闭的 channel 不能再发送数据
	//close(msg)
	////msg <- 3 //已经关闭的 channel 不能再发送数据，所以这样错误 panic: send on closed channel
	//wg.Wait()

	//// 2.已经关闭的 channel 还可以继续消费，直到数据取完
	//// 所以结论是已经关闭的channel能够继续取数据，知道数据取完为止，但不能继续发数据
	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//wg.Add(1)
	////go consumer5(msg)
	//go func(queue chan int) {
	//	defer wg.Done()
	//	for data := range queue {
	//		fmt.Println(data)
	//		time.Sleep(time.Second * 2) // 先消费1，然后发送方发送2后关闭，此时Sleep的时候2肯定进来了，然后是channel已经关闭了
	//	}
	//}(msg)
	//msg <- 2
	//close(msg) // 第二个发送完关闭 channel
	//wg.Wait()

	////使用 for 循环取数据，使用这种方式取数据，一旦取完取出来后还会继续取，只是值默认是 0
	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//wg.Add(1)
	////go consumer6(msg)
	//go func(queue chan int) {
	//	defer wg.Done()
	//	for {
	//		data := <-queue
	//		fmt.Println(data)
	//		time.Sleep(time.Second * 2)
	//	}
	//}(msg)
	//msg <- 2
	//close(msg)
	//wg.Wait()
	//// 针对上面取完数据再继续取值为0问题，修改如下：

	//// 希望取完后判断是否关闭，如果关闭就break不再取数据
	//// 取数据可接收两个值
	//var msg chan int
	//msg = make(chan int, 1)
	//msg <- 1
	//wg.Add(1)
	////go consumer7(msg)
	//go func(queue chan int) {
	//	defer wg.Done()
	//	for {
	//		data, ok := <-queue // ok 表示管道有没有关闭
	//		fmt.Println(data, ok)
	//		if !ok {
	//			break
	//		}
	//		time.Sleep(time.Second * 2)
	//	}
	//}(msg)
	//msg <- 2
	//close(msg)
	//wg.Wait()

	//// 无缓冲空间的使用，无缓冲空间现在一往里面放数据就报错，那应该怎么放呢？
	//var msg chan int
	//msg = make(chan int)
	//// == 无缓冲空间在放数据时对面一定要是有准备接收的，也就是在放之前先启动消费创建一个 goroutine，
	//// 但是会出现取出值之后再取值就有问题，第一次可以等待，第二次就不可以等待了 所以这里放1了之后进行关闭
	//wg.Add(1)
	////go consumer8(msg)
	//go func(queue chan int) { // 先启动一个消费者
	//	defer wg.Done()
	//	for {
	//		data, ok := <-queue
	//		fmt.Println(data, ok)
	//		if !ok {
	//			break
	//		}
	//		time.Sleep(time.Second * 2)
	//	}
	//}(msg)
	//// 当进行放数据到 msg 中的时候，这个时候会阻塞的，阻塞之前会获取一把锁，现在还没有人来竞争锁，因为没有goroutine,获取之后就要等，等别人消费，
	//// 那这把锁什么时候释放，肯定是要等到数据被消费之后
	//// channel 是多个goroutine之间线程安全的，也就是多个 goroutine 对同一个channel 进行操作是不会出现并发的数据一致性问题,这背后如何保证的？肯定是使用了锁，why?
	//msg <- 1 // 如果是一个没有缓冲的 channel,在没有启动一个消费者之前，放数据就会报错，除非在放数据之前启动了一个 goroutine 从这个channel中获取数据
	//close(msg)
	//wg.Wait()

	////上面的 channel 都是双向管道
	//var msg chan int
	//msg = make(chan int)
	//wg.Add(1)
	//go consumer9(msg)
	//msg <- 1
	//fmt.Println("等待返回值")
	//fmt.Println(<-msg) // 2
	//close(msg)
	//wg.Wait()

	//// 为了安全，还提供了单向 channel
	//var msg chan<- int // 只能放值，不能取值
	////var msg <-chan int // 只能取值，不能放值
	//msg = make(chan int, 10)
	//msg <- 1
	//msg <- 2
	////data := <-msg // 编译都不通过

	//var msg chan int // 只能放值，不能取值
	//msg = make(chan int, 10)
	//wg.Add(1)
	//go consumer10(msg) // 普通channel可以直接转为单向channel，反过来不行
	//msg <- 1
	//close(msg)
	//wg.Wait()
}

func consumer(queue chan int) {
	defer wg.Done()
	data := <-queue
	fmt.Println(data)
}

func consumer2(queue chan int) {
	defer wg.Done()
	data := <-queue
	data2 := <-queue
	data3 := <-queue
	fmt.Println(data, data2, data3)
}

func consumer3(queue chan int) {
	defer wg.Done()
	for data := range queue {
		fmt.Println(data)
	}
}

func consumer4(queue chan int) {
	defer wg.Done()
	for data := range queue {
		fmt.Println(data)
	}
}

func consumer5(queue chan int) {
	defer wg.Done()
	for data := range queue {
		fmt.Println(data)
		time.Sleep(time.Second * 2) // 先消费1，然后发送方发送2后关闭，此时Sleep的时候2肯定进来了，然后是channel已经关闭了
	}
}

func consumer6(queue chan int) {
	defer wg.Done()
	for {
		data := <-queue
		fmt.Println(data)
		time.Sleep(time.Second * 2)
	}
}
func consumer7(queue chan int) {
	defer wg.Done()
	for {
		data, ok := <-queue // ok 表示管道有没有关闭
		if !ok {
			break
		}
		fmt.Println(data, ok)
		time.Sleep(time.Second * 2)
	}
}

func consumer8(queue chan int) {
	defer wg.Done()
	for {
		data, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(data, ok)
		time.Sleep(time.Second * 2)
	}
}

func consumer9(queue chan int) {
	defer wg.Done()
	for {
		data, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(data, ok)
		time.Sleep(time.Second * 2)
		queue <- 2
	}
}

func consumer10(queue <-chan int) {
	defer wg.Done()
	for {
		data, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(data, ok)
		time.Sleep(time.Second * 2)
	}
}

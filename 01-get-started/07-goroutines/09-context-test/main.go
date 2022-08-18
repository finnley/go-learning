package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
需求：
主线程中需要启动两个goroutine分别去完成两个不同的任务。比如：
1.启动一个goroutine去监控某台服务器的CPU使用情况

*/

var wg sync.WaitGroup

// 场景一：比如现在启动一个goroutine去监控某台服务器的CPU使用情况
func cpuInfo() {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	fmt.Println("CPU信息读取完成")
}

// 场景二：现在需求升级了，CPU一般不是只监控一次，而是不停的监控，
// 就需要每隔一段时间报告下， 所以需要没过一段时间报告下

// 场景二的第一种处理方式
var stop bool

func cpuInfo2() {
	defer wg.Done()
	for {
		if stop {
			break
		}
		time.Sleep(time.Second * 2)
		fmt.Println("CPU信息读取完成")
	}
}

// 场景二的第二种处理方式
// 使用channel的方式：阻塞写法
var stopChan chan bool = make(chan bool)

func cpuInfo3() {
	defer wg.Done()
	for {
		select {
		case <-stopChan: // 单纯这样写是有问题的，因为这是阻塞的，会一直等待发进来，发进来之后才会进行后面的逻辑
			fmt.Println("退出CPU监控333")
			return
		}
		// 如果select阻塞，下面两行一次执行的机会都没有
		time.Sleep(time.Second * 2)
		fmt.Println("这里不会读取到")
	}
}

// 使用channel的方式：非阻塞写法
func cpuInfo4() {
	defer wg.Done()
	for {
		select {
		case <-stopChan: // 单纯这样写是有问题的，因为这是阻塞的，会一直等待发进来，发进来之后才会进行后面的逻辑
			fmt.Println("退出CPU监控")
			return
		default:
			// 为解决select出现阻塞导致的下面两行一次执行的机会都没有这一问题
			// 如果不希望阻塞，而是每次进来判断一下，所以需要加个default，不至于被阻塞住
			// 如果有case和default会优先满足case
			time.Sleep(time.Second * 2)
			fmt.Println("这里会读取到：CPU信息读取完成")
		}
	}
}

// 主协程中想要取消生成的goroutine里面逻辑的时候，场景的使用有channel，但是golang还提供了 context
// golang 提供了context 更加能优雅的解决上面的问题
// 可以在主的goroutine中取消子goroutine，使用方式和select差不多，但是有些区别
func cpuInfo5(ctx context.Context) {
	defer wg.Done()
	go memberInfo5(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("监控CPU退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("获取CPU成功")
		}
	}
}

// context的特性：
// 父context被取消，那么这个父context生成的子context也会被取消，可以达成链式取消
func cpuInfo6(ctx context.Context) {
	defer wg.Done()
	ctx2, _ := context.WithCancel(ctx)
	go memberInfo5(ctx2)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("监控退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("获取CPU成功")
		}
	}
}

func memberInfo5(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("监控内存退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("获取内存成功")
		}
	}
}

// timeout
func cpuInfo7(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("监控退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("获取CPU成功")
		}
	}
}

func main() {
	//// ======== 场景一 ========
	//// 现在启动一个goroutine去监控某台服务器CPU使用情况
	//wg.Add(1)
	//go cpuInfo()
	//wg.Wait()
	//fmt.Println("信息监控完成")

	//// ======== 场景二 ========
	//// 现在既希望可以读取CPU信息，也可以中断CPU的信息读取，这个怎么处理？
	//// 一种方式：监控全局变量
	//wg.Add(1)
	//go cpuInfo2()
	//// 下面两行不能写在wait之后，因为如果写在wait之后会在goroutine结束之后，就永远得不到执行，
	//// 这是因为goroutine里面是个死循环
	//time.Sleep(time.Second * 6)
	//stop = true
	//wg.Wait()
	//fmt.Println("信息监控完成")
	//// 弊端：这种全局变量的方式可以完成，但是并不优雅

	//// 使用channel的方式：阻塞写法
	//wg.Add(1)
	//go cpuInfo3()
	//time.Sleep(time.Second * 6)
	//stopChan <- true
	//wg.Wait()
	//fmt.Println("信息监控完成")

	//// 使用channel的方式
	//wg.Add(1)
	//go cpuInfo4()
	//time.Sleep(time.Second * 6)
	//stopChan <- true
	//wg.Wait()
	//fmt.Println("信息监控完成")

	//// 使用context的方式
	//wg.Add(1)
	//ctx, cancel := context.WithCancel(context.Background()) // context.Background()这个写法比较固定
	//go cpuInfo5(ctx)
	//time.Sleep(time.Second * 6)
	//cancel()
	//wg.Wait()
	//fmt.Println("信息监控完成")

	//// 使用context.cancel的方式
	//wg.Add(2)
	//ctx, cancel := context.WithCancel(context.Background()) // context.Background()这个写法比较固定
	//go cpuInfo5(ctx)
	////go memberInfo5(ctx)
	//// 或者注释掉上面一行在cpuInfo中嵌套调用
	//time.Sleep(time.Second * 6)
	//cancel()
	//wg.Wait()
	//fmt.Println("信息监控完成")

	// 使用channel的方式
	wg.Add(1)
	// 3秒后过期
	//ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	//go cpuInfo7(ctx)

	// 如果不等3s，等1s钟
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go cpuInfo7(ctx)
	time.Sleep(time.Second)
	// 1s后手动取消
	cancel() // cancel可以采用默认的方式，也可以随时手动取消，但是取消一定要在超时之前取消，之后是没有用的
	wg.Wait()
	fmt.Println("信息监控完成")
}

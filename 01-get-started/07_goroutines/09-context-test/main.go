package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 比如现在启动一个goroutine去监控某台服务器的CPU使用情况
func cpuInfo() {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	fmt.Println("CPU信息读取完成")
}

var stop bool

// 现在需求升级了，CPU不是只监控一次，而是不停的监控
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

// 使用channel的方式
var stopChan chan bool = make(chan bool)

func cpuInfo3() {
	defer wg.Done()
	for {
		select {
		case <-stopChan: // 单纯这样写是有问题的，因为这是阻塞的，会一直等待发进来，发进来之后才会进行后面的逻辑
			fmt.Println("退出CPU监控")
			return
		}
		// 如果select阻塞，下面两行一次执行的机会都没有
		//time.Sleep(time.Second * 2)
		//fmt.Println("CPU信息读取完成")
	}
}

func cpuInfo4() {
	defer wg.Done()
	for {
		select {
		case <-stopChan: // 单纯这样写是有问题的，因为这是阻塞的，会一直等待发进来，发进来之后才会进行后面的逻辑
			fmt.Println("退出CPU监控")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("CPU信息读取完成")
		}
		// 如果select阻塞，下面两行一次执行的机会都没有
		// 并不希望阻塞，而是每次进来判断一下，所以需要加个default，不至于被阻塞住
		// 如果有case和default会有限满足case
	}
}

// golang 提供了context 更加能优雅的解决上面的问题
// 可以在主的goroutine中取消子goroutine，使用上和select差不多，但是有些区别

func cpuInfo5(ctx context.Context) {
	defer wg.Done()
	go memberInfo5(ctx)
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
	// 现在启动一个goroutine去监控某台服务器CPU使用情况
	wg.Add(1)
	go cpuInfo()
	wg.Wait()
	fmt.Println("信息监控完成")

	//// 现在既希望可以读取CPU信息，也可以中断CPU的信息读取，这个怎么处理？
	//// 一种方式是监控全局变量
	//wg.Add(1)
	//go cpuInfo2()
	//time.Sleep(time.Second * 6)
	//stop = true
	//wg.Wait()
	//fmt.Println("信息监控完成")
	//// 这种全局变量的方式可以完成，但是并不优雅

	//// 使用channel的方式
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

	//// 使用channel的方式
	//wg.Add(1)
	//ctx, cancel := context.WithCancel(context.Background()) // context.Background()这个写法比较固定
	//go cpuInfo5(ctx)
	//time.Sleep(time.Second * 6)
	//cancel()
	//wg.Wait()
	//fmt.Println("信息监控完成")

	//// 使用channel的方式
	//wg.Add(2)
	//ctx, cancel := context.WithCancel(context.Background()) // context.Background()这个写法比较固定
	//go cpuInfo5(ctx)
	////go memberInfo5(ctx)
	//// 或者注释掉上面一行在cpuInfo中嵌套调用
	//time.Sleep(time.Second * 6)
	//cancel()
	//wg.Wait()
	//fmt.Println("信息监控完成")

	//// 使用channel的方式
	//wg.Add(1)
	////ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	////go cpuInfo7(ctx)
	//// 如果不等3s，等1s钟
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//go cpuInfo7(ctx)
	//time.Sleep(time.Second)
	//// 1s后手动取消
	//cancel() // cancel可以采用默认的方式，也可以随时手动取消，但是取消一定要在超时之前取消，之后是没有用的
	//wg.Wait()
	//fmt.Println("信息监控完成")
}

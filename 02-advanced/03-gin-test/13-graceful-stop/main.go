package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 优雅退出:
	// 当我们关闭程序的时候应该做得后续处理
	// 应用还在运行，运行过程有些数据还没有入库，就突然关闭了，程序被迫终端，就有一部分数据没有处理完，比如处理订单数据就非常危险，会出现数据不一致问题
	// 所以我们需要让程序在突然中断是做些处理，没处理完的数据先处理下再关闭，另外还可以将数据坐下日志记录，等以后还原
	// 类似这些处理都是优雅退出

	// 微服务 启动之前或者启动之后会做一些事情：将当前的服务的ip地址和端口号注册到注册中心，希望我这个服务可以让别人发现，
	// 如果没有做优雅退出，导致在结束这个进程的之后并没有告知注册中心

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	// 使用这个方式启动，主进程不会停止，会一直挂着，如果想要结束，就需要将启动的过程放到另外的协程中运行，这样主线程就可以等到它的信号比如Ctrl+C或者KILL信号等
	// 当接收到信号之后就将主线程退出
	//router.Run(":8083")

	go func() {
		router.Run(":8083")
	}()
	// 如果想要接收到信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 处理后续逻辑
	fmt.Println("关闭 server 中...")
	fmt.Println("注销服务")
}

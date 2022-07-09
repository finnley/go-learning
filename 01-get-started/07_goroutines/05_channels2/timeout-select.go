package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	result := make(chan interface{})
	timeout := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		timeout <- struct{}{}
	}()

	go func() {
		r := rand.Intn(10)
		fmt.Println(r)
		time.Sleep(time.Second * time.Duration(r))
		result <- r
	}()

	select {
	case <-timeout:
		// 如果能从timeout读取出数据，就表示超时了
		fmt.Println("超时...")
	case r := <-result:
		fmt.Println("success", r)
	}
}

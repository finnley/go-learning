package main

import (
	"fmt"
	"time"
)

func timeTicker(interval time.Duration) <-chan time.Time {
	timeChannel := make(chan time.Time)
	go func() {
		for {
			time.Sleep(interval)
			timeChannel <- time.Now()
		}
	}()

	return timeChannel
}

func main() {
	//fmt.Println(time.Now())
	//for now := range timeTicker(time.Second * 3) {
	//	fmt.Println(now)
	//}

	// 一种方式结束
	//endTime := time.Now().Add(time.Second * 20)
	//for now := range time.Tick(time.Second * 3) {
	//	fmt.Println(now)
	//	if now.After(endTime) {
	//		break
	//	}
	//}

	endTime := time.Now().Add(time.Second * 20)
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	for now := range ticker.C {
		fmt.Println(now)
		if now.After(endTime) {
			break
		}
	}
}

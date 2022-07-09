package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	now1 := time.Now().Unix()
	fmt.Println(now1)

	now2 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now2)

	latestOperationTime := 1655565190 // 2022-06-18 23:13:10
	earliestRequestTime := time.Unix(int64(latestOperationTime), 0).Add(5 * 60 * time.Second)
	fmt.Println(earliestRequestTime.Unix())
	fmt.Println(time.Now().Unix())

	if earliestRequestTime.Before(time.Now()) {
		fmt.Println("aaa")
	}
}

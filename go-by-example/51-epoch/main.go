package main

import (
	"fmt"
	"time"
)

func main() {

	// 分别使用 time.Now 的 Unix 和 UnixNano， 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now) // 2021-09-10 09:34:08.6299427 +0800 CST m=+0.002928401

	millis := nanos / 1000000
	fmt.Println(secs) // 1631237648
	fmt.Println(millis) // 1631237648629
	fmt.Println(nanos) // 1631237648629942700

	fmt.Println(time.Unix(secs, 0)) // 2021-09-10 09:34:08 +0800 CST
	fmt.Println(time.Unix(0, nanos)) // 2021-09-10 09:34:08.6299427 +0800 CST
}

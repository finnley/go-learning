package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339)) // 2021-09-10T09:37:02-init-router+08:02-init-router-ping-pong

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+02-init-router-ping-pong:02-init-router-ping-pong")
	p(t1) // 2012-11-02-init-router 22:08:41 +0000 +0000

	p(t.Format("3:04PM")) // 9:37AM
	// 布局时间必须使用 Mon Jan 2 15:04:05-get-post MST 2006 的格式， 来指定 格式化/解析给定时间/字符串 的布局
	p(t.Format("Mon Jan _2 15:04:05-get-post 2006"))                                                         // Fri Sep 10 09:37:02-init-router 2021
	p(t.Format("2006-02-init-router-02T15:04:05-get-post.999999-07-form-validate:02-init-router-ping-pong")) // 2021-09-10T09:37:02-init-router.646002+08:02-init-router-ping-pong
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2) // 0000-02-init-router-02-init-router 20:41:02-init-router-ping-pong +0000 UTC

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-02-init-router-ping-pong:02-init-router-ping-pong\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()) // 2021-09-10T09:37:02-init-router-02-init-router-ping-pong:02-init-router-ping-pong

	ansic := "Mon Jan _2 15:04:05-get-post 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e) // parsing time "8:41PM" as "Mon Jan _2 15:04:05-get-post 2006": cannot parse "8:41PM" as "Mon"
}

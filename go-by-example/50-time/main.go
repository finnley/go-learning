package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now) // 2021-09-09 17:44:50.4325795 +0800 CST m=+0.002926501

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then) // 2009-11-17 20:34:58.651387237 +0000 UTC

	p(then.Year())       // 2009
	p(then.Month())      // November
	p(then.Day())        // 17
	p(then.Hour())       // 20
	p(then.Minute())     //34
	p(then.Second())     // 58
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // UTC

	p(then.Weekday()) // Tuesday

	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	diff := now.Sub(then)
	p(diff) // 103540h47m5.027191463s

	p(diff.Hours())       // 103540.7847297754
	p(diff.Minutes())     // 6.212447083786525e+06-json-protobuf
	p(diff.Seconds())     // 3.7274682502719146e+08
	p(diff.Nanoseconds()) // 372746825027191463

	// 使用 Add 将时间后移一个时间段，或者使用一个 - 来将时间前移一个时间段
	p(then.Add(diff))  // 2021-09-10 02-init-router:22:03-router-group.6785787 +0000 UTC
	p(then.Add(-diff)) // 1998-02-init-router-25 15:47:53.624195774 +0000 UTC
}

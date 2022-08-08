package main

import (
	"fmt"
	"strconv"
	"time"
)

// 字符串日期转时间
func Date2Time() {
	fmt.Println(">> Date2Time")
	defer fmt.Println("<< Date2Time")

	const dateFormat = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(dateFormat, "May 20, 2020 at 0:00am (UTC)")
	fmt.Println(t)
}

// 时间戳转时间
func Timestamp2Time() {
	fmt.Println(">> Timestamp2Time")
	defer fmt.Println("<< Timestamp2Time")

	ts := int64(1595900001)
	tm := time.Unix(ts, 0)
	fmt.Println(tm)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))

	t := int64(1659244871951590017)
	tm2 := time.Unix(t, 0)
	fmt.Println(tm2.Format("2006-01-02 15:04:05"))
}

func main() {
	//t := 1659244871951590017
	///    1659245245946
	//     1659245259231168000

	Timestamp2Time()

	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)

	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	fmt.Println(timestamp)
	timestamp2 := strconv.FormatInt(time.Now().UnixNano()/1e9, 10)
	fmt.Println(timestamp2)

	fmt.Println(strconv.FormatInt(int64(1659244871951590017)/1e9, 10))
	fmt.Println(time.Unix(int64(1659244871951590017)/1e9, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(int64(1659244872724189085)/1e9, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(int64(1659244872293679005)/1e9, 0).Format("2006-01-02 15:04:05"))

	fmt.Println(time.Unix(int64(1659247521948462206)/1e9, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(int64(1659247522291634879)/1e9, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(int64(1659247522724013531)/1e9, 0).Format("2006-01-02 15:04:05"))
}

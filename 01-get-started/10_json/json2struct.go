package main

import (
	"encoding/json"
	"fmt"
)

// 结构体刷新小写字母开头不会转换，会转成对应的零值
type Student2 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score []int  `json:"score"`
	Addr  string `json:"addr"`
}

func main() {
	slice := []byte(`
{
    "Addr": "江苏省",
    "Age": 20,
    "Name": "嘿嘿嘿",
    "Score": [
        100,
        29,
        59,
        80
    ]
}
`)
	var stu Student2
	err := json.Unmarshal(slice, &stu)
	if err != nil {
		fmt.Println("json转换失败")
	} else {
		fmt.Println(stu)
	}
}

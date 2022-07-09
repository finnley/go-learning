package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score []int
	Addr  string
}

func main() {
	stu := Student{
		Name:  "哈哈",
		Age:   20,
		Score: []int{100, 28, 84, 59},
		Addr:  "江苏省",
	}
	// 生成JSON数据格式
	//slice, err := json.Marshal(stu)
	// 生成格式化json信息
	slice, err := json.MarshalIndent(stu, "*", "----")
	if err != nil {
		fmt.Println("json转换失败")
	} else {
		fmt.Println(string(slice))
	}
}

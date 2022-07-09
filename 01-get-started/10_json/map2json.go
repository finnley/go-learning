package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 创建字典保存键值对信息
	m := make(map[string]interface{})
	m["Name"] = "嘿嘿嘿"
	m["Age"] = 20
	m["Score"] = []int{100, 29, 59, 80}
	m["Addr"] = "江苏省"

	//slice, err := json.Marshal(m)
	//格式化json
	slice, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Println("json格式转换失败")
	} else {
		fmt.Println(string(slice))
	}
}

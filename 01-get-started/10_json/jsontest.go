package main

import (
	"encoding/json"
	"fmt"
)

type StudentTest struct {
	Name  string
	Age   int
	Score []int
	Addr  string
}

func main() {
	stu := StudentTest{
		Name:  "张三",
		Age:   30,
		Score: []int{100, 23, 98},
		Addr:  "江苏省",
	}
	json2, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json转换错误")
	} else {
		fmt.Println(string(json2))
	}

	m := make(map[string]interface{})
	m["name"] = "李四"
	m["age"] = 28
	m["score"] = []int{199, 343}
	m["addr"] = "江苏省"
	json3, _ := json.Marshal(m)
	fmt.Println(string(json3))

	json4 := []byte(`
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
	var student4 StudentTest
	err4 := json.Unmarshal(json4, &student4)
	if err4 != nil {
		fmt.Println("转换失败")
	} else {
		fmt.Println(student4)
	}
}

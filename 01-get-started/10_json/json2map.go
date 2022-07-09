package main

import (
	"encoding/json"
	"fmt"
)

/**
json转map步骤：
1.现将json转成接口类型的数据
2.对接口类型的数据进行类型断言判断里面每个数据的类容，本身转成的是map，键是string，值是interface
需要注意上面接口类型的数据格式是map格式，但是是个接口类型
*/
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

	// 定义一个接口类型变量
	var i interface{}
	// json转map 进行Unmarshal操作之后类型还是接口类型，但是格式是个map格式
	err := json.Unmarshal(slice, &i)
	if err != nil {
		fmt.Println("json转换失败")
	} else {
		fmt.Println(i)         // map[Addr:江苏省 Age:20 Name:嘿嘿嘿 Score:[100 29 59 80]]
		fmt.Printf("%T\n", &i) // *interface {}
	}
	//switch i.(type) {
	//
	//}
	m := i.(map[string]interface{})
	fmt.Println(m)        // map[Addr:江苏省 Age:20 Name:嘿嘿嘿 Score:[100 29 59 80]]
	fmt.Printf("%T\n", m) // map[string]interface {}
	for _, v := range m {
		//fmt.Println(k, v)
		switch val := v.(type) {
		case string:
			fmt.Println("string: ", val)
		case int:
			fmt.Println("int:", val)
		case float64:
			fmt.Println("float64:", val)
		case []int:
			fmt.Println("[]int:", val)
		case []interface{}:
			fmt.Println("interface")
			for i, s := range val {
				fmt.Println(i, s)
			}
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 空接口：可以存储任意类型数据，空接口类型数据无法计算
	var i interface{} = 123 // 等号后面可接任意数据类型
	var a interface{} = 123 // 等号后面可接任意数据类型
	t1 := reflect.TypeOf(i)
	t2 := reflect.TypeOf(a)
	fmt.Println(t1, t2) // int int
	if t1 == t2 {
		v1 := reflect.ValueOf(i).Int()
		v2 := reflect.ValueOf(a).Int()
		fmt.Println(v1, v2) // 123 123
		fmt.Println(v1 + v2)
	}
}

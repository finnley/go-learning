package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	typeof := reflect.TypeOf(People{})
	// 通过反射将结构体标签作为循环的值
	// 遍历结构体成员
	for i := 0; i < typeof.NumField(); i++ {
		field := typeof.Field(i)
		tag := field.Tag.Get("json")
		fmt.Println(tag)
	}
}

/**
name
age
*/

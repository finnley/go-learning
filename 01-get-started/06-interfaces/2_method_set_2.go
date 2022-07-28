package main

import "fmt"

type people struct {
	Name string
}

// 类型断言
// 1.和map的条件判断类似
func main444() {
	var i interface{} = &people{Name: "哈哈"}
	// 类型断言 1.和map的条件判断类似
	if v, ok := i.(*people); ok {
		fmt.Println(v.Name) // 哈哈
		fmt.Println(v)      // &{哈哈}
	} else {
		fmt.Println("类型匹配失败")
	}
}

// 2.使用switch进行类型断言操作选择 i.(type)
func main() {
	var i interface{} = people{"嘿嘿 "} // people
	//var i interface{} = &people{"嘿嘿 "} // *people
	// 2.使用switch进行类型断言操作选择 i.(type)
	switch i.(type) {
	case people:
		fmt.Println("people")
	case *people:
		fmt.Println("*people")
	default:
		fmt.Println("类型匹配失败")
	}
}

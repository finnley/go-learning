package main

import "fmt"

// 闭包：将匿名函数作为函数的返回值
// seq() 后面的 func() int 这个整体是seq函数的返回值类型
// 匿名函数可以长时间的驻留在内存中
func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 调用的时候f对应的就是seq里面的匿名函数，这个过程就是闭包，调用f就相当于执行了seq里面的匿名函数流程
	f := seq()
	value := f()
	fmt.Println(value) // 1
	value = f()
	fmt.Println(value) // 2
	value = f()
	fmt.Println(value) // 3

	f1 := seq()
	value1 := f1()
	fmt.Println(value1) // 1
	value1 = f1()
	fmt.Println(value1) // 2
}

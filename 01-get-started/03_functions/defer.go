package main

import "fmt"

// defer 将函数加载到内存中并没有执行，而是在函数出栈时执行
func main() {
	//fmt.Println("A")
	//fmt.Println("B")
	//fmt.Println("C")
	//fmt.Println("D")

	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	defer fmt.Println("D")

	//i := 10
	//defer fmt.Println(i)
	//i += 10
	//fmt.Println(i)
}

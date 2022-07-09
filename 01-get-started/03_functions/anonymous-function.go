package main

import "fmt"

// 定义函数类型
func test() {

}

// 第二种：定义一个函数类型
// 使用type 为一个类型起一个别名，func(int, int)为函数定义的格式
type FUNCTION func(int, int)

// 匿名函数: 在函数内部定义函数，函数没有函数名
func main() {
	////第一种：定义匿名函数并调用
	//func(a int, b int) {
	//	fmt.Println(a + b)
	//}(10, 20)

	//// 打印函数信息，函数在代码区的内存地址
	//fmt.Println(test) // 0x109f050

	//定义函数没有调用会出现错误，可以定义一个函数类型
	//定义函数类型变量f，FUNCTION是一个函数类型，类似var s string s = "golang"
	//var f FUNCTION // var f func(int, int)
	//f = func(a int, b int) {
	//	fmt.Println(a + b)
	//}
	//fmt.Println(f) // 0x109f0e0
	//// 可以通过f去调用一个匿名函数函数
	//f(10, 20) // 30

	// 定义与调用合并
	f := func(a int, b int) {
		fmt.Println(a + b)
	}
	fmt.Printf("%T\n", f) // func(int, int)
	fmt.Println(f)        // 0x109f0e0
	// 可以通过f去调用一个匿名函数函数
	f(10, 20) // 30
}

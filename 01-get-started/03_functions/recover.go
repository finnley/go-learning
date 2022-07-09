package main

import (
	"errors"
	"fmt"
)

func demo(i int) {
	// recover错误拦截一定要出现在panic错误之前，要未雨绸缪
	// 如果代码中出现了错误，或者错误拦截，后续代码是不会继续执行的，比如demo最后一行的打印，但是超出该函数内的可以继续执行，比如main中的打印
	// 有点像try catch
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("err: ", err)
	//	}
	//}()
	var arr [10]int
	arr[i] = 123
	fmt.Println(arr)
}

// 一般性错误处理
func errorTest(a int, b int) (v int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	v = a / b
	return
}

func main() {
	// 如果没有错误，recover是可以执行到的，只不过是没有拦截到错误信息
	demo(10) // panic: runtime error: index out of range [10] with length 10
	fmt.Println("程序继续执行...")
}

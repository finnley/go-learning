package main

import "fmt"

func adder() func(int) int {
	sum := 0
	// 闭包
	// 函数体首先会有局部变量，func 是函数体，v是参数也是局部变量，sum不是函数体里面定义的，sum是外面定义的，是自由变量
	// 函数返回的时候返回的是一个闭包
	// 返回的是这个函数以及对sum的引用，以及sum变量会保存下来，保存到函数里面去
	return func(v int) int {
		sum += v
		return sum
	}
}

// 正统函数式编程
// 返回累加之后的值，还有一个是下一轮的函数
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

// 函数式编程
func main() {
	//a := adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("0+1+...+%d = %d\n", i, a(i))
	//}

	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0+1+...+%d = %d\n", i, s)
	}
}

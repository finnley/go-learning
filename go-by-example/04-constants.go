package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

/**
常量可以用于计算吗？
Line15: 常量表达式可以执行任意精度的运算

数值类型常量没有确定的类型，直到被给定某个类型，比如显示类型转化

数值类型常量可以根据上下文需要自动确定类型，比如Line:20
*/

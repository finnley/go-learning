package main

import "fmt"

// 3. 定义结构体
type Course struct {
	name  string
	price int
}

// 4. 定义接口
type Callable interface {
}

// 5. 定义函数别名 并没有函数体
type handle func(str string)

func main() {
	var aa byte
	fmt.Printf("%T\n", aa) // uint8
	// 1. 给一个类型定义别名 别名一般用于兼容性处理，所以不需要过多关注
	// var a byte; // byte is an alias for uint8
	// 实际byte本质上仍然是uint8，编译的时候仍然会将byte替换为uint8，type 也就是可以给另外内置的类型取一个别名
	type myByte = byte    // byte is an alias for uint8
	var b myByte          // var b uint8
	fmt.Printf("%T\n", b) // uint8
	// 那为什么会有byte,而不直接使用uint8呢？就是为了强调处理的对象是字节类型，也是为了代码的可读性，uint8到底是处理数字还是字节

	// 2. 基于一个已有的类型定义一个新的类型（没有等号）
	type myInt int // 可以看做是将int类型的所有代码都弄过来，int有的myInt也有
	var i myInt
	fmt.Printf("%T\n", i) // main.myInt

	// 3. 定义结构体
	// 4. 定义接口
	// 5. 定义函数别名
}

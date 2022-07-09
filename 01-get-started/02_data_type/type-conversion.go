package main

import "fmt"

func main0201() {
	// 菜市场 黄瓜3.25/斤 5斤
	price := 3.25 // float64
	weight := 5   // int

	fmt.Printf("%T\n", price)  // float64
	fmt.Printf("%T\n", weight) // int

	// 1.类型转换：要转的数据类型(变量名)
	//sum := price * weight // invalid operation: price * weight (mismatched types float64 and int)
	//total := price * float64(weight)
	//fmt.Println(total)

	// 浮点型转整形，会忽略小数部分，保留整数部分，不会四舍五入，可能会丢失精度，大类型转成小类型也可能造成精度丢失 int(3.99) = 3
	//total := int(price) * weight
	//fmt.Println(total)

	// 2.类型转换：要转的数据类型(表达式)
	total := int(price * float64(weight)) // 16
	fmt.Println(total)
}

//type int1 int
//type int2 int
type int1 = int
type int2 = int

func main() {
	//var a int8 = 123
	//var b int = 234
	//// 整形有不同的格式，但是不同的整形也是不允许混合计算的，需要类型转换
	////fmt.Println(a + b) // error: Invalid operation: a + b (mismatched types int8 and int)
	//fmt.Println(int(a) + b) // ok

	//// ch 是rune类型，rune 等价于int32，在底层是 type rune = int32，所以是可以相互计算的
	//// type 如果带等号（=）可以看成和那个类型是一样的，如果不带等号（=）就是两个不同的类型
	//ch := 'A'
	//fmt.Printf("%T\n", ch) // int32
	//var b int32 = 100
	//fmt.Println(ch + b) // 165

	// type 后面没有带等号（=）go会认为int1和int2是两个不同的类型
	var a int1 = 123
	var b int2 = 123
	fmt.Println(a + b) // error: Invalid operation: a + b (mismatched types int1 and int2)
	fmt.Println(a + int1(b))
}

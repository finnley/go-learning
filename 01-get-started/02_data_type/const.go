package main

import "fmt"

func main() {
	// 1. 常量的数据类型可以是布尔，数字，字符串
	// 2. 不曾使用的常量，在编译的时候不会报错

	// 常量定义
	// 1. 省略类型
	// 2. 定义多个常量（类型相同）
	// 3. 定义多个常量（类型不同）
	// 4. 定义多个常量，省略类型
	const aaa = 1
	const (
		bbb             = 1
		ccc             = 1
		ddd, eee        = "ddd", "eee"
		fff, ggg        = 100, "gggg"
		hhh      int    = 10
		iii      string = "iii"
		//j string // 不能只声明不赋值
	)

	// 常量组如不指定类型和初始化值，该类型和值和上一行的类型一致
	const (
		x int = 16
		y
		s = "abc"
		z
	)
	fmt.Println(x, y, s, z)                 // 16 16 abc abc
	fmt.Printf("%T %T %T %T\n", x, y, s, z) // int int string string

	// go没有枚举类型，可以通过常量和枚举关键字，枚举关键字：iota
	// iota是常量，常用作常量计数器，枚举，只和行有关
	// 几十个计数器，是第几个常量，从0开始
	const (
		a    = iota       //0
		b                 //1
		c                 //2
		d    = "ha"       //独立值，07-iota += 1 = 3
		e                 //"ha"   07-iota += 1 = 4
		f    = 100        //100    07-iota +=1 = 5
		g                 //100    07-iota +=1 = 6
		h, i = iota, iota //7, 7  07-iota + 1 = 7
		j    = iota       //8      恢复计数 ==> 8
		//k                 //9
		//l                 //10
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j) // 0 1 2 ha ha 100 100 7 7 8 9 10

	const (
		aa = iota // 0
		bb = iota // 1
		cc = iota // 2
	)
	fmt.Println(aa, bb, cc) // 0 1 2

	// 下面的这个要注意下，B1值是200
	const (
		A1 = (iota + 1) * 100 // 100
		B1                    // 200
		C1                    // 300
		D1 = 10               //10
		E1 = "hello"          // hello
		F1                    // hello
		G1 = iota             // 6
	)
	fmt.Println(A1, B1, C1, D1, E1, F1, G1) // 100 200 300 10 hello hello 6
}

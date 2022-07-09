package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 字符串基本操作
	// 1. 字符串长度
	var name string = "china:大中国"
	// 返回的是字节的长度 英文或者字符使用 ASCII 编码，一旦涉及到中文就不行
	// 中文使用的 unicode 字符集，存储的时候需要编码
	// utf8 编码，可以用 3个 字节表示一个中文，也可以用 1个字节表示英文
	fmt.Println(len(name)) // 15 = 6(英文) + 9(中文)
	// 类型转换 可以将 name 转成 rune(int32) 数组
	// 将每一个字符转成 rune
	name_arr := []rune(name)
	name_arr2 := []int32(name)                // 和上一行是等价的
	fmt.Println(len(name_arr))                // 9
	fmt.Println(len(name_arr2))               // 9
	fmt.Println(name_arr)                     // [99 104 105 110 97 58 22823 20013 22269]
	fmt.Println(name_arr2)                    // [99 104 105 110 97 58 22823 20013 22269]
	fmt.Println(utf8.RuneCountInString(name)) // 9 对应9个unicode字符

	fmt.Println([]rune(name)) // [99 104 105 110 97 58 22823 20013 22269]
	fmt.Println(name[7])      // 164
	fmt.Println(name_arr[7])  // 20013

	s := "hello, world"
	fmt.Println(len(s))     // 12
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	fmt.Println([]rune(s))  // [104 101 108 108 111 44 32 119 111 114 108 100]

	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i]) // 63 68 69 6e 61 3a e5 a4 a7 e4 b8 ad e5 9b bd
	}

	//for i := 0; i < len(name); i++ {
	//	fmt.Printf("%c ", name[i]) // 63 68 69 6e 61 3a e5 a4 a7 e4 b8 ad e5 9b bd
	//}
	println()
	var b = "hello world"
	for i := 0; i < len(b); i++ {
		fmt.Printf("%x ", b[i]) // 68 65 6c 6c 6f 20 77 6f 72 6c 64
	}

	println()
	var c = "Go语言"
	for i := 0; i < len(c); i++ {
		fmt.Printf("%x ", c[i]) // 47 6f e8 af ad e8 a8 80
	}

	fmt.Println()
	str := "Hello, 世界"
	// len() 统计的是自己，str有13个字节
	/**
	0       H       1
	1       e       1
	2       l       1
	3       l       1
	4       o       1
	5       ,       1
	6               1
	7       世      3
	10      界      3
	*/
	for i := 0; i < len(str); {
		// 显式地调用utf8.DecodeRuneInString解码
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}

	/**
	0       'H'     72
	1       'e'     101
	2       'l'     108
	3       'l'     108
	4       'o'     111
	5       ','     44
	6       ' '     32
	7       '世'    19990
	*/
	// * Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

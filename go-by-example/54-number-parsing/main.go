package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234

	// 在使用 ParseInt 解析整型数时， 例子中的参数 0 表示自动推断字符串所表示的数字的进制。 64 表示返回的整型数是以 64 位存储的
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i) // 123

	// ParseInt 会自动识别出字符串是十六进制数
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d) // 456

	// ParseUint 也是可用的
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u) // 789

	// Atoi 是一个基础的 10 进制整型数转换函数
	k, _ := strconv.Atoi("135")
	fmt.Println(k) // 135

	_, e := strconv.Atoi("wat")
	fmt.Println(e) // strconv.Atoi: parsing "wat": invalid syntax
}

package main

import "fmt"

func main() {
	// 转义符
	name := "china:\"大中国\""
	fmt.Println(name) // china:"大中国"

	date := `2021\09\09`
	fmt.Println(date) // 2021\09\09
}

package main

import "fmt"

func main() {
	name, desc := "中国", "上海"
	func(name string) {
		fmt.Println(name, desc) // 中国 上海
		//name, desc = "江苏", "东台"
		name, desc := "江苏", "东台"
		fmt.Println(name, desc) // 江苏 东台
	}(name)
	fmt.Println(name, desc) // 中国 东台
}

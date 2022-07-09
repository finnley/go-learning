package main

import "fmt"

func main() {
	var x, y int = 3, 6
	fmt.Println(&x, &y)
	
	argf2(&x, &y)
	fmt.Println(x, y)
}

func argf2(a, b *int) {
	fmt.Println(a, b)

	*a = *a + *b
	*b = 888
}

// 请问下这里传参时是将x,y的地址复制了一份放到了一个两个新变量里面传到argf1里面的吗
func argf1(a, b int) {
	fmt.Println(&a, &b)
	a = a + b
	b = 9999
}

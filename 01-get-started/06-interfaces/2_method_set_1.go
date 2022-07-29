package main

import "fmt"

type Animal interface {
	GetName() string
}

type People struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func (p *People) GetName() string {
	return p.Name
}

func main() {
	var n Animal
	//p1 := People{}
	p2 := &People{}
	//n = p1 // 无法编译
	// 如果是引用传递必须是取地址操作，否则就像上一行无法通过编译
	n = p2
	fmt.Println(n)
}

package main

import "fmt"

type Person3 struct {
	Name string
	Age  int
}

type Student3 struct {
	Person3
	Score int
}

func (p *Person3) Print() {
	fmt.Println(*p)
}

func (s *Student3) Print() {
	fmt.Println(*s)
}

func main() {
	// 实例化 student 对象
	stu := Student3{
		Person3: Person3{
			Name: "小周",
			Age:  19,
		},
		Score: 100,
	}
	// 采用就近原则，默认使用本结构体对应的方法
	stu.Print()         // {{小周 19} 100}
	stu.Person3.Print() // {小周 19}
}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student2 struct {
	Person
	Score int
}

func (p *Person) SayHello() {
	fmt.Println(*p)
}

func main() {
	// 实例化Student对象
	// 子类可以直接掉员工父类方法
	stu := Student2{
		Person: Person{
			Name: "小张",
			Age:  20,
		},
		Score: 100,
	}
	// 子类直接可以调用父类方法，但是并没有将 student 的所有信息都打印出来，因为score 在父类中是没有的
	stu.SayHello() // {小张 20}
}

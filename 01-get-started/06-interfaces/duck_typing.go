package main

import "fmt"

// 先定义一个接口，和使用此接口作为参数的函数：
type IGreeting interface {
	sayHello()
}

func greeting(i IGreeting) {
	i.sayHello()
}

// 再来定义两个结构体：
type Golang struct {
}

func (g Golang) sayHello() {
	fmt.Println("Hi, I am Go!")
}

type PHP struct {
}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP")
}

func main() {
	golang := Golang{}
	php := PHP{}

	greeting(golang)
	greeting(php)
}

/**
If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.

在 main 函数中，调用调用 greeting() 函数时，传入了 golang, php 对象，它们并没有显式地声明实现了 IGreeting 类型，只是实现了接口所规定的 sayHello() 函数。实际上，编译器在调用 sayHello() 函数时，会隐式地将 golang, php 对象转换成 IGreeting 类型，这也是静态语言的类型检查功能。
*/

package main

import "fmt"

// e.g. interface
type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是鱼🐟，假装自己是一只鸭🦆\n")
}

// 定义了一个新的类型，注意是新类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是山寨鱼🐟，嘎嘎嘎\n")
}

// 定义了一个新类型（type 用法之一：基于一个已有类型定义一个新类型）
type StrongFakeFish Fish

// 新类型自己的方法
func (f StrongFakeFish) Swim() {
	fmt.Printf("我是华强北山寨鱼🐟，嘎嘎嘎\n")
}

func main() {
	// e.g.
	fake := FakeFish{}
	// fake 无法调用原来 Fish 的方法
	//fake.Swim() // 这一句会编译错误
	fake.FakeSwim() // 我是山寨鱼🐟，嘎嘎嘎

	// 转换为Fish
	td := Fish(fake) // 类型转换
	// 类型转换类似于：
	//a := 10
	//var b int64 = int64(a)
	// 真的变成了🐟
	td.Swim() // 我是鱼🐟，假装自己是一只鸭?

	sFake := StrongFakeFish{}
	// 这里就是调用共了自己的方法
	sFake.Swim() // 我是华强北山寨鱼🐟，嘎嘎嘎

	td = Fish(sFake)
	// 真的变成了鱼
	td.Swim() // 我是鱼🐟，假装自己是一只鸭?

}

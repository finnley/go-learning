package main

import "fmt"

// e.g.
// ToyDuck 玩具鸭
type ToyDuck struct {
	Color string
	Price uint64
}

func (t ToyDuck) Swim() {
	fmt.Printf("门前一条河，游过一群鸭，我是%s, %f\n", t.Color, t.Price)
}

func main() {
	// e.g.
	// duck1 是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Swim()

	// duck2 是结构体的实例
	duck2 := ToyDuck{}
	duck2.Swim()

	// duck3 是 *ToyDuck
	// new 出来的也是指针
	// new 分配好内存，并把内存初始化好，内存单元的每个比特位都置为零值
	duck3 := new(ToyDuck)
	duck3.Swim()

	// 当你声明这样的时候，Go 就帮你分配好内存
	// 不用担心空指针的为你，因为它压根就不是指针
	var duck4 ToyDuck // 这种写法光声明，编译器会帮我们处理好为 var duck4 ToyDuck = ToyDuck{}
	duck4.Swim()

	// duck5 就是一个指针
	// 这是一个指针，但是不直到指针指向哪里
	//var duck5 *ToyDuck // 其实也是帮分配好了内存，这个内存只能放指针，并没有对指针赋值，所以这个指针是个 nil
	//duck5.Swim()       // 这边回直接panic掉

	// 赋值，初始化按字段名字赋值
	duck6 := ToyDuck{
		Color: "黄色",
		Price: 100,
	}
	duck6.Swim()

	// 初始化字段顺序赋值，不建议使用
	duck7 := ToyDuck{
		"蓝色",
		1024,
	}
	duck7.Swim()

	/// 后面再单独赋值
	duck8 := ToyDuck{}
	duck8.Color = "橘色"
}

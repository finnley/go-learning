package main

import (
	"fmt"
)

// 接口是一个协议-下面是自己创建的一个协议，程序员协议
// 1. 能够写代码，2. 能够解决bug 就是一个程序员，至于是男是女，多少岁不关心，其实就是一组方法的集合
type Programmer interface {
	Coding() string // 方法只是声明
	Debug() string
}

type Designer interface {
	Design() string
}

type Manager interface {
	Programmer
	Designer
	Manage() string
}

type UIDesigner struct {
}

func (d UIDesigner) Designer() string {
	fmt.Println("UI设计")
	return "UI设计"
}

// java 中一种类型只要继承了一个接口，如果继承了这个接口，那么这个接口里面的所有方法都必须要全部实现

type Pythoner struct {
	UIDesigner
	lib   []string
	kj    []string
	years int
}

func (p Pythoner) Coding() string {
	fmt.Println("Python开发者")
	return "Python开发者"
}

func (p Pythoner) Debug() string {
	fmt.Println("我会Python Debug")
	return "我会Python Debug"
}

func (p Pythoner) Manage() string {
	fmt.Println("我会管理")
	return "我会Python 管理"
}

func (p Pythoner) Design() string {
	fmt.Println("我会设计")
	return "我会Python 设计"
}

type Go struct {
}

func (p Go) Coding() string {
	fmt.Println("Go开发者")
	return "Go开发者"
}

func (p Go) Debug() string {
	fmt.Println("我会Go Debug")
	return "我会Go Debug"
}

// 对于 Pythoner 这个结构体来说，实现任何方法都可以，可以只实现 Coding ,也可以只实现 Debug,甚至两个都不实现，只要不全部实现 Coding,Debug 的话，那Pythoner就不是一个 Programmer 类型，为何要care是不是 Programmer 呢？
// 04-interface: 在 go 中接口是一种类型，抽象类型，既然是类型，就可以基于这个类型声明一个对象

// 多态
// 比如支付，有微信支付，支付宝支付，银行卡支付，系统可以支持各种类型的支付，每一种支付类型都有统一的接口
// 定一个协议，1. 创建订单 2. 支付，3. 查询支付状态 4. 退款
type AliPay struct {
}

type Wechat struct {
}

type Bank struct {
}

var b Bank
var a AliPay
var w Wechat

// 如果后期接入一种新的支付，或者取消已有支付，就必须要来修改源码
//var x = Tongyong
//x = Bank{}
//x = AliPay{}
//x = Wechat{}
//x.create()
//x.pay()
//x.query()

func HandlePro(p Programmer) {

}

type MyError struct {
}

func (m MyError) Error() string {
	return "错误"
}

// 1. Pythoner 本身自己就是一个类型，
func main() {
	// 接口帮完成了go语言的多态
	var pro Programmer = Pythoner{}
	pro.Coding()

	var pros []Programmer // 一堆程序员 不管是做什么的
	pros = append(pros, Pythoner{}, Go{})

	// 接口虽然是一种类型，但是接口是一种抽象类型，只有方法声明，没有实现
	p := Pythoner{}
	fmt.Printf("%T\n", p)   // main.Pythoner
	fmt.Printf("%T\n", pro) // main.Pythoner

	// 接口也支持继承
	// 1.go struct 组合 组合一起实现了所有的接口方法也是可以的
	// 2.接口本身也支持组合
	var m Manager = Pythoner{}
	fmt.Println(m)

	//var err error = error{}
	//var err error = MyError{}
	//var err error = errors.New("错误") // 可以这么使用，但是实际中这种使用并不多，使用比较多的是下面这种
	s := "文件不存在"
	var err error = fmt.Errorf("错误: %s\n", s)
	fmt.Println(err)
}

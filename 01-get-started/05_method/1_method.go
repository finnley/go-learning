package main

import "fmt"

// 面向对象中有属性和行为
// 可以通过结构体来描述对象的属性
// 可以通过方法来描述对象的行为

// 定义结构体，表示对象属性
type Student1 struct {
	Name string
	Age  int
}

// 值对象接受者
// 方法中不能修改对象属性
func (stu Student1) SayHello1() {
	//fmt.Println("大家好，我是", stu.Name)

	// 值对象接受者区别：方法中尝试修改对象属性
	stu.Name = "小赵"
	fmt.Println(stu) // {小赵 19}
}

// 指针对象接受者
// 方法中可以修改对象属性
func (stu *Student1) SayHello2() {
	//fmt.Println("大家好，我是", stu.Name)

	// 指针对象接受者区别：方法中尝试修改对象属性
	stu.Name = "小王"
}

func main() {
	// 调用时必须实例化对象，也就是对结构体进行初始化
	stu := Student1{
		Name: "小张",
		Age:  19,
	}
	//// 值对象接收者调用
	//stu.SayHello1() // 大家好，我是 小张
	//// 指针对象接收者调用
	//stu.SayHello2()    // 大家好，我是 小张
	//(&stu).SayHello2() // 大家好，我是 小张

	// 值对象接受者区别：方法中尝试修改对象属性
	//stu.SayHello1()  // {小赵 19} 仅修改了方法里面的值
	//fmt.Println(stu) // {小张 19} 但对象的属性并没有真正修改
	//指针对象接受者区别：方法中尝试修改对象属性
	stu.SayHello2()
	fmt.Println(stu) // {小王 19} 对象的属性被修改

	fmt.Println(stu.SayHello2) // 0x109f2b0
}

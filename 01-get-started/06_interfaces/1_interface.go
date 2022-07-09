package main

import "fmt"

// 接口定义
type Humaner interface {
	SayHi()
}

// 定义完接口一般需要实例化具体对象，用来实现接口
type Person struct {
	Name string
	Age  int
}

// 两个子类继承Person
type teacher struct {
	Person
	Subject string
}

// 两个子类继承Person
type Student struct {
	Person
	Score int
}

// 实现接口
func (t *teacher) SayHi() {
	//func (t teacher) SayHi() {
	fmt.Printf("大家好，我是%s， 我%d岁了， 我教%s\n", t.Name, t.Age, t.Subject)
}

// 实现接口
func (s *Student) SayHi() {
	//func (s Student) SayHi() {
	fmt.Printf("大家好，我是%s， 我%d岁了， 我的成绩%d\n", s.Name, s.Age, s.Score)
}

func (t *teacher) Print() {

}

func main222() {
	tea := teacher{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		Subject: "英语",
	}
	stu := Student{
		Person: Person{
			Name: "小强",
			Age:  20,
		},
		Score: 100,
	}

	// 创建接口类型变量
	var h1, h2 Humaner
	// 将对象地址赋值给接口类型的变量，就可以通过地址找到对象信息，就可以通过接口调用不同对象
	h1 = &tea
	h2 = &stu
	// 如果是值接收者可传地址，可传值
	//h1 = tea
	//h2 = stu
	// 通过接口调用
	h1.SayHi() // 大家好，我是张三， 我30岁了， 我教英语
	h2.SayHi() // 大家好，我是小强， 我20岁了， 我的成绩100
}

// 多态：将接口类型作为函数参数
func SayHi(h Humaner) {
	// 不管什么对象，只要满足接口都可以调用对应方法
	// 如果是学生就可以走学生的方法，如果是老师就走老师的方法
	// 也就是传入不同的人物形态就走不同的人物结果
	h.SayHi()
}

func main() {
	// 1.初始化对象
	tea := teacher{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		Subject: "英语",
	}
	stu := Student{
		Person: Person{
			Name: "小强",
			Age:  20,
		},
		Score: 100,
	}
	// 2.调用，传入不同对象的引用，也就是传入对象地址，然后函数自动调用
	SayHi(&tea) // 大家好，我是张三， 我30岁了， 我教英语
	SayHi(&stu) // 大家好，我是小强， 我20岁了， 我的成绩100
	//SayHi(tea) // 大家好，我是张三， 我30岁了， 我教英语
	//SayHi(stu) // 大家好，我是小强， 我20岁了， 我的成绩100
}

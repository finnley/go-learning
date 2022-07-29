package main

import "fmt"

type EmptyInterface interface {
}

type Course struct {
	name  string
	price int
	url   string
}

type Printer interface {
	printInfo() string
}

func (c *Course) printInfo() string {
	return "课程信息"
}

type AliOss struct {
}

type LocalFile struct {
}

func store(x interface{}) {
	switch v := x.(type) {
	case AliOss:
		// todo 此处要做一些特殊处理，比如设置权限等
		fmt.Println(v)
	case LocalFile:
		//todo 检查路径
		fmt.Println(v)
	}

}

func print(x interface{}) {
	// 看下x是不是int类型,如果是int类型 v,就是x的值,ok就是true,否则ok是false
	//if v, ok := x.(int); ok {
	//	fmt.Printf("%d，整数\n", v)
	//}
	//if s, ok := x.(string); ok {
	//	fmt.Printf("%s，字符串\n", s)
	//}

	switch v := x.(type) {
	case string:
		fmt.Printf("%s，字符串\n", v)
	case int:
		fmt.Printf("%d，整数\n", v)
	}
	//fmt.Printf("%v\n", x)
}

func main() {
	// 空接口
	var i interface{}
	// 空接口类似java和python中的object
	// 1. 空接口的第一个用途可以把任何类型赋值给空接口变量 类似多态，只不过有些特殊，什么类型都可以接收
	i = Course{}
	fmt.Println(i) // { 0 }
	i = 10

	// 2.参数传递
	print(i)
	i = "golang"
	print(i)
	i = []string{"python"}
	print(i)
	fmt.Println(i)

	// 3.空接口可以作为map的值
	//var teacherInfo = make(map[string]string)
	var teacherInfo = make(map[string]interface{})
	teacherInfo["name"] = "张三"
	teacherInfo["age"] = "18"
	teacherInfo["age"] = 18
	teacherInfo["weight"] = "60.4"
	teacherInfo["weight"] = 60.5
	fmt.Printf("%v\n", teacherInfo)

	// 接口的一个坑，接口引入了指针，编译就会报错
	// 接口有一个默认规范，接口名称一般以 er 结尾
	//c := &Course{}
	//var c Printer = Course{}
	//c.printInfo()

	// 类型断言
	// 指明值的时候可以将变量指明interface 类型，但是有时候也想知道类型是什么
	//i.(int) // 看下i是不是int类型
}

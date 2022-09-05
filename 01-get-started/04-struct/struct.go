package main

import "fmt"

type student struct {
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main03() {
	// 定义结构体变量
	stu := student{"小明", 20, "男", 100, "江苏"}
	stu.sex = "女"            // 修改成员
	fmt.Println(stu)         // {小明 20 女 100 江苏}
	fmt.Printf("%p\n", &stu) // 0xc00010e040
	fmt.Println(&stu.name)   // 0xc00010e040  16Byte
	fmt.Println(&stu.age)    // 0xc00010e050
	fmt.Println(&stu.sex)    // 0xc00010e058
	fmt.Println(&stu.score)  // 0xc00010e068
	fmt.Println(&stu.addr)   // 0xc00010e070
}

type Student struct {
	name string
	age  int
}

func main() {
	m := make(map[int]Student)
	stu := Student{"小张", 20}
	m[1001] = stu
	m[1002] = Student{"小王", 50}
	fmt.Println(m)
	// 可以读取结构体中某个成员信息
	fmt.Println(m[1001].name, m[1001].age) // 小张 20

	// 但是却不能修改结构体中某个成员，
	//m[1001].name = "法师" // error: cannot assign to struct field m[1001].name in map 不能改变结构体内容的元素

	// 如果想修改结构体必须整体操作
	for k, v := range m {
		v.name = "小钱"
		//m[k].name = "小号"
		fmt.Println(k, v)
	}
}

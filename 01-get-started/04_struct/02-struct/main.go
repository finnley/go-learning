package main

import (
	"fmt"
	"unsafe"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

func (c Course) printCourseInfo() {
	fmt.Printf("课程名: %s, 课程价格: %d, 课程地址: %s\n", c.Name, c.Price, c.Url)
}

func (c *Course) setPrice(price int) {
	c.Price = price // 等价于 (*c).Price = price

}

// 1. 结构体的方法只能和结构体在一个包中
// 2. 内置的int类型不能加方法，接受者可以是任意类型，如果要给int加方法可以使用下面方法
type MyInt int

func (m MyInt) toString() {

}

func main() {
	// go不支持面向对象
	// 面向对象三大特征：封装，继承，多态。go也是可以实现的
	// go不支持方法重载，抽象基类
	// 定义 struct go没有class
	// 1. 实例化 key-value 形式
	var c Course = Course{
		Name:  "django",
		Price: 100,
		Url:   "learn.com",
	}
	//访问
	fmt.Println(c.Name, c.Url, c.Price)

	// 大小写在go中的重要性，影响一些对象的可见性
	// 一个包中的变量或者结构体如果首字母是小写，那么对于另一个包不可见
	// 结构体定义的名称以及属性首字母大小写要注意

	// 2. 实例化 顺序形式
	c2 := Course{"python", 90, "cat.com"}
	fmt.Println(c2.Name, c2.Url, c2.Price)

	// 3.如果一个指向结构体的指针 通过结构体指针获取对象的值
	c3 := &Course{"golang", 200, "go.com"}
	fmt.Printf("%T\n", c3)                 // *main.Course
	fmt.Println(c3.Name, c3.Url, c3.Price) // 和下面效果一样 这里是go语言的语法糖 如果是一个指向结构体的指针，go内部会将c3.Name转成(*c3).Name
	fmt.Println((*c3).Name, (*c3).Url, (*c3).Price)

	// 4.零值 如果不给结构体赋值，go会默认给每隔字段采用默认值，默认值是类型的默认值，并不是nil对象，实例化的时候会开辟一块空间，并把空间的每一个值都赋值为默认值
	c4 := Course{}
	fmt.Println(c4.Price) // 0

	// 5.多种方式零值初始结构体
	var c5 Course = Course{}
	var c6 Course
	var c7 *Course = new(Course) //指针类型使用new初始化，new分配内存，并将这块内存初始化好
	var cc *Course = &Course{}
	// 上面三种都是零值初始化方式
	// 零值初始化
	fmt.Println("零值初始化")
	fmt.Println(c5.Price) // 0
	fmt.Println(c6.Price) // 0
	fmt.Println(c7.Price) // 0
	fmt.Println(cc.Price) // 0

	//var c8 *Course // 指针如果只申明不赋值，默认值是 close_channel,c6不是指针，是结构体类型，c6和常用的 var c9 int 是没有区别的，
	//int在申明的时候内存会分配好并设置默认值，结构体也是一样
	//fmt.Println(c8.Price) // panic: runtime error: invalid memory address or close_channel pointer dereference

	// 6.结构体是值类型
	c8 := Course{"python", 90, "cat.com"}
	c9 := c8
	fmt.Println(c8)
	fmt.Println(c9)
	c8.Price = 200
	fmt.Println(c8)
	fmt.Println(c9)

	// 7.结构体的大小 占用内存的大小，可以使用sizeof查看对象占用的内存
	fmt.Println(unsafe.Sizeof(1))                          // 8 1是int64,占用8个字节
	fmt.Println(unsafe.Sizeof("scrapy"))                   // 16
	fmt.Println(unsafe.Sizeof("scrapy-python-golang-php")) // 16 无论多长都是16 这里涉及到golang string类型的本质
	fmt.Println(unsafe.Sizeof(""))                         // 16
	fmt.Println(unsafe.Sizeof(c8))                         // 40字节 = 16+8+16
	// string是个结构体 当计算大小时候实际上是计算结构体里面数量占了多少个数，当请求字符串长度的时候就是计算头多大
	type string2 struct {
		Data uintptr // 指针占8个长度 Data类似切片，这个是个指针，指向的是数据域
		Len  int     // 长度64位系统占8个长度 Len表示数据域有多少字符串
	}

	// 8. slice的大小
	type slice struct {
		array unsafe.Pointer // 底层数组地址
		len   int            // 长度 1
		cap   int            // 容量
	}
	s1 := []string{"django", "tornado", "scrapy", "celery"}
	fmt.Println(unsafe.Sizeof(s1)) // 24字节 8+8+8

	m1 := map[string]string{
		"lan":  "php",
		"lan2": "python",
		"lan3": "golang",
	}
	fmt.Println("m1: ", unsafe.Sizeof(m1)) // 8

	// 9.结构体方法 可以达到封装数据和封装方法的效果
	c10 := Course{"scrapy", 100, "baidu.com"}
	c10.printCourseInfo() // 课程名: scrapy, 课程价格: 100, 课程地址: baidu.com
	// 上面的调用类似下面的调用效果一样
	Course.printCourseInfo(c10) // 课程名: scrapy, 课程价格: 100, 课程地址: baidu.com

	c10.setPrice(200)      // 修改了c10的price 本质是 Course.setPrice(c10, 200)
	fmt.Println(c10.Price) // 并没有改变，还是100 为什么？
	// 本质是这句 Course.setPrice(c10, 200), 只是写成了 c10.setPrice(200) 这种写法叫做语法糖
	// 函数参数的传递是怎么传递的？c10是结构体，结构体是值传递 也就是c10传进去会被复制一份，如果想要改变值，就需要设置为指针类型,
	/**
	func (c *Course) setPrice(price int) {
		c.Price = price
	}
	*/

	// 这句 c10.setPrice(200) 等价于 (&c10).setPrice(200)   // 修改了c10的price

	// 结构体的接受者有两种形式 1. 值传递， 2。 指针传递
	// a.如果想要改结构体的值；b.如果结构体的数据很大，采用值传递会复制值，性能就会降低，这种情况可以使用指针传递
	// 如果不想产生副作用，就不用指针传递，因为有可能不小心改了里面的值

	// 10.内嵌结构体实现继承 通过组合实现
	t := Teacher{
		Name:  "finnley",
		Age:   58,
		Title: "Golang",
	}
	cccc := Course2{
		Teacher: t,
		Name:    "php",
		Price:   100,
		Url:     "aa.com",
	}
	cccc.courseInfo()
	//上面可以改成匿名嵌套
	cccc.courseInfo2()
}

type Teacher struct {
	Name  string
	Age   int
	Title string
}

//05-type Course2 struct {
//	Teacher Teacher // 嵌套 将另一个结构体嵌套
//	Name    string
//	Price   int
//	Url     string
//}

type Course2 struct {
	Teacher // 匿名嵌套
	Name    string
	Price   int
	Url     string
}

func (t Teacher) teacherInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 职称: %s\n", t.Name, t.Age, t.Title)
}

func (c Course2) courseInfo() {
	fmt.Printf("课程名: %s, 价格: %d, 讲师信息: 姓名:%s %d %s\n", c.Name, c.Price, c.Teacher.Name, c.Teacher.Age, c.Teacher.Title)
}

func (c Course2) courseInfo2() {
	fmt.Printf("课程名: %s, 价格: %d, 讲师信息: 姓名:%s %d %s\n", c.Name, c.Price, c.Teacher.Name, c.Age, c.Title)
}

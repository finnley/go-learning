package main

import "fmt"

func main() {
	// Go 语言中的数组和 Python 中的 list 可以对应起来理解，但是Go中slice和python的list更像 ，但是也存在区别：
	// Python 中的 list是可以随意变动的，而Go中的array大小一旦定义，则不允许变更，但是Go中的 slice 可以灵活变动
	// Go语言中的Array特点： 1.数组大小确定 2.类型一致

	// ******** 定义 var courses [10]string ********
	fmt.Println("******** 定义 ********")
	//// 1.根据值类型进行类型推断 定义了一个空数组
	//var courses = [10]string{}
	//fmt.Println(courses)         // [         ]
	//fmt.Printf("%#v\n", courses) // [10]string{"", "", "", "", "", "", "", "", "", ""}

	//var courses = [10]string{"django", "scrapy", "tornado"} // ok
	// 前后数组大小要一致
	//var courses [10]string = [5]string{"django", "scrapy", "tornado"} // error: 前后数组大小要一致，所以推荐上面写法，不指定数组类型，使用类型推断

	// 2.还可以这样简写，声明并初始化赋值
	//courses := [5]string{"django", "scrapy", "tornado", 100} // error: 数组类型不一致，有string 和 int类型，因为静态语言要求严格
	courses := [5]string{"django", "scrapy", "tornado"}
	fmt.Println(courses)         // [django scrapy tornado  ]
	fmt.Printf("%#v\n", courses) // [5]string{"django", "scrapy", "tornado", "", ""}

	// 3.数组的另一种创建方式，根据值得个数动态初始化数组大小 ...
	var a1 = [4]float32{}
	fmt.Println(a1) // [0 0 0 0]
	//a1 = [4]float32{1.0, 1.1, 1.2, 1.3, 1.4} // Index out of bounds: 4
	var b1 = [5]int{'A', 'B', 'a', 'b'}
	fmt.Println(b1) // [65 66 97 98 0]
	// 省略号，表示这是数组，长度是根据初始化的时候里面的数据多少决定的
	d1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(d1) // [1 2 3 4 5]
	// 根据数组索引下标初始化数组元素，比如第4个值设置成100，其他是默认值
	e1 := [5]int{4: 100}
	fmt.Println(e1) // [0 0 0 0 100]
	f1 := [...]int{0: 1, 4: 8, 9: 100}
	fmt.Println(f1) // [1 0 0 0 8 0 0 0 0 100]

	// array of 3 integers 定义一个数组，没有设置初始值值时，数组的元素为数组元素类型的零值
	var a [3]int
	// print the first element a[0] = 0
	fmt.Println(a[0]) // 0
	// print the last element, a[2] = 0
	fmt.Println(a[len(a)-1]) // 0

	// ******** Array 操作：修改值、取值 ********
	fmt.Println("\n******** Array 操作 ********")
	// 数组不支持删除、添加操作，因为Go语言数组大小一旦定义就不允许修改，大小一旦固定内存空间就固定了，不可能重新分配
	// 取值
	fmt.Println("\n******** 取值 ********")
	fmt.Println(courses[0]) // django

	// 修改值
	fmt.Println("\n******** 改值 ********")
	courses[0] = "php"
	fmt.Println(courses) // [php scrapy tornado  ]

	// 求长度
	fmt.Println("\n******** 求长度 ********")
	fmt.Println(len(f1)) // 10

	// 遍历
	fmt.Println("\n******** 遍历 ********")
	// print the indices and elements
	for i, v := range courses {
		fmt.Printf("%d %#v\n", i, v)
	}
	/**
	0 "php"
	1 "scrapy"
	2 "tornado"
	3 ""
	4 ""
	*/

	sum := 0
	for _, v := range f1 {
		sum += v
	}
	fmt.Println("sum of f1: ", sum) // sum of f1:  109

	sum = 0
	for i := 0; i < len(f1); i++ {
		sum += f1[i]
	}
	fmt.Println("sum of f1: ", sum) // sum of f1:  109

	// ******** 数组是值类型 ********
	fmt.Println("******** 数组是值类型 ********")
	// 数组是值类型
	courseA := [3]string{"django", "scrapy", "tornado"}
	courseB := [...]string{"c/c++", "java", ".net", "python", "php"}
	// courseA和courseB都是数组类型，但是不是同一种类型
	// Go语言中数组长度是数组类型的一部分，类型相同长度不同不是同一数组类型
	fmt.Printf("%T\n", courseA) // [3]string
	fmt.Printf("%T\n", courseB) // [5]string
	printArray(courseB)
	fmt.Println(courseB) // [c/c++ java .net python php]

	aa := [3]int{1, 2, 3}
	bb := aa
	bb[0] = 4
	fmt.Println(aa) // [1 2 3]
	fmt.Println(bb) // [4 2 3]

	printAddrOfArrays()
}

// 数组作为参数，实际调用时是值传递
func printArray(toPrint [5]string) {
	toPrint[0] = "Go"
	fmt.Println(toPrint) // [Go java .net python php]
}

func printAddrOfArrays() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a: %p\n", &a) // a: 0xc000020180
	b := a
	copyArray(a)
	fmt.Printf("b: %p\n", &b) // b: 0xc0000201b0
}

func copyArray(c [5]int) {
	fmt.Printf("c: %p\n", &c) // c: 0xc0000201e0
}

// 数组：一组具有相同数据类型在内存中有序存储的数据集合
// 相同数据类型
// 有序存储
// 数据集合
// 数组的长度在定义后不可以修改（不但固定不能修改）
// 格式：var 数组名 [元素个数]元素数据类型
func main0202() {
	//var arr [10]int // 默认值为该数组元素类型的0值
	//// 1. 使用数组名+下标进行数组初始化 下标范围[0, 数组最大元素个数-1]
	//arr[0] = 123
	////arr[-1] = 123
	//fmt.Println(arr) // [123 0 0 0 0 0 0 0 0 0]

	// 2. 通过赋值的方式进行数组初始化
	//var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 3. 通过简短方式初始化并赋值
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(arr) // [1 2 3 4 5 6 7 8 9 10]

	// 数据长度 len(array)
	//fmt.Println(len(arr)) // 10

	// 遍历数组元素 for
	//for i := 0; i < len(arr); i++ {
	//	fmt.Printf("%d ", arr[i]) // 1 2 3 4 5 6 7 8 9 10
	//}
	//fmt.Println()

	// i: 下标 v: 值
	for i, v := range arr {
		fmt.Printf("%d=>%d\n", i, v)
	}
}

func main03() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		fmt.Println(&arr[i]) // 0xc0000b8000 0xc0000b8008 0xc0000b8010  ...
	}
	/**
	0xc0000b8000
	0xc0000b8008
	0xc0000b8010
	0xc0000b8018
	0xc0000b8020
	0xc0000b8028
	0xc0000b8030
	0xc0000b8038
	0xc0000b8040
	0xc0000b8048
	*/
}

/**
数据：
1. 固定长度
2. 类型唯一
3. 数据序列

数组长度是数组类型的一个组成部分，[3]int 和 [4]int就不是同一类型，而是两个不同类型，数组长度必须是常量表达式，因为数组的长度需要在编译阶段确定
*/

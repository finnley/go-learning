package main

import (
	"fmt"
)

func main() {
	//courses := []string{}
	slice := make([]int, 1)
	slice[0] = 10
	fmt.Println(slice)
}

func main0203() {
	// 切片初始化方式
	//fmt.Println("******** 切片初始化方式 ********")
	//var courses = []string{"django", "scrapy", "tornado"}
	//fmt.Printf("%T\n", courses)                     // []string
	//fmt.Println(courses)                            // [django scrapy tornado]
	//fmt.Printf("values of courses: %#v\n", courses) // values of courses: []string{"django", "scrapy", "tornado"}
	////courses[3] = "php"
	////courses = append(courses, "php")
	//fmt.Println(courses)

	// 第二种: 初始化方法 make
	fmt.Println("\n******** 初始化方法 make ********")
	// 1.切片不是没有长度限制吗，为什么使用 make 初始化的时候需要传递一个长度？
	// 2.传递了长度之后是不是就意味着像数组一样长度不能变了呢？
	courses2 := make([]string, 5)
	fmt.Printf("%T\n", courses2)           // []string
	fmt.Println(len(courses2))             // 5
	fmt.Println(courses2)                  // 5
	fmt.Printf("course2: %#v\n", courses2) // course2: []string{"", "", "", "", ""}
	//// 如果传两个参数，表示申请5个元素并初始化5个元素，
	//// 申请0个元素就就不初始化
	//courses2_copy := make([]string, 0, 5) // 空切片
	//fmt.Println(courses2_copy)
	//fmt.Printf("course2_copy: %#v\n", courses2_copy) // []string{}
	//names3 := make([]string, 3)
	//names3[0] = "a"
	//names3[1] = "b"
	//names3[2] = "c"
	//fmt.Println(names3[0], names3[1], names3[2])
	////names[3] = "d"
	////fmt.Println(names3[3])

	//// 第三种: 通过数组转切片
	//fmt.Println("\n******** 通过数组转切片 ********")
	//var courses3 = [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	//subCourse := courses3[1:4]
	//fmt.Printf("%T\n", subCourse) // []string
	//fmt.Println(subCourse)        // [scrapy tornado python]
	//subCourse[0] = "php"
	//fmt.Println(subCourse) // [php tornado python]
	//fmt.Println(courses3)  // [django php tornado python asyncio]

	//// 第四种方式: new
	//fmt.Println("\n******** new ********")
	//subCourse2 := *new([]string)
	//fmt.Println(subCourse2) // []

	// 数组是值传递，切片是引用传递
	fmt.Println("\n******** 数组是值传递，切片是引用传递 ********")
	var courses4 = [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	subCourse4 := courses4[1:4]
	fmt.Println(subCourse4) // [scrapy tornado python]
	fmt.Println("====== replace ======")
	replace(subCourse4)
	fmt.Println(courses4) // [django java tornado python asyncio]

	fmt.Println("\n******** 切片操作 ********")
	// slice 是动态数组，可以动态添加值
	var courses5 = [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	subCourse5 := courses5[1:4] // [scrapy tornado python]
	fmt.Println(subCourse5[1])  // tornado
	subCourse5[1] = "php"
	fmt.Println(subCourse5) // [scrapy php python]
	fmt.Println(courses5)   // [django scrapy php python asyncio]
	// 在slice截取
	fmt.Println("\n******** slice的截取 ********")
	var courses6 = [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	subCourse6 := courses6[1:4] // [scrapy tornado python]
	subCourse62 := subCourse6[1:3]
	fmt.Printf("%T, %v\n", subCourse62, subCourse62) // []string, [tornado python]
	subCourse621 := subCourse6[1:]
	fmt.Printf("%v\n", subCourse621) // [tornado python]
	subCourse622 := subCourse6[1:4]
	fmt.Printf("%v\n", subCourse622) // [tornado python asyncio]
	fmt.Printf("%v\n", courses6)     // [django scrapy tornado python asyncio]

	//subCourse623 := subCourse6[1:5]
	//fmt.Printf("%v\n", subCourse623) // panic: runtime error: slice bounds out of range [:5] with capacity 4

	// append 向切片追加元素
	fmt.Println("\n******** append ********")
	subCourse62 = append(subCourse62, "java")
	fmt.Printf("%v\n", subCourse62) // [tornado python java]
	subCourse621 = subCourse6[1:]
	fmt.Printf("%v\n", subCourse621) // [tornado python]
	subCourse622 = subCourse6[1:4]
	fmt.Printf("%v\n", subCourse622) // [tornado python java]
	fmt.Printf("%v\n", courses6)     // [django scrapy tornado python java]

	// append 函数追加多个元素
	fmt.Println("\n******** 函数追加多个元素 ********")
	subCourse62 = append(subCourse62, "java", "php", "c#")
	fmt.Printf("%v\n", subCourse62) // [tornado python java java php c#]
	fmt.Printf("%v\n", courses6)    // [django scrapy tornado python java] // 这里因为扩容了，内存地址变了

	fmt.Println("\n******** copy ********")
	// copy 拷贝的时候目标对象长度需要设置好，虽然是切片，但是还是有空间的，不使用make，默认的方法长度是0，复制过来的时候没有空间
	//subCourse7 := []string{} // []
	subCourse7 := make([]string, len(subCourse62)) // [tornado python java java php c#]
	copy(subCourse7, subCourse62)
	fmt.Printf("%v\n", subCourse7) // [tornado python java java php c#]

	// 切片合并
	fmt.Println("\n******** 切片合并 ********")
	appendedCourse := []string{"java", "php", "c#"}
	var courses8 = [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	subCourses8 := courses8[1:4] // [scrapy tornado python]
	subCourses8 = append(subCourses8, appendedCourse...)
	fmt.Printf("%v\n", subCourses8) // [scrapy tornado python java php c#]
	fmt.Printf("%v\n", courses8)    // [django scrapy tornado python asyncio]

	// 从切片中删除元素
	fmt.Println("\n******** 切片删除 ********")
	deleteCourses := [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	courseSlice := deleteCourses[:]
	fmt.Printf("%v\n", courseSlice) // [django scrapy tornado python asyncio]
	courseSlice = append(courseSlice[:1], courseSlice[2:]...)
	fmt.Printf("%v\n", courseSlice)   // [django tornado python asyncio]
	fmt.Printf("%v\n", deleteCourses) // [django tornado python asyncio asyncio]
}

func replace(mySlice []string) {
	mySlice[0] = "java"
	fmt.Println(mySlice) // [java tornado python]
}

//切片：一组具有相同数据类型在内存中有序存储的可扩容的数据集合
func main0209() {
	//// 切片的定义和使用
	//var slice []int
	//fmt.Println(len(slice)) // 0
	//// 定义了一个切片，直接往里面添加元素会出错
	////slice[0] = 1
	//fmt.Println(slice) // panic: runtime error: index out of range [0] with length 0

	// 使用make方式定义并初始化长度为10的切片
	// make([]数据类型, 长度)
	//var slice []int = make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]
	var slice []int = make([]int, 0, 10) // []
	//slice[0] = 1
	fmt.Println(slice)

	//append 对切片进行扩容操作
	//var slice []int
	//fmt.Println(len(slice)) // 0 计算长度
	//fmt.Println(cap(slice)) // 0 计算容量
	//slice = append(slice, 1, 2, 3, 4, 5)
	//fmt.Println(slice)
	//fmt.Println(len(slice)) // 5 计算长度
	//fmt.Println(cap(slice)) // 6 计算容量

	// 切片扩容机制 源码growthslice
	// 切片扩容为偶数值，小于1024为上一次的两倍，大于1024为上一次的1/4
	// 在使用append扩容的时候内存地址一定会发生改变的，扩容后就是两个不同的内存地址了，见源码growslice memmove一定修改了内存地址
	// 当切片进行拷贝的时候内存地址如果超出了容量的限制一定会修改内存地址的，此时两个切片就不会相互影响了，如果没扩容之前指向同一地址就会影响
}

func main0204() {
	// 切片截取
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 切片名[起始下标, 结束下标, 容量] 左开右闭 包含起始下标，不包含结束下标
	//s := slice[2:7] // [3, 4, 5, 6, 7]
	//fmt.Println(s)

	//s := slice[2:] // [3 4 5 6 7 8 9 10]
	//fmt.Println(s)

	//s := slice[:5] // [1 2 3 4 5]
	//fmt.Println(s)

	s := slice[2:5:6] // [3, 4, 5] len=5-2=3 cap=6-2=4
	//fmt.Println(s)
	//fmt.Println(len(s))
	//fmt.Println(cap(s))
	s[0] = 333
	// 切片的截取是将新的切片指向源切片的内存地址 修改一个会影响另外一个
	fmt.Println(s)     // [333, 4, 5]
	fmt.Println(slice) // [1 2 333 4 5 6 7 8 9 10]
}

func main0207() {
	slice := []int{1, 2, 3, 4, 5}
	//s := slice // 使用这种方式相当于在内存中拷贝了一份，但是两个指向同一个底层数组
	//s[2] = 333
	//fmt.Println(s) // [1 2 333 4 5]
	//fmt.Println(slice) // [1 2 333 4 5]
	// 在存储两个内容完全相同，但是不会相互影响，使用copy就不会相互影响，可以看做数据也会拷贝一份
	s := make([]int, 5)
	copy(s, slice)
	s[2] = 333
	fmt.Println(s)     // [1 2 333 4 5]
	fmt.Println(slice) // [1 2 3 4 5]

	//fmt.Println("切片队列应用============")
	//queue := []string{}
	//// 队列操作：
	//// push
	//queue = append(queue, "a", "b")
	//queue = append(queue, "c")
	//// pop
	//x := queue[0]
	//queue = queue[1:]
	//fmt.Println("1: ", x)
	//
	//x = queue[0]
	//queue = queue[1:]
	//fmt.Println("2: ", x)
	//
	//x = queue[0]
	//queue = queue[1:]
	//fmt.Println("3: ", x)
	//
	//fmt.Println("切片栈应用============")
	//stack := []string{}
	//// 栈操作：
	//// push
	//stack = append(stack, "a", "b")
	//stack = append(stack, "c")
	//// pop
	//y := stack[len(stack)-1]
	//stack = stack[:len(stack)-1]
	//fmt.Println("1: ", y)
	//
	//y = stack[len(stack)-1]
	//stack = stack[:len(stack)-1]
	//fmt.Println("2: ", y)
	//
	//y = stack[len(stack)-1]
	//stack = stack[:len(stack)-1]
	//fmt.Println("3: ", y)
	//
	//fmt.Println("切片操作 ==============")
	//nums := []int{0, 1, 2, 3, 4, 5} // len=6 cap=6
	//numChildren := nums[1:3]
	//fmt.Printf("%T, %#v\n", numChildren, numChildren)
	//nums = make([]int, 0, 3)
	//nums = []int{2, 3, 4}
	//// 移除中间元素3(index)
	//// 切片操作, copy
	//// [index:] [3, 4]
	//// [index+1:] [4]
	//// copy([index:], [index+1:]) [2, 4, 5]
	//copy(nums[1:], nums[2:])
	//fmt.Println(nums)
	//fmt.Println(nums[:len(nums)-1])
	//nums = []int{1, 3, 5, 9, 19}
	//fmt.Println(sort.SearchInts(nums, 8))
	//
	//scores := map[string]float32{"aa": 13, "bb": 34, "cc": 45}
	//fmt.Println(scores)
	//delete(scores, "aa")
	//fmt.Println(scores)
	//delete(scores, "dd")
	//fmt.Println(scores)
	//
	////var nilMap map[string]string
	////fmt.Println(nilMap)
	////fmt.Println(nilMap["xx"])
	////nilMap["xx"] = "yy"
	////fmt.Println(nilMap)
	//
	//// [168 180 166 176 169]
	//heigh := []int{168, 180, 166, 176, 169}
	//for j := 0; j < len(heigh)-1; j++ {
	//	for i := 0; i < len(heigh)-1; i++ {
	//		if heigh[i] > heigh[i+1] {
	//			heigh[i], heigh[i+1] = heigh[i+1], heigh[i]
	//		}
	//		fmt.Println(i, heigh)
	//	}
	//	fmt.Println(heigh)
	//}
}
